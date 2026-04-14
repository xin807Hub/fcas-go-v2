package traffic

import (
	"fcas_server/model/traffic"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AlarmLogSvc struct {
	Log   *zap.Logger
	Mysql *gorm.DB
}

func NewAlarmLogSvc(log *zap.Logger, mysql *gorm.DB) *AlarmLogSvc {
	return &AlarmLogSvc{
		Log:   log.Named("[TrafficAlarmLog-流量告警日志]"),
		Mysql: mysql,
	}
}

func (svc AlarmLogSvc) List(params traffic.AlarmLogListRequest) (output []*traffic.AlarmLog, total int64, err error) {
	tx := svc.Mysql.Model(traffic.AlarmLog{})
	if params.StartTime != nil && params.EndTime != nil {
		tx = tx.Where("? <= start_time AND start_time < ?", params.StartTime, params.EndTime)
	}
	if params.AppTypeId != nil {
		tx = tx.Where("app_type_id = ?", params.AppTypeId)
	}
	if params.AppId != nil {
		tx = tx.Where("app_id = ?", params.AppId)
	}
	if len(params.LinkIds) != 0 {
		tx = tx.Where("link_id IN ?", params.LinkIds)
	}

	err = tx.Count(&total).Limit(params.Limit).Offset((params.Page - 1) * params.Limit).Find(&output).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Any("params", params), zap.Error(err))
		return nil, 0, err
	}

	return output, total, nil
}
