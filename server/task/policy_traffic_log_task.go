package task

import (
	"encoding/json"
	"fcas_server/global"
	policyModel "fcas_server/model/policy"
	"fcas_server/service/policy"
	"fcas_server/utils"
	"io"
	"strconv"
	"time"
)

const shuntPort = 8080

var (
	dimControlPolicyService = policy.DimControlPolicyService{}
)

func RunPolicyTrafficLogTask() {
	date := time.Time{}
	policyActionList, err := dimControlPolicyActionService.List()
	if err != nil {
		global.Log.Error("ControlPolicyAction查询失败：" + err.Error())
		return
	}

	controlPolicyList, err := dimControlPolicyService.List()
	if err != nil {
		global.Log.Error("ControlPolicy查询失败：" + err.Error())
		return
	}

	policyIdList := make([]int, 0)
	for _, item := range controlPolicyList {
		policyIdList = append(policyIdList, item.Id)
	}
	upList := make([]int, 0)
	dnList := make([]int, 0)

	acyionToPolicy := map[int]int{}
	ipToAction := map[string][]int{}
	resultMap := map[int]string{}

	resultDtoMap := map[int]*policyModel.DimControlPolicyLog{}

	for i := 0; i < len(policyActionList); i++ {
		entity := policyActionList[0]
		if !utils.Contains(policyIdList, entity.PolicyId) {
			continue
		}
		upList = append(upList, entity.UploadActionId)
		dnList = append(dnList, entity.DownloadActionId)
		acyionToPolicy[entity.DownloadActionId] = entity.PolicyId
		acyionToPolicy[entity.UploadActionId] = entity.PolicyId

		if ipToAction[entity.ShuntIp] != nil {
			ints := ipToAction[entity.ShuntIp]
			ints = append(ints, entity.UploadActionId)
			ints = append(ints, entity.DownloadActionId)
			ipToAction[entity.ShuntIp] = ints
		} else {
			ints := make([]int, 0)
			ints = append(ints, entity.UploadActionId)
			ints = append(ints, entity.DownloadActionId)
			ipToAction[entity.ShuntIp] = ints
		}
	}

	for key, value := range ipToAction {
		url := "http://" + key + ":" + strconv.Itoa(shuntPort) + "/controlAction-get"
		postData := map[string]interface{}{}
		postData["actionIds"] = value
		resp, err := utils.HttpRequest(url, "post", nil, nil, postData)
		if err == nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				global.Log.Error("分流器接口数据获取失败，分流器IP：" + key)
				continue
			}
			bodyRest := map[string]string{}
			json.Unmarshal(body, &bodyRest)
			dataArr := bodyRest["data"]
			if len(dataArr) > 0 {
				global.Log.Info("查询actionId对应流量数据成功，分流器IP：" + string(body))
				var dataMapArr []map[string]interface{}
				json.Unmarshal([]byte(dataArr), &dataMapArr)
				for j := 0; j < len(dataMapArr); j++ {
					if dataMapArr[j]["code"] == 1 {
						resultMap[dataMapArr[j]["actionId"].(int)] = resultMap[j]
					}
				}
			}
		} else {
			global.Log.Error("查询actionId对于的流量数据失败，分流器IP：" + key)
		}
	}

	for actionId, value := range resultMap {
		jsonMap := map[string]interface{}{}
		json.Unmarshal([]byte(value), &jsonMap)
		allTraffic := jsonMap["allTrafficData"]
		throughTraffic := jsonMap["throughTrafficData"]

		policyId := acyionToPolicy[actionId]

		if utils.Contains(upList, actionId) {
			if resultDtoMap[policyId] != nil {
				controlPolicyLog := resultDtoMap[policyId]
				controlPolicyLog.UpTraffic = allTraffic.(int) + controlPolicyLog.UpTraffic
				controlPolicyLog.UpPass = allTraffic.(int) + controlPolicyLog.UpPass
				controlPolicyLog.UpDiscard = throughTraffic.(int) + controlPolicyLog.UpDiscard
				resultDtoMap[policyId] = controlPolicyLog
			} else {
				controlPolicy := policyModel.DimControlPolicyLog{}
				controlPolicy.PolicyId = policyId
				controlPolicy.UpTraffic = allTraffic.(int)
				controlPolicy.UpPass = throughTraffic.(int)
				controlPolicy.UpDiscard = allTraffic.(int) - throughTraffic.(int)
				resultDtoMap[policyId] = &controlPolicy
			}
		} else if utils.Contains(dnList, actionId) {
			if resultDtoMap[policyId] != nil {
				controlPolicyLog := resultDtoMap[policyId]
				controlPolicyLog.DnTraffic = allTraffic.(int) + controlPolicyLog.DnTraffic
				controlPolicyLog.DnPass = throughTraffic.(int) + controlPolicyLog.DnPass
				controlPolicyLog.DnDiscard = allTraffic.(int) - throughTraffic.(int) + controlPolicyLog.DnDiscard
				resultDtoMap[policyId] = controlPolicyLog
			} else {
				controlPolicy := policyModel.DimControlPolicyLog{}
				controlPolicy.PolicyId = policyId
				controlPolicy.DnTraffic = allTraffic.(int)
				controlPolicy.DnPass = throughTraffic.(int)
				controlPolicy.DnDiscard = allTraffic.(int) - throughTraffic.(int)
				resultDtoMap[policyId] = &controlPolicy
			}
		}
	}

	logEntityArr := make([]policyModel.DimControlPolicyLog, 0)
	for _, valueDto := range resultDtoMap {
		valueDto.RecordTime = date
		logEntityArr = append(logEntityArr, *valueDto)
	}
	if len(logEntityArr) > 0 {
		if err := global.ServiceDB.Model(&policyModel.DimControlPolicyLog{}).CreateInBatches(logEntityArr, 10000).Error; err != nil {
			global.Log.Error("policyLog日志入库失败：" + err.Error())
		}
	}
}
