package configuration

import (
	"errors"
	"fcas_server/global"
	"fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fcas_server/model/system"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

type DimBypassSvc struct {
	Log   *zap.Logger
	Mysql *gorm.DB
}

func NewDimBypassSvc(log *zap.Logger, mysql *gorm.DB) *DimBypassSvc {
	return &DimBypassSvc{
		Log:   log.Named("[DimBypass-Bypass管理]"),
		Mysql: mysql,
	}
}

func (svc DimBypassSvc) GetById(id string) (result configuration.DimBypass, err error) {
	tx := svc.Mysql.Model(configuration.DimBypass{}).Debug()

	// 分页
	err = tx.First(&result, id).Error
	if err != nil {
		svc.Log.Error("根据ID获取信息失败", zap.Error(err))
		return result, err
	}

	return result, nil
}

func (svc DimBypassSvc) List(req req.ListRequest) (result []*configuration.DimBypass, total int64, err error) {
	tx := svc.Mysql.Model(configuration.DimBypass{}).Debug()
	if req.Key != "" {
		key := fmt.Sprint("%", req.Key, "%")
		tx = tx.Where("bypass_name LIKE ?", key).Or("bypass_ip LIKE ?", key)
	}

	// 分页
	err = tx.Count(&total).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&result).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Error(err))
		return result, total, err
	}

	// 获取bypass-status并填充
	svc.fillBypassStatus(result)

	return result, total, nil
}

func (svc DimBypassSvc) Save(req req.DimBypassSaveRequest) error {
	tx := svc.Mysql.Debug()

	err := tx.Create(&configuration.DimBypass{
		OlpId:      req.OlpId,
		BypassName: req.BypassName,
		BypassIp:   req.BypassIp,
		BypassPort: req.BypassPort,
		Remark:     req.Remark,
	}).Error
	if err != nil {
		svc.Log.Error("创建失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimBypassSvc) Update(req req.DimBypassSaveRequest) error {
	tx := svc.Mysql.Debug()

	err := tx.Save(&configuration.DimBypass{
		Id:         req.ID,
		OlpId:      req.OlpId,
		BypassName: req.BypassName,
		BypassIp:   req.BypassIp,
		BypassPort: req.BypassPort,
		Remark:     req.Remark,
	}).Error
	if err != nil {
		svc.Log.Error("更新失败", zap.Error(err))
		return err
	}

	return nil
}

func (svc DimBypassSvc) Delete(ids []int) error {
	tx := svc.Mysql.Debug()

	err := tx.Delete(&configuration.DimBypass{}, ids).Error
	if err != nil {
		svc.Log.Error("删除失败", zap.Error(err))
		return err
	}

	return nil
}

// ValidateUnique 校验名称是否存在, id为0表示新增操作, id不为0表示更新操作（更新时需要排除当前记录）
func (svc DimBypassSvc) ValidateUnique(name string, id int) error {
	tx := svc.Mysql.Model(configuration.DimBypass{})
	if id != 0 {
		tx = tx.Where("id != ?", id)
	}
	var count int64
	err := tx.Where("bypass_name = ?", name).Count(&count).Error
	if err != nil {
		svc.Log.Error("校验bypass名称时出现异常", zap.Error(err))
		return fmt.Errorf("校验bypass名称时出现异常: %w", err)
	}

	if count > 0 {
		return errors.New("bypass名称已存在")
	}

	return nil
}

func (svc DimBypassSvc) ValidateBypassPassword(currUserId uint, password string) (bool, error) {

	var currUser system.SysUser
	if err := global.SystemDB.Model(system.SysUser{}).Where("id = ?", currUserId).First(&currUser).Error; err != nil {
		return false, fmt.Errorf("获取当前用户设置的Bypass切换密码失败: %w", err)
	}

	if password != currUser.BypassPassword {
		//return map[string]any{
		//	"succ": false,
		//	"msg":  "Bypass切换密码错误，请重新输入",
		//}, nil
		return false, nil
	}

	return true, nil
}

func (svc DimBypassSvc) SetBypassStatus(dimBypass configuration.DimBypass) error {

	if err := dimBypass.SetStatus(dimBypass.Status); err != nil {
		svc.Log.Error("设置bypass-status失败", zap.Error(err))
		return err
	}

	return nil
}

// fillBypassStatus 填充bypass-status，并发处理
func (svc DimBypassSvc) fillBypassStatus(records []*configuration.DimBypass) {
	numWorker := 10
	jobs := make(chan *configuration.DimBypass, len(records))
	var wg sync.WaitGroup

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobs {
				// 获取bypass-status并填充
				if err := item.GetStatus(); err != nil {
					svc.Log.Error("填充bypass-status失败", zap.Error(err))
					continue
				}
			}
		}()
	}

	for _, item := range records {
		jobs <- item
	}
	close(jobs)
	wg.Wait()
}
