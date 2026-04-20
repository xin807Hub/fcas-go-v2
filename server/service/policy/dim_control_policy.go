package policy

import (
	"encoding/json"
	"errors"
	"fcas_server/global"
	modelConfiguration "fcas_server/model/configuration"
	"fcas_server/model/policy"
	"fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"sort"
	"strconv"
	"strings"
	"time"
)

type DimControlPolicyService struct {
}

func (DimControlPolicyService) List() (result []policy.DimControlPolicy, err error) {
	if err = global.ServiceDB.Model(&policy.DimControlPolicy{}).Find(&result).Error; err != nil {
		global.Log.Error("查询流控策略失败", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (DimControlPolicyService) GetPolicyIdList() (idList []int, err error) {
	if err = global.ServiceDB.Model(&policy.DimControlPolicy{}).Select("id").Scan(&idList).Error; err != nil {
		global.Log.Error("查询流控策略ID失败", zap.Error(err))
		return nil, err
	}
	return idList, nil
}

func (DimControlPolicyService) PageControlPolicy(policyReq policy.DimControlPolicyReq) (result []policy.DimControlPolicyResp, total int64, err error) {
	limit := policyReq.Limit
	offset := policyReq.Limit * (policyReq.Page - 1)

	sql := `select %s from (
            select a.*, c.crowd_name as user_crowd_name from (
                select a.*, g.group_name as user_crowd_group_name from (
                    select c.*, b.app_name from (
                        select w.*, a.app_type_name from dim_control_policy w
                        left join (select distinct app_type_id, app_type_name from dim_app_classify) a
                          on w.app_type_id = a.app_type_id
                    ) c
                    left join (select distinct app_type_id, app_id, app_name from dim_app_classify) b
                      on c.app_id = b.app_id and c.app_type_id = b.app_type_id
                ) a left join dim_user_crowd_group g on a.user_crowd_group_id = g.id
            ) a left join dim_user_crowd c on a.user_crowd_id = c.id
        ) a left join dim_user_info u on a.user_id = u.id where 1=1`

	params := make([]interface{}, 0)

	if policyReq.PolicyName != "" {
		sql += " and a.name like ?"
		params = append(params, "%"+policyReq.PolicyName+"%")
	}
	if policyReq.UserType != "" || policyReq.UerType != "" {
		userType := policyReq.UserType
		if userType == "" {
			userType = policyReq.UerType
		}
		sql += " and a.user_type = ?"
		params = append(params, userType)
	}
	if policyReq.UserCrowdGroupId != "" || policyReq.UerCrowdGroupId != "" {
		groupID := policyReq.UserCrowdGroupId
		if groupID == "" {
			groupID = policyReq.UerCrowdGroupId
		}
		sql += " and a.user_crowd_group_id = ?"
		params = append(params, groupID)
	}
	if policyReq.UserCrowdId != "" || policyReq.UerCrowdId != "" {
		crowdID := policyReq.UserCrowdId
		if crowdID == "" {
			crowdID = policyReq.UerCrowdId
		}
		sql += " and a.user_crowd_id = ?"
		params = append(params, crowdID)
	}
	if policyReq.UserId != "" || policyReq.UerId != "" {
		userID := policyReq.UserId
		if userID == "" {
			userID = policyReq.UerId
		}
		sql += " and a.user_id = ?"
		params = append(params, userID)
	}
	if policyReq.AppTypeId != 0 {
		sql += " and a.app_type_id = ?"
		params = append(params, policyReq.AppTypeId)
	}
	if policyReq.AppId != 0 {
		sql += " and a.app_id = ?"
		params = append(params, policyReq.AppId)
	}

	sqlCount := fmt.Sprintf(sql, "count(*)")
	if err = global.ServiceDB.Raw(sqlCount, params...).Scan(&total).Error; err != nil {
		global.Log.Error("统计流控策略失败", zap.Error(err))
		return nil, 0, err
	}

	sqlSelect := fmt.Sprintf(sql, "a.*, u.user_name")
	sqlSelect += " limit ? offset ?"
	params = append(params, limit, offset)

	if err = global.ServiceDB.Raw(sqlSelect, params...).Scan(&result).Error; err != nil {
		global.Log.Error("分页查询流控策略失败", zap.Error(err))
		return nil, 0, err
	}

	for i := 0; i < len(result); i++ {
		dto := result[i]
		var policyLog []policy.DimControlPolicyLog
		if err := global.ServiceDB.Model(&policy.DimControlPolicyLog{}).
			Where("policy_id = ?", dto.Id).
			Order("record_time desc").
			Find(&policyLog).Error; err != nil {
			global.Log.Error("查询流控策略日志失败", zap.Error(err))
			continue
		}
		if len(policyLog) > 1 {
			times := int((policyLog[0].RecordTime.UnixMilli() - policyLog[1].RecordTime.UnixMilli()) / 1000)
			result[i].UpTrafficSpeed = calculating(policyLog[0].UpTraffic, policyLog[1].UpTraffic, times)
			result[i].DnTrafficSpeed = calculating(policyLog[0].DnTraffic, policyLog[1].DnTraffic, times)
			result[i].UpPassSpeed = calculating(policyLog[0].UpPass, policyLog[1].UpPass, times)
			result[i].DnPassSpeed = calculating(policyLog[0].DnPass, policyLog[1].DnPass, times)
			result[i].UpDiscardSpeed = calculating(policyLog[0].UpDiscard, policyLog[1].UpDiscard, times)
			result[i].DnDisCardSpeed = calculating(policyLog[0].DnDiscard, policyLog[1].DnDiscard, times)
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
	if err := d.validateControlPolicy(controlPolicy); err != nil {
		return err
	}

	now := time.Now()
	isEffective, err := ifTimeEffect(controlPolicy, now)
	if err != nil {
		global.Log.Error("判断策略生效状态失败", zap.Error(err))
		return err
	}

	if controlPolicy.Id == 0 {
		exists, err := d.isNameExists(controlPolicy.Name, 0)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("策略名称已存在")
		}

		controlPolicy.CreateTime = now.Format(global.DateTimeLayout)
		controlPolicy.Status = global.UnBind
		if isEffective {
			controlPolicy.Status = global.Bind
		}

		if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Create(&controlPolicy).Error; err != nil {
			global.Log.Error("新增流控策略失败", zap.Error(err))
			return err
		}

		if controlPolicy.Status == global.Bind {
			policyInfoStr := BuildPolicyJson(controlPolicy, global.Bind)
			global.Log.Info("流控策略绑定: " + policyInfoStr)
			utils.SendMessage(policyInfoStr, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)
		}
		return nil
	}

	oldPolicy, err := d.mustGetByPolicyId(controlPolicy.Id)
	if err != nil {
		return err
	}

	if oldPolicy.Name != controlPolicy.Name {
		exists, err := d.isNameExists(controlPolicy.Name, controlPolicy.Id)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("策略名称已存在")
		}
	}

	controlPolicy.CreateTime = oldPolicy.CreateTime
	controlPolicy.Status = oldPolicy.Status
	if isEffective {
		controlPolicy.Status = global.Bind
	}

	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id = ?", controlPolicy.Id).Updates(map[string]interface{}{
		"name":                controlPolicy.Name,
		"user_type":           nullableInt(controlPolicy.UserType),
		"user_crowd_group_id": nullableInt(controlPolicy.UserCrowdGroupId),
		"user_crowd_id":       nullableInt(controlPolicy.UserCrowdId),
		"user_id":             nullableInt(controlPolicy.UserId),
		"ul_flow_rate":        nullableInt(controlPolicy.UlFlowRate),
		"dl_flow_rate":        nullableInt(controlPolicy.DlFlowRate),
		"start_time":          controlPolicy.StartTime,
		"end_time":            controlPolicy.EndTime,
		"remark":              controlPolicy.Remark,
		"flow_ctrl_type":      controlPolicy.FlowCtrlType,
		"app_type_id":         nullableInt(controlPolicy.AppTypeId),
		"app_id":              nullableInt(controlPolicy.AppId),
		"ip_data":             controlPolicy.IpData,
		"period_type":         nullablePointerInt(controlPolicy.PeriodType),
		"policy_period":       controlPolicy.PolicyPeriod,
		"link_ids":            controlPolicy.LinkIds,
		"status":              controlPolicy.Status,
	}).Error; err != nil {
		global.Log.Error("更新流控策略失败", zap.Error(err))
		return err
	}

	if controlPolicy.Status == global.Bind {
		policyInfoStr := BuildPolicyJson(controlPolicy, global.Bind)
		global.Log.Info("流控策略绑定: " + policyInfoStr)
		utils.SendMessage(policyInfoStr, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)
	}

	return nil
}

func ifTimeEffect(controlPolicy policy.DimControlPolicy, now time.Time) (bool, error) {
	start, err := time.ParseInLocation(global.DateTimeLayout, controlPolicy.StartTime, now.Location())
	if err != nil {
		return false, err
	}
	end, err := time.ParseInLocation(global.DateTimeLayout, controlPolicy.EndTime, now.Location())
	if err != nil {
		return false, err
	}
	if end.Before(start) {
		return false, errors.New("结束时间不能早于开始时间")
	}
	return (now.Equal(start) || now.After(start)) && (now.Equal(end) || now.Before(end)), nil
}

func (d DimControlPolicyService) InfoControlPolicy(id int) (result *policy.DimControlPolicy) {
	var data policy.DimControlPolicy
	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id = ?", id).First(&data).Error; err != nil {
		global.Log.Error("查询流控策略详情失败", zap.Error(err))
		return nil
	}
	return &data
}

func (d DimControlPolicyService) getByPolicyId(policyId int) (result *policy.DimControlPolicy) {
	data, err := d.mustGetByPolicyId(policyId)
	if err != nil {
		global.Log.Error("根据策略ID查询流控策略失败", zap.Error(err))
		return nil
	}
	return data
}

func (d DimControlPolicyService) mustGetByPolicyId(policyId int) (*policy.DimControlPolicy, error) {
	var data policy.DimControlPolicy
	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id = ?", policyId).First(&data).Error; err != nil {
		global.Log.Error("根据策略ID查询流控策略失败", zap.Error(err))
		return nil, err
	}
	return &data, nil
}

func (d DimControlPolicyService) DeleteControlPolicy(ids []int) error {
	var result []policy.DimControlPolicy
	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id in ?", ids).Find(&result).Error; err != nil {
		global.Log.Error("查询待删除流控策略失败", zap.Error(err))
		return err
	}

	for _, controlPolicy := range result {
		if strings.EqualFold(controlPolicy.Status, global.Bind) {
			policyInfo := BuildPolicyJson(controlPolicy, global.UnBind)
			global.Log.Info(strconv.Itoa(controlPolicy.Id) + " 流控策略解绑: " + policyInfo)
			utils.SendMessage(policyInfo, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)
		}
	}

	if err := (DimControlPolicyActionService{}).DeleteByPolicyIDs(ids); err != nil {
		return err
	}

	if err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("id in ?", ids).Delete(policy.DimControlPolicy{}).Error; err != nil {
		global.Log.Error("删除流控策略失败", zap.Error(err))
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

	timeInfo := struct {
		StartTime  string      `json:"start_time"`
		EndTime    string      `json:"end_time"`
		Period     string      `json:"period"`
		PeriodInfo interface{} `json:"period_info"`
	}{
		StartTime: normalizeDateTimeString(controlPolicy.StartTime),
		EndTime:   normalizeDateTimeString(controlPolicy.EndTime),
	}

	periodInfoList := make([]interface{}, 0)
	if controlPolicy.PolicyPeriod == "" {
		timeInfo.Period = "all"
		periodInfoList = append(periodInfoList, map[string]string{
			"start_period": "0",
			"end_period":   "0",
		})
	} else {
		if controlPolicy.PeriodType != nil && *controlPolicy.PeriodType == 2 {
			timeInfo.Period = "week"
		} else {
			timeInfo.Period = "day"
		}
		periodArr := strings.Split(controlPolicy.PolicyPeriod, ",")
		for _, periodItem := range periodArr {
			parts := strings.Split(periodItem, "-")
			if len(parts) != 2 {
				continue
			}
			periodInfoList = append(periodInfoList, map[string]string{
				"start_period": parts[0],
				"end_period":   parts[1],
			})
		}
		if len(periodInfoList) == 0 {
			timeInfo.Period = "all"
			periodInfoList = append(periodInfoList, map[string]string{
				"start_period": "0",
				"end_period":   "0",
			})
		}
	}
	timeInfo.PeriodInfo = periodInfoList
	dpiPolicy.TimeInfo = timeInfo

	control := struct {
		ThresholdUpIds int   `json:"threshold_up_ids"`
		ThresholdDnIds int   `json:"threshold_dn_ids"`
		LinkIds        []int `json:"link_ids"`
	}{
		ThresholdUpIds: thresholdValue(controlPolicy.UlFlowRate),
		ThresholdDnIds: thresholdValue(controlPolicy.DlFlowRate),
		LinkIds:        resolveLinkVlanIDs(controlPolicy.LinkIds),
	}
	dpiPolicy.ControlPolicy = control

	appInfoList := make([]interface{}, 0)
	if controlPolicy.FlowCtrlType == 1 && controlPolicy.AppTypeId != 0 {
		appID := 0
		if controlPolicy.AppId != 0 {
			appID = controlPolicy.AppId - controlPolicy.AppTypeId*10000
			if appID < 0 {
				appID = controlPolicy.AppId
			}
		}
		appInfoList = append(appInfoList, map[string]int{
			"app_type": controlPolicy.AppTypeId,
			"app_id":   appID,
		})
	}

	userInfoList := getUserInfoList(controlPolicy)
	tupleList := buildTupleJsonList(userInfoList, controlPolicy.IpData)

	featureList := map[string]interface{}{
		"combined": len(appInfoList) > 0 && len(tupleList) > 0,
		"tuple":    tupleList,
		"app_info": appInfoList,
	}

	dpiPolicy.FeatureList = featureList
	dpiPolicy.BindObj = bindFlag

	policyJSON, _ := json.Marshal(dpiPolicy)
	return string(policyJSON)
}

func getUserInfoList(controlPolicy policy.DimControlPolicy) (userInfos []modelConfiguration.DimUserInfo) {
	if controlPolicy.UserId != 0 {
		if err := global.ServiceDB.Model(&modelConfiguration.DimUserInfo{}).Where("id = ?", controlPolicy.UserId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据用户ID获取用户失败", zap.Error(err))
			return nil
		}
		return userInfos
	} else if controlPolicy.UserCrowdId != 0 {
		sql := `SELECT tb1.*
        FROM dim_user_info tb1,
             dim_user_crowd tb2,
             dim_user_crowd_relation tb3
        WHERE tb1.id = tb3.user_id
          AND tb2.id = tb3.crowd_id
          AND tb2.id = ?`

		if err := global.ServiceDB.Raw(sql, controlPolicy.UserCrowdId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据用户群ID获取用户失败", zap.Error(err))
			return nil
		}
		return userInfos
	} else if controlPolicy.UserCrowdGroupId != 0 {
		sql := `select distinct u.*
        from dim_user_info u,
             dim_user_crowd uc,
             dim_user_crowd_relation ur,
             dim_user_crowd_group_relation ugr
        where u.id = ur.user_id
          and uc.id = ur.crowd_id
          and uc.id = ugr.crowd_id
          and ugr.group_id = ?`
		if err := global.ServiceDB.Raw(sql, controlPolicy.UserCrowdGroupId).Find(&userInfos).Error; err != nil {
			global.Log.Error("根据用户群组ID获取用户失败", zap.Error(err))
			return nil
		}
		return userInfos
	}
	return userInfos
}

func buildTupleJsonList(userInfos []modelConfiguration.DimUserInfo, dstAddrs []policy.IpData) []interface{} {
	var tupleList = make([]interface{}, 0)

	transformUserAddrToSip := func(addr string) string {
		if strings.Contains(addr, "/") {
			if strings.Contains(addr, ":") {
				return ""
			}
			return utils.GetStartIp(addr) + "-" + utils.GetEndIp(addr)
		}
		if strings.Contains(addr, "-") {
			return addr
		}
		return addr + "-" + addr
	}

	switch {
	case len(userInfos) == 0 && len(dstAddrs) > 0:
		for _, addr := range dstAddrs {
			tupleItem := map[string]string{}
			if addr.StartIP != "" {
				tupleItem["dip"] = addr.StartIP + "-" + addr.EndIP
			}
			if addr.DstPort != "" {
				tupleItem["dport"] = addr.DstPort
			}
			tupleList = append(tupleList, tupleItem)
		}

	case len(userInfos) > 0 && len(dstAddrs) == 0:
		for _, user := range userInfos {
			for _, addr := range user.IpAddress {
				sip := transformUserAddrToSip(addr)
				if sip == "" {
					continue
				}
				tupleList = append(tupleList, map[string]string{
					"sip": sip,
				})
			}
		}

	case len(userInfos) > 0 && len(dstAddrs) > 0:
		for _, user := range userInfos {
			for _, userAddr := range user.IpAddress {
				sip := transformUserAddrToSip(userAddr)
				if sip == "" {
					continue
				}
				for _, dstAddr := range dstAddrs {
					tupleItem := map[string]string{
						"sip": sip,
					}
					if dstAddr.StartIP != "" {
						tupleItem["dip"] = dstAddr.StartIP + "-" + dstAddr.EndIP
					}
					if dstAddr.DstPort != "" {
						tupleItem["dport"] = dstAddr.DstPort
					}
					tupleList = append(tupleList, tupleItem)
				}
			}
		}
	}

	return tupleList
}

func (d DimControlPolicyService) validateControlPolicy(controlPolicy policy.DimControlPolicy) error {
	if strings.TrimSpace(controlPolicy.Name) == "" {
		return errors.New("策略名称不能为空")
	}
	if strings.TrimSpace(controlPolicy.LinkIds) == "" {
		return errors.New("链路不能为空")
	}
	if controlPolicy.UlFlowRate == 0 && controlPolicy.DlFlowRate == 0 {
		return errors.New("上行限速与下行限速至少填写一项")
	}
	if hasIPv6InPolicy(controlPolicy.IpData) {
		return errors.New("暂不支持IPv6")
	}
	return nil
}

func hasIPv6InPolicy(ipData []policy.IpData) bool {
	for _, item := range ipData {
		if strings.Contains(item.StartIP, ":") || strings.Contains(item.EndIP, ":") {
			return true
		}
	}
	return false
}

func (d DimControlPolicyService) isNameExists(name string, excludeID int) (bool, error) {
	tx := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("name = ?", name)
	if excludeID != 0 {
		tx = tx.Where("id <> ?", excludeID)
	}
	var count int64
	if err := tx.Count(&count).Error; err != nil {
		global.Log.Error("校验流控策略名称失败", zap.Error(err))
		return false, err
	}
	return count > 0, nil
}

func thresholdValue(value int) int {
	if value <= 0 {
		return -1
	}
	return value * 1024
}

func resolveLinkVlanIDs(linkIDs string) []int {
	rawIDs := parseCSVInts(linkIDs)
	if len(rawIDs) == 0 {
		return []int{}
	}

	var lines []modelConfiguration.DimLineInfo
	if err := global.ServiceDB.Model(&modelConfiguration.DimLineInfo{}).
		Where("id in ? OR line_vlan in ?", rawIDs, rawIDs).
		Find(&lines).Error; err != nil {
		global.Log.Error("解析链路VLAN失败", zap.Error(err))
		return rawIDs
	}

	idToVlan := make(map[int]int, len(lines))
	vlanSet := make(map[int]struct{}, len(lines))
	for _, line := range lines {
		idToVlan[line.ID] = line.LineVlan
		vlanSet[line.LineVlan] = struct{}{}
	}

	result := make([]int, 0, len(rawIDs))
	seen := make(map[int]struct{}, len(rawIDs))
	for _, id := range rawIDs {
		if vlan, ok := idToVlan[id]; ok {
			if _, exists := seen[vlan]; !exists {
				result = append(result, vlan)
				seen[vlan] = struct{}{}
			}
			continue
		}
		if _, ok := vlanSet[id]; ok {
			if _, exists := seen[id]; !exists {
				result = append(result, id)
				seen[id] = struct{}{}
			}
		}
	}
	sort.Ints(result)
	if len(result) == 0 {
		return rawIDs
	}
	return result
}

func parseCSVInts(value string) []int {
	parts := strings.Split(value, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		number, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		result = append(result, number)
	}
	return result
}

func normalizeDateTimeString(value string) string {
	if value == "" {
		return value
	}
	return strings.ReplaceAll(strings.ReplaceAll(value, "T", " "), "+08:00", "")
}

func nullableInt(value int) interface{} {
	if value == 0 {
		return nil
	}
	return value
}

func nullablePointerInt(value *int) interface{} {
	if value == nil || *value == 0 {
		return nil
	}
	return *value
}
