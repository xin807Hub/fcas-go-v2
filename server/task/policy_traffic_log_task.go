package task

import (
	"encoding/json"
	"fcas_server/global"
	policyModel "fcas_server/model/policy"
	"fcas_server/service/policy"
	"fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var dimControlPolicyService = policy.DimControlPolicyService{}

// TrafficResponse 定义响应数据结构体
type TrafficResponse struct {
	Data []TrafficData `json:"data"`
}

type TrafficData struct {
	Code               int    `json:"code"`
	Msg                string `json:"msg"`
	ActionID           string `json:"actionId"`
	AllTrafficData     int    `json:"allTrafficData"`
	ThroughTrafficData int    `json:"throughTrafficData"`
}

func RunPolicyTrafficLogTask() {
	policyIds, err := dimControlPolicyService.GetPolicyIdList()
	if err != nil {
		global.Log.Error("策略查询失败：" + err.Error())
		return
	}
	if len(policyIds) < 1 {
		global.Log.Info("暂无策略, 退出")
		return
	}

	policyActionList, err := dimControlPolicyActionService.List()
	if err != nil {
		global.Log.Error("Action查询失败：" + err.Error())
		return
	}
	if len(policyActionList) < 1 {
		global.Log.Info("暂无action, 退出")
		return
	}

	upActionIdArr, dnActionIdArr, actionIdPolicyIdMap, shuntIpActionIdsMap := initData(policyActionList, policyIds)

	// 以分流器分组，逐个分流拉取最新的数据, 将每个actionId的结果保存在map中
	resultMap := postShuntAndGetResp(shuntIpActionIdsMap)

	// 将结果持久化
	err = ProcessAndStoreTrafficData(resultMap, actionIdPolicyIdMap, upActionIdArr, dnActionIdArr)
	if err != nil {
		global.Log.Error("处理响应的结果集并批量入库失败", zap.Error(err))
	}
}

func initData(policyActionList []policyModel.DimControlPolicyAction, policyIdList []int) ([]int, []int, map[int]int, map[string][]int) {
	upActionIdArr := make([]int, 0)
	dnActionIdArr := make([]int, 0)
	actionIdPolicyIdMap := make(map[int]int)
	shuntIpActionIdsMap := make(map[string][]int)

	for i := 0; i < len(policyActionList); i++ {
		action := policyActionList[0]
		if !utils.Contains(policyIdList, action.PolicyId) {
			global.Log.Info(fmt.Sprintf("该action所对应的策略不存在，actionId=%d, policyId=%d, 退出", action.Id, action.PolicyId))
			continue
		}
		if action.UploadActionId == 0 {
			global.Log.Info(fmt.Sprintf("该action所对应的upload_action_id=0，actionId=%d, policyId=%d, 退出", action.Id, action.PolicyId))
			continue
		}
		if action.DownloadActionId == 0 {
			global.Log.Info(fmt.Sprintf("该action所对应的download_action_id=0，actionId=%d, policyId=%d, 退出", action.Id, action.PolicyId))
			continue
		}
		if action.UploadDeviceId == 0 {
			global.Log.Info(fmt.Sprintf("该action所对应的upload_device_id=0，actionId=%d, policyId=%d, 退出", action.Id, action.PolicyId))
			continue
		}

		upActionIdArr = append(upActionIdArr, action.UploadActionId)
		dnActionIdArr = append(dnActionIdArr, action.DownloadActionId)
		actionIdPolicyIdMap[action.DownloadActionId] = action.PolicyId
		actionIdPolicyIdMap[action.UploadActionId] = action.PolicyId
		if actionIds, exists := shuntIpActionIdsMap[action.ShuntIp]; exists {
			actionIds = append(actionIds, action.UploadActionId)
			actionIds = append(actionIds, action.DownloadActionId)
			shuntIpActionIdsMap[action.ShuntIp] = actionIds
		} else {
			initActionIds := make([]int, 0)
			initActionIds = append(initActionIds, action.UploadActionId)
			initActionIds = append(initActionIds, action.DownloadActionId)
			shuntIpActionIdsMap[action.ShuntIp] = initActionIds
		}
	}
	return upActionIdArr, dnActionIdArr, actionIdPolicyIdMap, shuntIpActionIdsMap
}

func postShuntAndGetResp(shuntIpActionIdsMap map[string][]int) map[int]TrafficData {
	var shuntPort = global.CONFIG.Policy.ShuntPort
	resultMap := make(map[int]TrafficData)
	for shuntIp, actionIds := range shuntIpActionIdsMap {
		global.Log.Info(fmt.Sprintf("开始请求分流数据：shuntIp = %s, port = %d, actionIds= %v", shuntIp, shuntPort, actionIds))

		url := fmt.Sprintf("http://%s:%d/controlAction-get", shuntIp, shuntPort)
		postParam := map[string]interface{}{
			"actionIds": actionIds,
		}
		resp, httpErr := utils.HttpRequest(url, "post", nil, nil, postParam)
		if httpErr != nil {
			global.Log.Error(fmt.Sprintf("查询actionId对于的流量数据失败，分流器IP：%s, post参数= %v", shuntIp, postParam), zap.Error(httpErr))
			continue
		}

		var response TrafficResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			global.Log.Error("解析分流的响应数据失败", zap.Error(err))
			continue
		}
		global.Log.Info(fmt.Sprintf("查询actionId对应流量数据成功，body：%v", response))

		if len(response.Data) > 0 {
			for _, data := range response.Data {
				if data.Code == 1 {
					actionId, atoiErr := strconv.Atoi(data.ActionID)
					if atoiErr != nil {
						global.Log.Error(fmt.Sprintf("该actionId转成数字失败，data = %v", data))
						continue
					}
					resultMap[actionId] = data
				} else {
					global.Log.Error(fmt.Sprintf("该action的响应失败，data = %v", data))
					continue
				}
			}
		} else {
			global.Log.Info(fmt.Sprintf("response.Data 无响应数据，response = %v", response))
		}
	}
	return resultMap
}

// ProcessAndStoreTrafficData 处理结果集并持久化
func ProcessAndStoreTrafficData(
	resultMap map[int]TrafficData,
	actionToPolicy map[int]int,
	upList, dnList []int,
) error {
	// 初始化结果集map
	policyIdLogMap := make(map[int]*policyModel.DimControlPolicyLog)
	currentTime := time.Now()

	// 需要根据policyId的维度重组日志
	for actionId, trafficData := range resultMap {
		policyID := actionToPolicy[actionId]
		allTraffic := trafficData.AllTrafficData
		throughTraffic := trafficData.ThroughTrafficData
		discarded := allTraffic - throughTraffic

		if contains(upList, actionId) {
			if entity, exists := policyIdLogMap[policyID]; exists {
				entity.UpTraffic += allTraffic
				entity.UpPass += throughTraffic
				entity.UpDiscard += discarded
			} else {
				policyIdLogMap[policyID] = &policyModel.DimControlPolicyLog{
					PolicyId:   policyID,
					UpTraffic:  allTraffic,
					UpPass:     throughTraffic,
					UpDiscard:  discarded,
					RecordTime: currentTime,
				}
			}
			continue
		}

		if contains(dnList, actionId) {
			if entity, exists := policyIdLogMap[policyID]; exists {
				entity.DnTraffic += allTraffic
				entity.DnPass += throughTraffic
				entity.DnDiscard += discarded
			} else {
				policyIdLogMap[policyID] = &policyModel.DimControlPolicyLog{
					PolicyId:   policyID,
					DnTraffic:  allTraffic,
					DnPass:     throughTraffic,
					DnDiscard:  discarded,
					RecordTime: currentTime, // Set to current time
				}
			}
		}
	}

	var logEntities []*policyModel.DimControlPolicyLog
	for _, logs := range policyIdLogMap {
		logEntities = append(logEntities, logs)
	}

	// 批量存储日志
	if len(logEntities) > 0 {
		if result := global.ServiceDB.CreateInBatches(logEntities, 1000); result.Error != nil {
			global.Log.Error("policyLog日志入库失败：", zap.Error(result.Error))
			return result.Error
		} else {
			global.Log.Info(fmt.Sprintf("分流日志数据成功入库%d条：", result.RowsAffected))
		}
	}
	return nil
}

// contains checks if a slice contains a specific string value
func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
