package configuration

import (
	"errors"
	"fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DimDeviceInfoSvc struct {
	Mysql *gorm.DB
	Log   *zap.Logger
}

func NewDimDeviceInfoSvc(log *zap.Logger, mysql *gorm.DB) *DimDeviceInfoSvc {
	return &DimDeviceInfoSvc{
		Log:   log.Named("[DimDeviceInfo-设备管理]"),
		Mysql: mysql,
	}
}

func (svc DimDeviceInfoSvc) GetById(id string) (result configuration.DimDeviceInfo, err error) {
	tx := svc.Mysql.Model(configuration.DimDeviceInfo{}).Debug()

	err = tx.First(&result, id).Error
	if err != nil {
		svc.Log.Error("根据ID获取信息失败", zap.Error(err))
		return result, err
	}

	return result, nil
}

func (svc DimDeviceInfoSvc) List(req req.ListRequest) (result []*configuration.DimDeviceInfo, total int64, err error) {
	tx := svc.Mysql.Model(configuration.DimDeviceInfo{}).Debug()
	if req.Key != "" {
		key := fmt.Sprint("%", req.Key, "%")
		tx = tx.Where("device_name LIKE ?", key).Or("device_ip LIKE ?", key)
	}

	// 分页
	err = tx.Count(&total).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&result).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Error(err))
		return result, total, err
	}

	// 填充snmp相关信息
	for _, r := range result {
		if err := r.SetSnmpFields(); err != nil {
			svc.Log.Error("填充snmp相关信息失败", zap.Any("device_ip", r.DeviceIp), zap.Error(err))
		}
	}

	return result, total, nil
}

func (svc DimDeviceInfoSvc) Save(req req.DimDeviceInfoSaveRequest) error {
	tx := svc.Mysql.Debug()

	err := tx.Create(&configuration.DimDeviceInfo{
		DeviceName: req.DeviceName,
		DeviceIp:   req.DeviceIp,
		SnmpName:   req.SnmpName,
		UdpPort:    req.UdpPort,
		Remark:     req.Remark,
	}).Error
	if err != nil {
		svc.Log.Error("创建失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimDeviceInfoSvc) Update(req req.DimDeviceInfoSaveRequest) error {
	tx := svc.Mysql.Debug()

	err := tx.Save(&configuration.DimDeviceInfo{
		ID:         req.ID,
		DeviceName: req.DeviceName,
		DeviceIp:   req.DeviceIp,
		SnmpName:   req.SnmpName,
		UdpPort:    req.UdpPort,
		Remark:     req.Remark,
	}).Error
	if err != nil {
		svc.Log.Error("更新失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimDeviceInfoSvc) Delete(ids []int) error {
	tx := svc.Mysql.Debug()

	err := tx.Delete(&configuration.DimDeviceInfo{}, ids).Error
	if err != nil {
		svc.Log.Error("删除失败", zap.Error(err))
		return err
	}

	return nil
}

// ValidateUnique 校验名称是否存在, id为0表示新增操作, id不为0表示更新操作（更新时需要排除当前记录）
func (svc DimDeviceInfoSvc) ValidateUnique(name string, id int) error {
	tx := svc.Mysql.Model(configuration.DimDeviceInfo{})
	if id != 0 {
		tx = tx.Where("id != ?", id)
	}
	var count int64
	err := tx.Where("device_name = ?", name).Count(&count).Error
	if err != nil {
		svc.Log.Error("校验设备名称时出现异常", zap.Error(err))
		return fmt.Errorf("校验设备名称时出现异常: %w", err)
	}

	if count > 0 {
		return errors.New("设备名称已存在")
	}

	return nil
}
