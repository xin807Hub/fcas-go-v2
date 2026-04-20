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
	global.Log.Info("过期流控策略定时任务 start...")
	now := time.Now()

	var expiredPolicies []policy.DimControlPolicy
	err := global.ServiceDB.Model(&policy.DimControlPolicy{}).
		Where("end_time < ? AND status = ?", now, global.Bind).
		Find(&expiredPolicies).Error
	if err != nil {
		global.Log.Error("查询过期流控策略失败", zap.Error(err))
		return
	}
	for _, expiredPolicy := range expiredPolicies {
		policyJSON := policy2.BuildPolicyJson(expiredPolicy, global.UnBind)
		global.Log.Info(fmt.Sprintf("expired policy id=%d, payload=%s", expiredPolicy.Id, policyJSON))
		utils.SendMessage(policyJSON, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)

		err = global.ServiceDB.Exec(`update dim_control_policy set status = ? where id = ?`, global.UnBind, expiredPolicy.Id).Error
		if err != nil {
			global.Log.Error("更新流控策略状态失败", zap.Error(err))
		}
	}
	global.Log.Info("过期流控策略定时任务 end...")

	global.Log.Info("到期生效流控策略定时任务 start...")
	var unBindPolicies []policy.DimControlPolicy
	err = global.ServiceDB.Model(&policy.DimControlPolicy{}).
		Where("start_time <= ? AND end_time >= ? AND status = ?", now, now, global.UnBind).
		Find(&unBindPolicies).Error
	if err != nil {
		global.Log.Error("查询到期生效流控策略失败", zap.Error(err))
		return
	}

	for _, unBindPolicy := range unBindPolicies {
		bindPolicyJSON := policy2.BuildPolicyJson(unBindPolicy, global.Bind)
		global.Log.Info(fmt.Sprintf("bind policy id=%d, payload=%s", unBindPolicy.Id, bindPolicyJSON))
		utils.SendMessage(bindPolicyJSON, global.CONFIG.Policy.SendUrl, strconv.Itoa(policy.ControlPolicyMsgType), global.CONFIG.Policy.Dir)

		err = global.ServiceDB.Exec(`update dim_control_policy set status = ? where id = ?`, global.Bind, unBindPolicy.Id).Error
		if err != nil {
			global.Log.Error("更新流控策略状态失败", zap.Error(err))
		}
	}
	global.Log.Info("到期生效流控策略定时任务 end...")
}
