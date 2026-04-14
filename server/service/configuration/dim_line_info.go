package configuration

import (
	"errors"
	"fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DimLineInfoSvc struct {
	Log   *zap.Logger
	Mysql *gorm.DB
}

func NewDimLineInfoSvc(log *zap.Logger, mysql *gorm.DB) *DimLineInfoSvc {
	return &DimLineInfoSvc{
		Log:   log.Named("[DimLineInfo-链路管理]"),
		Mysql: mysql,
	}
}

func (svc DimLineInfoSvc) GetById(id string) (result configuration.DimLineInfo, err error) {
	tx := svc.Mysql.Model(configuration.DimLineInfo{})

	err = tx.First(&result, id).Error
	if err != nil {
		svc.Log.Error("根据ID获取信息失败", zap.Error(err))
		return result, err
	}

	return result, nil
}

func (svc DimLineInfoSvc) List(req req.ListRequest) (result []*configuration.DimLineInfo, total int64, err error) {
	tx := svc.Mysql.Model(configuration.DimLineInfo{})
	if req.Key != "" {
		key := fmt.Sprint("%", req.Key, "%")
		tx = tx.Where("line_name LIKE ?", key).Or("line_vlan LIKE ?", key)
	}

	// 分页
	err = tx.Count(&total).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&result).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Error(err))
		return result, total, err
	}

	return result, total, nil
}

func (svc DimLineInfoSvc) Save(req req.DimLineInfoSaveRequest) error {

	err := svc.Mysql.Create(&configuration.DimLineInfo{
		LineName: req.LineName,
		LineNum:  req.LineNum,
		LineVlan: req.LineVlan,
		Remark:   req.Remark,
	}).Error
	if err != nil {
		svc.Log.Error("创建失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimLineInfoSvc) Update(req req.DimLineInfoSaveRequest) error {

	err := svc.Mysql.Save(&configuration.DimLineInfo{
		ID:       req.ID,
		LineName: req.LineName,
		LineNum:  req.LineNum,
		LineVlan: req.LineVlan,
		Remark:   req.Remark,
	}).Error
	if err != nil {
		svc.Log.Error("更新失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimLineInfoSvc) Delete(ids []int) error {

	err := svc.Mysql.Delete(&configuration.DimLineInfo{}, ids).Error
	if err != nil {
		svc.Log.Error("删除失败", zap.Error(err))
		return err
	}

	return nil
}

// ValidateUnique 校验名称是否存在, id为0表示新增操作, id不为0表示更新操作（更新时需要排除当前记录）
func (svc DimLineInfoSvc) ValidateUnique(name string, vlan, id int) error {
	tx := svc.Mysql.Model(configuration.DimLineInfo{})
	if id != 0 {
		tx = tx.Where("id != ?", id)
	}
	var count int64
	err := tx.Where("line_name = ?", name).Or("line_vlan = ?", vlan).Count(&count).Error
	if err != nil {
		svc.Log.Error("校验名称和vlan时出现异常", zap.Error(err))
		return fmt.Errorf("校验名称和vlan时出现异常: %w", err)
	}

	if count > 0 {
		return errors.New("链路名称或vlan已存在")
	}

	return nil
}
