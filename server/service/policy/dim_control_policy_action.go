package policy

import (
	"fcas_server/global"
	"fcas_server/model/policy"
	"go.uber.org/zap"
)

var dimControlPolicyService = DimControlPolicyService{}

type DimControlPolicyActionService struct {
}

func (DimControlPolicyActionService) List() ([]policy.DimControlPolicyAction, error) {
	var actions []policy.DimControlPolicyAction
	if err := global.ServiceDB.Model(&policy.DimControlPolicyAction{}).Find(&actions).Error; err != nil {
		global.Log.Error("查询策略Action 失败", zap.Error(err))
		return nil, err
	}
	return actions, nil
}

func (DimControlPolicyActionService) SaveOrUpdate(action policy.DimControlPolicyAction) error {
	var result []policy.DimControlPolicyAction
	if err := global.ServiceDB.Model(&policy.DimControlPolicyAction{}).
		Where("policy_id", action.PolicyId).
		Where("vlan_id", action.VlanId).
		Where("upload_device_id", action.UploadDeviceId).Find(&result).Error; err != nil {
		global.Log.Error("策略配置PolicyAction查询失败", zap.Error(err))
		return err
	}
	if len(result) > 0 {
		action.Id = result[0].Id
	}
	if err := global.ServiceDB.Save(&action).Error; err != nil {
		global.Log.Error("策略配置PolicyAction保存失败", zap.Error(err))
		return err
	}
	// 更新controlPolicy表中字段
	controlPolicy := dimControlPolicyService.getByPolicyId(action.PolicyId)
	if controlPolicy != nil {
		controlPolicy.Remark = action.ToString()
		_ = dimControlPolicyService.SaveOrUpdateControlPolicy(*controlPolicy)
	}
	return nil
}
