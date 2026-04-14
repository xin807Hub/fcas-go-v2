package policy

import (
	"encoding/json"
	"fcas_server/global"
	modelConfiguration "fcas_server/model/configuration"
	"fcas_server/model/policy"
	"fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type DimControlPolicyService struct {
}

func (DimControlPolicyService) List() (result []policy.DimControlPolicyResp, err error) {
	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Find(&result).Error; err != nil {
		global.Log.Error("策略配置control失败", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (DimControlPolicyService) PageControlPolicy(policyReq policy.DimControlPolicyReq) (result []policy.DimControlPolicyResp, total int64, err error) {
	limit := policyReq.Limit
	offset := policyReq.Limit * (policyReq.Page - 1)

	sql := `select %s from (
            select a.*, c.crowd_name as user_crowd_name from (
                select a.*, g.group_name as user_crowd_group_name from (
                    select c.*, b.app_name from (
                        select w.*, a.app_type_name from dim_control_policy w left join (select DISTINCT app_type_id, app_type_name from dim_app_classify) a on w.app_type_id = a.app_type_id
                    ) c
                    left join (select DISTINCT app_type_id, app_id, app_name from dim_app_classify) b on c.app_id = b.app_id and c.app_type_id = b.app_type_id
                ) a left join dim_user_crowd_group g on a.user_crowd_group_id=g.id
            ) a left join dim_user_crowd c on a.user_crowd_id = c.id
        ) a left join dim_user_info u on a.user_id = u.id where 1=1 `

	params := make([]interface{}, 0)

	if len(policyReq.PolicyName) > 0 {
		sql += " and a.name like ? "
		params = append(params, "%"+policyReq.PolicyName+"%")
	}
	if len(policyReq.UerType) > 0 {
		sql += " and a.user_type = ? "
		params = append(params, policyReq.UerType)
	}
	if len(policyReq.UerCrowdGroupId) > 0 {
		sql += " and a.user_crowd_group_id = ? "
		params = append(params, policyReq.UerCrowdGroupId)
	}
	if len(policyReq.UerCrowdId) > 0 {
		sql += " and a.user_crowd_id = ? "
		params = append(params, policyReq.UerCrowdId)
	}
	if len(policyReq.UerName) > 0 {
		sql += " and a.user_crowd_name = ? "
		params = append(params, policyReq.UerName)
	}
	if len(policyReq.UerId) > 0 {
		sql += " and a.user_id = ? "
		params = append(params, policyReq.UerId)
	}
	if policyReq.AppTypeId != 0 {
		sql += " and a.app_type_id = ? "
		params = append(params, policyReq.AppTypeId)
	}
	if policyReq.AppTypeId != 0 {
		sql += " and a.app_id = ? "
		params = append(params, policyReq.AppId)
	}

	sqlCount := fmt.Sprintf(sql, "count(*)")

	if err = global.ServiceDB.Raw(sqlCount, params...).Scan(&total).Error; err != nil {
		global.Log.Error("策略控制配置count查询失败", zap.Error(err))
		return nil, 0, err
	}

	sqlSelect := fmt.Sprintf(sql, "a.*, u.user_name")

	sqlSelect += " limit ? offset ?"
	params = append(params, limit)
	params = append(params, offset)

	if err := global.ServiceDB.Raw(sqlSelect, params...).Scan(&result).Error; err != nil {
		global.Log.Error("策略控制配置page信息查询失败", zap.Error(err))
		return nil, 0, err
	}

	// 遍历设置值 DimControlPolicyLog
	for i := 0; i < len(result); i++ {
		dto := result[i]
		var policyLog []policy.DimControlPolicyLog
		if err := global.ServiceDB.Model(&policy.DimControlPolicyLog{}).Where("policy_id = ?", dto.Id).Find(&policyLog).Error; err != nil {
			global.Log.Error("策略控制配置日志信息查询失败", zap.Error(err))
			continue
		}
		if len(policyLog) > 1 {
			times := (policyLog[0].RecordTime.UnixMilli() - policyLog[1].RecordTime.UnixMilli()) / 1000
			result[i].UpTrafficSpeed = calculating(policyLog[0].UpTraffic, policyLog[1].UpTraffic, int(times))
			result[i].DnTrafficSpeed = calculating(policyLog[0].DnTraffic, policyLog[1].DnTraffic, int(times))
			result[i].UpPassSpeed = calculating(policyLog[0].UpPass, policyLog[1].UpPass, int(times))
			result[i].DnPassSpeed = calculating(policyLog[0].DnPass, policyLog[1].DnPass, int(times))
			result[i].UpDiscardSpeed = calculating(policyLog[0].UpDiscard, policyLog[1].UpDiscard, int(times))
			result[i].DnDisCardSpeed = calculating(policyLog[0].DnDiscard, policyLog[1].DnDiscard, int(times))
		}
	}

	return result, total, nil
}

func calculating(value1 int, value2 int, times int) int {
	if value1-value2 > 0 && times > 0 {
		value := value1 - value2
		return value * 8 / times
	}
	return 0
}

func (d DimControlPolicyService) SaveOrUpdateControlPolicy(controlPolicy policy.DimControlPolicy) error {
	if controlPolicy.Id == 0 { // 新增
		controlPolicy.CreateTime = time.Now().Format(global.DateTimeLayout)
		if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Save(&controlPolicy).Error; err != nil {
			global.Log.Error("策略配置control失败", zap.Error(err))
			return err
		}
	} else { // 编辑
		if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).
			Where("id = ?", controlPolicy.Id).
			Updates(&controlPolicy).Error; err != nil {
			global.Log.Error("策略配置control失败", zap.Error(err))
			return err
		}
	}

	// 保存策略后,根据生效时间发送绑定/解绑策略
	ifPolicyEffect, err := ifTimeEffect(controlPolicy)
	if err != nil {
		return err
	}
	var sendType = ""
	if ifPolicyEffect {
		// 下发类型
		sendType = global.Bind
	} else {
		sendType = global.UnBind
	}
	policyInfoStr := BuildPolicyJson(controlPolicy, sendType)
	global.Log.Info(" 策略绑定：" + policyInfoStr)
	utils.SendMessage(policyInfoStr, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)
	return nil
}

func ifTimeEffect(controlPolicy policy.DimControlPolicy) (bool, error) {
	// 解析开始时间和结束时间
	start, err := time.Parse(global.DateTimeLayout, controlPolicy.StartTime)
	if err != nil {
		fmt.Println("解析开始时间失败:", err)
		return false, err
	}

	end, err := time.Parse(global.DateTimeLayout, controlPolicy.EndTime)
	if err != nil {
		fmt.Println("解析结束时间失败:", err)
		return false, err
	}

	// 获取当前时间
	now := time.Now()

	// 判断当前时间是否在开始时间和结束时间之间
	if now.After(start) && now.Before(end) {
		fmt.Println("当前时间处于生效时间内")
		return true, nil
	} else {
		fmt.Println("当前时间不在生效时间内")
		return false, nil
	}
}

func (d DimControlPolicyService) InfoControlPolicy(id int) (result *policy.DimControlPolicy) {
	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id = ?", id).Find(&result).Error; err != nil {
		global.Log.Error("查询策略配置control失败", zap.Error(err))
		return nil
	}
	return result
}

func (d DimControlPolicyService) getByPolicyId(policyId int) (result *policy.DimControlPolicy) {
	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("policy_id = ?", policyId).Find(&result).Error; err != nil {
		global.Log.Error("策略policyId查询配置control失败", zap.Error(err))
		return nil
	}
	return result
}

func (d DimControlPolicyService) DeleteControlPolicy(ids []int) error {
	var result []policy.DimControlPolicy
	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id in ?", ids).Find(&result).Error; err != nil {
		global.Log.Error("查询策略配置control失败", zap.Error(err))
		return err
	}

	for _, whitePolicy := range result {
		// 发送解绑策略
		policyInfo := BuildPolicyJson(whitePolicy, global.UnBind)
		global.Log.Info(strconv.Itoa(whitePolicy.Id) + " 策略解绑：" + policyInfo)
		utils.SendMessage(policyInfo, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)
	}

	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id in ?", ids).Delete(policy.DimControlPolicy{}).Error; err != nil {
		global.Log.Error("删除策略配置control失败", zap.Error(err))
		return err
	}
	return nil
}

func BuildPolicyJson(controlPolicy policy.DimControlPolicy, bindFlag string) string {
	dpiPolicy := struct {
		MsgType       int         `json:"msg_type"`
		PolicyId      int         `json:"policy_id"`
		TimeInfo      interface{} `json:"time_info"`
		ControlPolicy interface{} `json:"control_policy"`
		FeatureList   interface{} `json:"feature_list"`
		BindObj       string      `json:"bind_obj"`
	}{
		MsgType:  policy.ControlPolicyMsgType,
		PolicyId: controlPolicy.Id,
	}

	// 构造 time_info
	timeInfo := struct {
		StartTime  string      `json:"start_time"`
		EndTime    string      `json:"end_time"`
		Period     string      `json:"period"`
		PeriodInfo interface{} `json:"period_info"`
	}{
		StartTime: controlPolicy.StartTime,
		EndTime:   controlPolicy.EndTime,
	}

	periodInfoList := make([]interface{}, 0)
	if len(controlPolicy.PolicyPeriod) == 0 {
		timeInfo.Period = "all"
		periodInfo := struct {
			StartPeriod string `json:"start_period"`
			EndPeriod   string `json:"end_period"`
		}{
			StartPeriod: "0",
			EndPeriod:   "0",
		}
		periodInfoList = append(periodInfoList, periodInfo)
	} else {
		if controlPolicy.PeriodType == 1 {
			timeInfo.Period = "day"
		} else if controlPolicy.PeriodType == 2 {
			timeInfo.Period = "week"
		}
		periodArr := strings.Split(controlPolicy.PolicyPeriod, ",")
		for i := 0; i < len(periodArr); i++ {
			periodItem := periodArr[i]
			periodInfo := struct {
				StartPeriod string `json:"start_period"`
				EndPeriod   string `json:"end_period"`
			}{
				StartPeriod: strings.Split(periodItem, "-")[0],
				EndPeriod:   strings.Split(periodItem, "-")[1],
			}
			periodInfoList = append(periodInfoList, periodInfo)
		}
	}
	timeInfo.PeriodInfo = periodInfoList
	dpiPolicy.TimeInfo = timeInfo

	// 构造control_policy
	control := struct {
		ThresholdUpIds int   `json:"threshold_up_ids"`
		ThreshlodDnIds int   `json:"threshlod_dn_ids"`
		LinkIds        []int `json:"link_ids"`
	}{
		ThresholdUpIds: controlPolicy.UlFlowRate * 1024,
		ThreshlodDnIds: controlPolicy.DlFlowRate * 1024,
	}

	lineList := getListByLinkIds(strings.Split(controlPolicy.LinkIds, ","))

	lineValArr := make([]int, 0)
	for i := 0; i < len(lineList); i++ {
		lineDto := lineList[i]
		lineValArr = append(lineValArr, lineDto.LineVlan)
	}
	control.LinkIds = lineValArr

	dpiPolicy.ControlPolicy = control

	// 构造 feature_list 构造应用大小类的json
	appInfoList := make([]interface{}, 0)

	if controlPolicy.FlowCtrlType == 1 { // 应用大小类
		if controlPolicy.AppTypeId != 0 {
			appInfo := struct {
				AppType int `json:"app_type"`
				AppId   int `json:"app_id"`
			}{
				AppType: controlPolicy.AppTypeId,
				AppId:   controlPolicy.AppId - controlPolicy.AppTypeId*10000,
			}
			appInfoList = append(appInfoList, appInfo)
		}
	}

	// 构造五元组信息
	userInfoList := getUserInfoList(controlPolicy)
	// 构造feature_list
	targetArr := make([]string, 0)
	dip := controlPolicy.DstIp
	if len(dip) > 0 {
		var dipArr []map[string]interface{}
		err := json.Unmarshal([]byte(dip), &dipArr)
		if err == nil {
			for i := 0; i < len(dipArr); i++ {
				mapDip := dipArr[i]
				dipStart := mapDip["dipStart"]
				dipEnd := mapDip["dipEnd"]
				dport := mapDip["dport"]
				targetItem := struct {
					DipStart interface{} `json:"dip_start"`
					DipEnd   interface{} `json:"dip_end"`
					Dport    interface{} `json:"dport"`
				}{
					DipStart: dipStart,
					DipEnd:   dipEnd,
					Dport:    dport,
				}

				byts, _ := json.Marshal(targetItem)
				targetArr = append(targetArr, string(byts))
			}
		}
	}

	tupleList := buildTupleJsonList(userInfoList, targetArr)

	featureList := map[string]interface{}{}
	if len(appInfoList) > 0 && len(tupleList) > 0 {
		featureList["combined"] = true
	} else {
		featureList["combined"] = false
	}
	featureList["tuple"] = tupleList
	featureList["app_info"] = appInfoList

	dpiPolicy.FeatureList = featureList
	dpiPolicy.BindObj = bindFlag
	policyJson, _ := json.Marshal(dpiPolicy)

	return string(policyJson)
}

func getUserInfoList(controlPolicy policy.DimControlPolicy) (userInfos []modelConfiguration.DimUserInfo) {
	// 构造user_list
	if controlPolicy.UserId != 0 {
		if err := global.ServiceDB.Model(&modelConfiguration.DimUserInfo{}).Where("id = ?", controlPolicy.UserId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据ID获取信息失败", zap.Error(err))
			return nil
		}
		return userInfos
	} else if controlPolicy.UserCrowdId != 0 {
		sql := `SELECT
            tb1.*
        FROM
            dim_user_info tb1,
            dim_user_crowd tb2,
            dim_user_crowd_relation tb3
        where tb1.id = tb3.user_id
        AND tb2.id = tb3.crowd_id AND tb2.id = ?`

		if err := global.ServiceDB.Raw(sql, controlPolicy.UserCrowdId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据CrowdId获取用户信息失败", zap.Error(err))
			return nil
		}
		return userInfos
	} else if controlPolicy.UserCrowdGroupId != 0 {
		sql := `select distinct u.* from
            dim_user_info u,dim_user_crowd uc,dim_user_crowd_relation ur, dim_user_crowd_group_relation ugr
        where u.id = ur.user_id and uc.id = ur.crowd_id and uc.id = ugr.crowd_id and ugr.group_id = ? `
		if err := global.ServiceDB.Raw(sql, controlPolicy.UserCrowdGroupId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据UserCrowdGroupId获取用户信息失败", zap.Error(err))
			return nil
		}
		return userInfos
	}
	return userInfos
}

// 构造tupleList
func buildTupleJsonList(userInfos []modelConfiguration.DimUserInfo, dstAddrArray []string) []interface{} {
	var tupleList = make([]interface{}, 0)
	if len(userInfos) == 0 && len(dstAddrArray) > 0 {
		for i := 0; i < len(dstAddrArray); i++ {
			tupleItem := map[string]string{}
			jsonStr := dstAddrArray[i]
			mapStr := map[string]string{}
			err := json.Unmarshal([]byte(jsonStr), mapStr)
			if err != nil {
				dipStart := mapStr["dipStart"]
				dipEnd := mapStr["dipEnd"]
				dport := mapStr["dport"]
				if dipStart != "0" {
					tupleItem["dip"] = dipStart + "-" + dipEnd
				}
				if dport != "-1" {
					tupleItem["dport"] = dport
				}
				tupleList = append(tupleList, tupleItem)
			}
		}
	} else if len(userInfos) > 0 && len(dstAddrArray) == 0 {
		for i := 0; i < len(userInfos); i++ {
			userDto := userInfos[i]
			ipSectionArray := userDto.IpAddress
			for j := 0; j < len(ipSectionArray); j++ {
				ipSection := ipSectionArray[j]
				tupleItem := map[string]string{}
				if strings.Contains(ipSection, "/") {
					if strings.Contains(ipSection, ":") {
					} else {
						startIp := utils.GetStartIp(ipSection)
						endIp := utils.GetEndIp(ipSection)
						tupleItem["sip"] = startIp + "-" + endIp
					}
				} else if strings.Contains(ipSection, "-") {
					tupleItem["sip"] = ipSection
				} else { // 单个IP
					tupleItem["sip"] = ipSection + "-" + ipSection
				}
				tupleList = append(tupleList, tupleItem)
			}
		}
	} else if len(userInfos) > 0 && len(dstAddrArray) > 0 {
		for i := 0; i < len(userInfos); i++ {
			userDto := userInfos[i]
			ipSectionArray := userDto.IpAddress
			for j := 0; j < len(ipSectionArray); j++ {
				ipSection := ipSectionArray[j]
				var sip string
				if strings.Contains(ipSection, "/") {
					if strings.Contains(ipSection, ":") {
						startIp := utils.GetStartIp(ipSection)
						endIp := utils.GetEndIp(ipSection)
						sip = startIp + "-" + endIp
					}
				} else if strings.Contains(ipSection, "-") {
					sip = ipSection
				} else { // 单个IP
					sip = ipSection + "-" + ipSection
				}

				for z := 0; z < len(dstAddrArray); z++ {
					targetItem := dstAddrArray[z]
					tupleItem := map[string]string{}
					mapStr := map[string]string{}
					_ = json.Unmarshal([]byte(targetItem), mapStr)

					dipStart := mapStr["dipStart"]
					dipEnd := mapStr["dipEnd"]
					dport := mapStr["dport"]

					tupleItem["sip"] = sip
					if dipStart != "0" {
						tupleItem["dip"] = dipStart + "-" + dipEnd
					}
					if dport != "-1" {
						tupleItem["dport"] = dport
					}
					tupleList = append(tupleList, tupleItem)
				}
			}
		}
	}
	return tupleList
}

func getListByLinkIds(linkIds []string) (rest []modelConfiguration.DimLineInfo) {
	if err := global.ServiceDB.Model(modelConfiguration.DimLineInfo{}).Where("id in ?", linkIds).Find(&rest).Error; err != nil {
		global.Log.Error("查询DimLineInfo失败", zap.Error(err))
		return nil
	}
	return rest
}
