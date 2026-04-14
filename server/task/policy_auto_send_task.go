package task

import (
	"fcas_server/global"
	"fcas_server/model/policy"
	policy2 "fcas_server/service/policy"
	"fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func PolicyAutoSendTask() {
	global.Log.Info("过期策略定时任务 start...")
	// 获取当前时间
	now := time.Now()
	// 查询过期策略
	var expiredPolicies []policy.DimControlPolicy
	err := global.ServiceDB.Model(&policy.DimControlPolicy{}).Where(" end_time < ? AND status = ?", now, global.Bind).Find(&expiredPolicies).Error
	if err != nil {
		global.Log.Error("查询过期策略出错:", zap.Error(err))
		return
	}
	for _, expiredPolicy := range expiredPolicies {
		global.Log.Info(fmt.Sprintf("Expired Policy ID: %d\n", expiredPolicy.Id))

		policyJson := policy2.BuildPolicyJson(expiredPolicy, global.UnBind)

		global.Log.Info(fmt.Sprintf("policyJson = %s\n", policyJson))

		// 发送消息
		utils.SendMessage(policyJson, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)

		// 解绑后，更新策略状态
		err = global.ServiceDB.Exec(`update dim_control_policy SET status = ? where id = ?`, global.UnBind, expiredPolicy.Id).Error
		if err != nil {
			global.Log.Error("更新策略状态出错:", zap.Error(err))
			continue
		}
	}
	global.Log.Info("过期策略定时任务 end...")

	global.Log.Info("延期生效策略定时任务 start...")
	// 查询延期策略
	var unBindPolicies []policy.DimControlPolicy
	err = global.ServiceDB.Model(&policy.DimControlPolicy{}).Where("start_time >= ? AND status = ?", now, global.UnBind).Find(&unBindPolicies).Error
	if err != nil {
		global.Log.Info("查询延期策略出错:", zap.Error(err))
		return
	}

	for _, unBindPolicy := range unBindPolicies {
		global.Log.Info(fmt.Sprintf("UnBind Policy ID: %d\n", unBindPolicy.Id))

		bindPolicyJson := policy2.BuildPolicyJson(unBindPolicy, global.Bind)
		global.Log.Info(fmt.Sprintf("unBindPolicyJson = %s\n", bindPolicyJson))

		// 发送消息
		utils.SendMessage(bindPolicyJson, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)

		// 绑定后，更新策略状态
		err = global.ServiceDB.Exec(`update dim_control_policy SET status = ? WHERE id = ?`, global.Bind, unBindPolicy.Id).Error
		if err != nil {
			global.Log.Error("更新策略状态出错:", zap.Error(err))
			continue
		}
	}
	global.Log.Info("延期生效策略定时任务 end...")
}
