package policy

import (
	"encoding/json"
	"fcas_server/global"
	"fcas_server/model/configuration"
	"fcas_server/model/policy"
	"fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type DimWhitePolicyService struct {
}

func (DimWhitePolicyService) PageWhitePolicy(policyReq policy.DimWhitePolicyReq) (result []policy.DimWhitePolicyResp, total int64, err error) {

	limit := policyReq.Limit
	offset := policyReq.Limit * (policyReq.Page - 1)

	sql := `select %s from (
            select a.*, c.crowd_name as user_crowd_name from (
                select a.*, g.group_name as user_crowd_group_name from (
                    select c.*, b.app_name from (
                        select w.*, a.app_type_name from dim_white_policy w left join (select DISTINCT app_type_id, app_type_name from dim_app_classify) a on w.app_type_id = a.app_type_id
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
	if policyReq.AppId != 0 {
		sql += " and a.app_id = ? "
		params = append(params, policyReq.AppId)
	}

	sqlCount := fmt.Sprintf(sql, "count(*)")

	if err = global.ServiceDB.Raw(sqlCount, params...).Scan(&total).Error; err != nil {
		global.Log.Error("策略配置count查询失败", zap.Error(err))
		return nil, 0, err
	}

	sqlSelect := fmt.Sprintf(sql, "a.*, u.user_name")

	sqlSelect += " limit ? offset ?"
	params = append(params, limit)
	params = append(params, offset)

	if err := global.ServiceDB.Raw(sqlSelect, params...).Scan(&result).Error; err != nil {
		global.Log.Error("策略配置page信息查询失败", zap.Error(err))
		return nil, 0, err
	}
	return result, total, nil

}

func (d DimWhitePolicyService) SaveOrUpdateWhitePolicy(whitePolicy policy.DimWhitePolicy) error {
	if whitePolicy.Id == 0 {
		whitePolicy.CreateTime = time.Now()
		if err := global.ServiceDB.Model(&policy.DimWhitePolicy{}).Save(&whitePolicy).Error; err != nil {
			global.Log.Error("策略配置save失败", zap.Error(err))
			return err
		}
	} else {
		if err := global.ServiceDB.Model(&policy.DimWhitePolicy{}).Where("id = ?", whitePolicy.Id).Updates(&whitePolicy).Error; err != nil {
			global.Log.Error("策略配置update失败", zap.Error(err))
			return err
		}
	}
	//保存策略后, 发送绑定策略
	policyInfo := d.buildPolicyJson(whitePolicy, global.Bind)
	global.Log.Info("策略绑定：" + policyInfo)
	//utils.SendMessage(policyInfo, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.WhitePolicyMsgType), global.CONFIG.Policy.Dir)
	return nil
}

func (d DimWhitePolicyService) GetById(id int) (result *policy.DimWhitePolicyReq) {
	if err := global.ServiceDB.Model(&policy.DimWhitePolicyReq{}).Where("id = ?", id).Find(result).Error; err != nil {
		global.Log.Error("策略白名单配置信息查询失败", zap.Error(err))
		return nil
	}
	return result
}

func (d DimWhitePolicyService) DeleteWhitePolicy(ids []int) error {
	var result []policy.DimWhitePolicy
	if err := global.ServiceDB.Model(&policy.DimWhitePolicy{}).Where("id IN ?", ids).Find(&result).Error; err != nil {
		global.Log.Error("查询策略配置失败", zap.Error(err))
		return err
	}

	for _, whitePolicy := range result {
		// 发送解绑策略
		policyInfo := d.buildPolicyJson(whitePolicy, global.UnBind)
		global.Log.Info(strconv.Itoa(whitePolicy.Id) + " 策略解绑：" + policyInfo)
		//utils.SendMessage(policyInfo, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.WhitePolicyMsgType), global.CONFIG.Policy.Dir)
	}

	if err := global.ServiceDB.Model(&policy.DimWhitePolicy{}).Where("id in ?", ids).Delete(policy.DimWhitePolicy{}).Error; err != nil {
		global.Log.Error("删除策略配置失败", zap.Error(err))
		return err
	}
	return nil
}

func (d DimWhitePolicyService) buildPolicyJson(whitePolicy policy.DimWhitePolicy, bindFlag string) string {
	policy := struct {
		MsgType       int         `json:"msg_type"`
		PolicyId      int         `json:"policy_id"`
		TimeInfo      interface{} `json:"time_info"`
		ControlPolicy interface{} `json:"control_policy"`
		FeatureList   interface{} `json:"feature_list"`
		BindObj       interface{} `json:"bind_obj"`
	}{
		MsgType:  policy.WhitePolicyMsgType,
		PolicyId: whitePolicy.Id,
	}
	policy.TimeInfo = struct { // time_info
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}{
		StartTime: strings.Replace(strings.Replace(whitePolicy.StartTime, "T", " ", -1), "+08:00", "", -1),
		EndTime:   strings.Replace(strings.Replace(whitePolicy.EndTime, "T", " ", -1), "+08:00", "", -1),
	}
	policy.ControlPolicy = struct { // control_policy
		UlTos int `json:"ul_tos"`
		DlTos int `json:"dl_tos"`
	}{
		UlTos: whitePolicy.UlTos,
		DlTos: whitePolicy.DlTos,
	}
	// 构造feature_list 构造应用大小类的json
	appInfoList := append([]interface{}{}, 0)
	if whitePolicy.AppTypeId != 0 {
		appInfo := struct {
			AppType int `json:"app_type"`
			AppId   int `json:"app_id"`
		}{
			AppType: whitePolicy.AppTypeId,
			AppId:   whitePolicy.AppId,
		}
		appInfoList = append(appInfoList, appInfo)
	}
	dimUserInfoList := d.getUserInfoList(whitePolicy)
	tupList := d.buildUserJsonList(dimUserInfoList)

	featureList := struct {
		Combined bool          `json:"combined"`
		AppInfo  []interface{} `json:"app_info"`
		Tuple    []interface{} `json:"tuple"`
	}{}
	if len(tupList) > 0 {
		featureList.Combined = true
	} else {
		featureList.Combined = false
	}
	featureList.AppInfo = appInfoList
	featureList.Tuple = tupList

	policy.FeatureList = featureList
	policy.BindObj = bindFlag
	jsonBt, _ := json.Marshal(policy)

	return string(jsonBt)
}

func (DimWhitePolicyService) getUserInfoList(whitePolicy policy.DimWhitePolicy) (userInfos []configuration.DimUserInfo) {
	// 构造user_list
	if whitePolicy.UserId != 0 {
		if err := global.ServiceDB.Model(&configuration.DimUserInfo{}).Where("id = ?", whitePolicy.UserId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据ID获取信息失败", zap.Error(err))
			return nil
		}
		return userInfos
	} else if whitePolicy.UserCrowdId != 0 {
		sql := `SELECT
            tb1.*
        FROM
            dim_user_info tb1,
            dim_user_crowd tb2,
            dim_user_crowd_relation tb3
        where tb1.id = tb3.user_id
        AND tb2.id = tb3.crowd_id AND tb2.id = ?`

		if err := global.ServiceDB.Raw(sql, whitePolicy.UserCrowdId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据CrowdId获取用户信息失败", zap.Error(err))
			return nil
		}
		return userInfos
	} else if whitePolicy.UserCrowdGroupId != 0 {
		sql := `select distinct u.* from
            dim_user_info u,dim_user_crowd uc,dim_user_crowd_relation ur, dim_user_crowd_group_relation ugr
        where u.id = ur.user_id and uc.id = ur.crowd_id and uc.id = ugr.crowd_id and ugr.group_id = ? `
		if err := global.ServiceDB.Raw(sql, whitePolicy.UserCrowdGroupId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据UserCrowdGroupId获取用户信息失败", zap.Error(err))
			return nil
		}
		return userInfos
	}
	return userInfos
}

func (DimWhitePolicyService) buildUserJsonList(userInfos []configuration.DimUserInfo) []interface{} {
	tupList := make([]interface{}, 0)
	for i := 0; i < len(userInfos); i++ {
		dto := userInfos[i]
		for _, ipVal := range dto.IpAddress {
			if strings.Contains(ipVal, "/") {
				if strings.Contains(ipVal, ":") {
				} else {
					startIp := utils.GetStartIp(ipVal)
					endIp := utils.GetEndIp(ipVal)
					rest := struct {
						Sip string `json:"sip"`
					}{}
					rest.Sip = startIp + `-` + endIp
					tupList = append(tupList, rest)
				}
			} else {
				rest := struct {
					Sip string `json:"sip"`
				}{}
				rest.Sip = ipVal
				tupList = append(tupList, rest)
			}
		}
	}
	return tupList
}
