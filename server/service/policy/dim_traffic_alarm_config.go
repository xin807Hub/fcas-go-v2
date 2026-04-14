package policy

import (
	"fcas_server/global"
	"fcas_server/model/policy"
	"go.uber.org/zap"
	"time"
)

type DimTrafficAlarmConfigService struct {
}

func (DimTrafficAlarmConfigService) PageDimTrafficAlarmConfig(req policy.DimTrafficAlarmConfigReq) (result []policy.DimTrafficAlarmConfig, total int64, err error) {
	limit := req.Limit
	offset := req.Limit * (req.Page - 1)

	dbAlarm := global.ServiceDB.Model(&policy.DimTrafficAlarmConfig{})
	if len(req.Name) > 0 {
		dbAlarm.Where("name like ?", "%"+req.Name+"%")
	}
	if req.AppTypeId != 0 {
		dbAlarm.Where("app_type_id = ?", req.AppTypeId)
	}
	if req.AppId != 0 {
		dbAlarm.Where("app_id = ?", req.AppId)
	}

	if err := dbAlarm.Count(&total).Limit(limit).Offset(offset).Order("id desc").Find(&result).Error; err != nil {
		global.Log.Error("策略告警配置信息查询失败", zap.Error(err))
		return nil, 0, err
	}

	return result, total, nil
}

func (DimTrafficAlarmConfigService) GetById(req policy.DimTrafficAlarmConfigReq) (rest *policy.DimTrafficAlarmConfig) {
	if err := global.ServiceDB.Model(&policy.DimTrafficAlarmConfig{}).Where("id = ?", req.Id).Find(rest).Error; err != nil {
		global.Log.Error("策略告警单个配置信息查询失败", zap.Error(err))
		return nil
	}
	return rest
}

func (DimTrafficAlarmConfigService) SaveOrUpdateAlarmConfig(alarmConfig policy.DimTrafficAlarmConfig) error {
	if alarmConfig.Id == 0 {
		if alarmConfig.BaseValue != 0 {
			increaseVal := float64(alarmConfig.BaseValue) * (alarmConfig.IncreaseRatio / 100)
			decreaseVal := float64(alarmConfig.BaseValue) * (alarmConfig.DecreaseRatio / 100)
			increaseBaseVal := increaseVal + float64(alarmConfig.BaseValue)
			decreaseBaseVal := float64(alarmConfig.BaseValue) - decreaseVal
			alarmConfig.IncreaseBaseValue = int(increaseBaseVal)
			alarmConfig.DecreaseBaseValue = int(decreaseBaseVal)
		}
		if err := global.ServiceDB.Model(policy.DimTrafficAlarmConfig{}).Save(&alarmConfig).Error; err != nil {
			global.Log.Error("策略告警配置保存失败", zap.Error(err))
			return err
		}
	} else {
		if err := global.ServiceDB.Model(policy.DimTrafficAlarmConfig{}).Where("id = ?", alarmConfig.Id).Updates(&alarmConfig).Error; err != nil {
			global.Log.Error("策略告警配置修改失败", zap.Error(err))
			return err
		}
	}
	return nil
}

func (DimTrafficAlarmConfigService) DeleteAlarmConfig(ids []int) error {
	if err := global.ServiceDB.Model(&policy.DimTrafficAlarmConfig{}).Where("id in ?", ids).Delete(policy.DimTrafficAlarmConfig{}).Error; err != nil {
		global.Log.Error("策略告警配置删除失败", zap.Error(err))
		return err
	}
	return nil
}

// 获取有效告警配置
func (DimTrafficAlarmConfigService) GetAlarmConfigValid() (result []policy.DimTrafficAlarmConfig) {
	nowTime := time.Now().Format(global.TimeLayout)
	if err := global.ServiceDB.Model(&policy.DimTrafficAlarmConfig{}).
		Where("start_time <= ?", nowTime).
		Where("end_time >= ?", nowTime).
		Find(&result).Error; err != nil {
		global.Log.Error("策略告警配置查询失败", zap.Error(err))
		return nil
	}
	return result
}
