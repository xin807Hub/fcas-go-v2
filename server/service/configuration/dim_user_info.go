package configuration

import (
	"bytes"
	"errors"
	"fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fcas_server/utils/addr_set"
	"fmt"
	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type DimUserInfoSvc struct {
	Log     *zap.Logger
	Mysql   *gorm.DB
	AddrSet *addr_set.AddrSet
}

func NewDimUserInfoSvc(log *zap.Logger, mysql *gorm.DB, addrSet *addr_set.AddrSet) *DimUserInfoSvc {
	return &DimUserInfoSvc{
		Log:     log.Named("[DimUserInfo-用户管理]"),
		Mysql:   mysql,
		AddrSet: addrSet,
	}
}

func (svc DimUserInfoSvc) GetById(id string) (result configuration.DimUserInfo, err error) {

	err = svc.Mysql.Model(configuration.DimUserInfo{}).First(&result, id).Error
	if err != nil {
		svc.Log.Error("根据ID获取信息失败", zap.Error(err))
		return result, err
	}

	return result, nil
}

func (svc DimUserInfoSvc) List(req req.ListRequest) (result []*configuration.DimUserInfo, total int64, err error) {
	tx := svc.Mysql.Model(configuration.DimUserInfo{})
	if req.Key != "" {
		tx = tx.Where("user_name LIKE ?", "%"+req.Key+"%").Or(`ip_address->"$[*]" LIKE ?`, "%"+req.Key+"%")
	}

	// 分页
	err = tx.Count(&total).Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&result).Error
	if err != nil {
		svc.Log.Error("获取列表信息失败", zap.Error(err))
		return result, total, err
	}

	return result, total, nil
}

func (svc DimUserInfoSvc) Save(req req.DimUserInfoSaveRequest) error {

	// 处理IP地址
	addrs, addrNum, err := svc.AddrSet.Add(req.IpAddress)
	if err != nil {
		err = fmt.Errorf("IP地址保存失败,该组IP全部或部分已经保存: %w", err)
		svc.Log.Error("创建失败", zap.Any("addrs", req.IpAddress), zap.Error(err))
		return err
	}

	// 保存到数据库
	err = svc.Mysql.Create(&configuration.DimUserInfo{
		UserName:     req.UserName,
		UserType:     req.UserType,
		Remark:       req.Remark,
		IpAddress:    addrs,
		IpAddressNum: addrNum,
	}).Error
	if err != nil {
		svc.Log.Error("创建失败", zap.Any("req", req), zap.Error(err))
		svc.AddrSet.Remove(addrs...) // 回滚IP地址
		return err
	}

	return nil
}

func (svc DimUserInfoSvc) Update(req req.DimUserInfoSaveRequest) error {

	// 获取原有IP地址
	oldAddr, err := svc.getAddrById(req.ID)
	if err != nil {
		return fmt.Errorf("获取原有IP地址失败: %w", err)
	}

	addrs, addrNum, err := svc.AddrSet.Update(oldAddr, req.IpAddress)
	if err != nil {
		err = fmt.Errorf("IP地址保存失败,该组IP全部或部分已经保存: %w", err)
		svc.Log.Error("更新失败", zap.Any("addrs", req.IpAddress), zap.Error(err))
		return err
	}

	// 更新数据库
	err = svc.Mysql.Save(&configuration.DimUserInfo{
		ID:           req.ID,
		UserName:     req.UserName,
		UserType:     req.UserType,
		Remark:       req.Remark,
		IpAddress:    addrs,
		IpAddressNum: addrNum,
	}).Error
	if err != nil {
		if _, _, err := svc.AddrSet.Add(oldAddr); err != nil {
			svc.Log.Error("更新失败且回滚IP地址失败", zap.Any("req", req), zap.Any("old_addrs", oldAddr), zap.Error(err))
			return err
		}
		svc.Log.Error("更新失败", zap.Any("req", req), zap.Error(err))
		return err
	}

	return nil
}

func (svc DimUserInfoSvc) Delete(ids []int) error {

	// 待删除用户保存的IP地址
	oldAddr, err := svc.getAddrById(ids...)
	if err != nil {
		return fmt.Errorf("获取待删除用户保存的IP地址失败: %w", err)
	}

	// 删除数据库，需检查是否被用户组引用
	if err := svc.Mysql.Transaction(func(tx *gorm.DB) error {
		var count int64
		if err := tx.Raw(`SELECT COUNT(*) FROM dim_user_crowd_relation WHERE user_id IN ?`, ids).Scan(&count).Error; err != nil {
			return fmt.Errorf("查询待删除用户是否被用户组引用失败: %w", err)
		}
		if count != 0 {
			return errors.New("该用户被用户组引用，请先解除引用")
		}
		if err := tx.Delete(&configuration.DimUserInfo{}, ids).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		svc.Log.Error("删除失败", zap.Any("ids", ids), zap.Error(err))
		return err
	}

	// 删除IP地址
	svc.AddrSet.Remove(oldAddr...)

	return nil
}

func (svc DimUserInfoSvc) Export(req req.ListRequest) (output []byte, err error) {
	req.Limit = -1
	req.Page = 2
	records, _, err := svc.List(req)
	if err != nil {
		return nil, err
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("用户信息")
	if err != nil {
		return nil, fmt.Errorf("创建sheet失败: %w", err)
	}

	row := sheet.AddRow()
	header := []string{"用户名称", "IP地址/IP地址段", "IP地址个数", "备注"}
	for _, item := range header {
		row.AddCell().SetString(item)
	}

	for _, r := range records {
		row := sheet.AddRow()
		row.AddCell().SetString(r.UserName)
		row.AddCell().SetString(strings.Join(r.IpAddress, ", "))
		row.AddCell().SetString(r.IpAddressNum)
		row.AddCell().SetString(r.Remark)
	}

	var buf bytes.Buffer
	if err := file.Write(&buf); err != nil {
		return nil, fmt.Errorf("failed to write xlsx file: %w", err)
	}

	return buf.Bytes(), nil
}

// getAddrById 根据ID获取IP地址
func (svc DimUserInfoSvc) getAddrById(ids ...int) ([]string, error) {
	var addrs []struct {
		IpAddr []string `gorm:"column:ip_address;type:json;json"`
	}
	err := svc.Mysql.Model(configuration.DimUserInfo{}).Where("id in (?)", ids).Pluck("ip_address", &addrs).Error
	if err != nil {
		return nil, err
	}

	var result []string
	for _, addr := range addrs {
		result = append(result, addr.IpAddr...)
	}
	return result, nil
}

// ValidateUnique 校验用户名称是否存在, id为0表示新增操作, id不为0表示更新操作（更新时需要排除当前记录）
func (svc DimUserInfoSvc) ValidateUnique(name string, id int) error {
	tx := svc.Mysql.Model(configuration.DimUserInfo{})
	if id != 0 {
		tx = tx.Where("id != ?", id)
	}
	var count int64
	err := tx.Where("user_name = ?", name).Count(&count).Error
	if err != nil {
		svc.Log.Error("校验用户名称时出现异常", zap.Error(err))
		return fmt.Errorf("校验用户名称时出现异常: %w", err)
	}

	if count > 0 {
		return errors.New("用户名称已存在")
	}

	return nil
}
