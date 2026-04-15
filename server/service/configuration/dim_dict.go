package configuration

import (
	"errors"
	"fcas_server/global"
	"fcas_server/model/common/dict"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

type DimDictService struct{}

func (svc DimDictService) GetDictByType(dictType string) (result interface{}, err error) {
	db := global.ServiceDB
	var list []dict.DimDict

	switch dictType {
	case global.DictTypeAllIsp:
		err = db.Table("isp_view").Find(&list).Error
	case global.DictTypeSelectIsp:
		ispList := strings.Split(global.CONFIG.SelectIsp, ",")
		err = db.Table("isp_view").Select("name").Where("name in (?)", ispList).Group("name").Find(&list).Error
	case global.DictTypeAppType:
		err = db.Table("app_type_view").Find(&list).Error
	case global.DictTypeAppId:
		err = db.Table("app_id_view").Find(&list).Error
	case global.DictTypeAppTypeIdTree:
		list, err = getAppTypeIdTree()
	default:
		return list, errors.New("该数据字典的类型不存在")
	}
	if err != nil {
		global.Log.Error(fmt.Sprintf("[dictType = %s] 获取字典信息失败!", dictType), zap.Error(err))
		return list, err
	}
	return list, nil
}

func getAppTypeIdTree() ([]dict.DimDict, error) {
	var list []dict.DimDict
	err := global.ServiceDB.Table("app_type_view").Find(&list).Error
	if err != nil {
		global.Log.Error("获取字典信息失败", zap.Error(err))
		return list, err
	}
	for i := range list {
		var childList []dict.DimDict
		err = global.ServiceDB.Table("dim_app_classify").
			Select("app_id as id,app_name as name").
			Where("app_type_id = ?", list[i].Id).
			Find(&childList).Error
		list[i].Children = childList
	}
	return list, err
}
