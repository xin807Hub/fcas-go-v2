package traffic

import (
	"errors"
	"fcas_server/global"
	"fcas_server/model/common/response"
	"fcas_server/model/traffic"
	utils2 "fcas_server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserActionService struct{}

const (
	DataTypeDstIp = "dst_ip"
	DataTypeAppId = "app_id"
)

func (service UserActionService) GetUserActionTable(param traffic.UserActionReqParam) (response.PageResult, error) {
	var result = response.PageResult{
		CurrPage: param.Page,
		PageSize: param.Limit,
	}
	var tableData []traffic.UserActionTable
	var selectStr string
	if param.DataType == DataTypeDstIp {
		selectStr = "dst_ip AS name,sumMerge(bytes_up_view) AS up_byte,sumMerge(bytes_dn_view) AS dn_byte,up_byte + dn_byte AS total_byte"
	} else if param.DataType == DataTypeAppId {
		selectStr = "app_id,dictGet('app_id_dict','name',app_id) AS name,sumMerge(bytes_up_view) AS up_byte,sumMerge(bytes_dn_view) AS dn_byte,up_byte + dn_byte AS total_byte"
	}
	tableNameNew := "bigdata_fcas_v2.dws_gen_traffic_hour"
	tableNameOld := "bigdata_fcas.dws_gen_traffic_hour"

	ckV2Db := global.V2ClickhouseDB.Table(tableNameNew)
	ckV1Db := global.V1ClickhouseDB.Table(tableNameOld)

	if param.SrcIp != "" {
		ckV2Db = ckV2Db.Where("src_ip Like ?", "%"+param.SrcIp+"%")
		ckV1Db = ckV1Db.Where("src_ip Like ?", "%"+param.SrcIp+"%")
	}
	if param.Limit == 0 {
		param.Limit = 50
	}

	var ckDb *gorm.DB
	timeRangeType := utils2.GetDbTypeByTimeRange(param.StartTime, param.EndTime)
	switch timeRangeType {
	case global.QueryNew:
		ckDb = ckV2Db.Select(selectStr).Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	case global.QueryCrossOld2New:
		ckDb = global.V2ClickhouseDB.Table("(? UNION ALL ?)",
			global.V2ClickhouseDB.Table(tableNameNew).Select("start_time,dst_ip,host,app_id,bytes_up_view,bytes_dn_view").Where("start_time >= ? AND start_time < ?", global.CONFIG.DeploymentDate, param.EndTime),
			global.V2ClickhouseDB.Table(tableNameOld).Select("start_time,dst_ip,host,app_id,bytes_up_view,bytes_dn_view").Where("start_time >= ? AND start_time < ?", param.StartTime, global.CONFIG.DeploymentDate)).
			Select(selectStr)
	case global.QueryOld:
		ckDb = ckV1Db.Select(selectStr).Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	default:
		return result, errors.New("开始、结束时间解析错误")
	}

	var total int64
	err := ckDb.Group(param.DataType).Count(&total).Limit(param.Limit).Offset(param.Limit * (param.Page - 1)).Order("total_byte DESC").Find(&tableData).Error
	if err != nil {
		global.Log.Error("获取用户行为分析表格数据错误", zap.Error(err))
		return result, err
	}

	result.List = tableData
	result.TotalCount = total
	return result, nil
}

func (service UserActionService) GetUserActionDetail(param traffic.UserActionReqParam) ([]traffic.UserActionDetail, error) {
	var result []traffic.UserActionDetail
	if param.DstIp == "" && param.AppId == 0 {
		return result, errors.New("查询详情时appId或dstIp不得为空")
	}
	var selectStr string
	if param.DataType == DataTypeDstIp {
		selectStr = "dictGet('app_id_dict','name',app_id) AS name,host,sumMerge(bytes_up_view) AS up_byte,sumMerge(bytes_dn_view) AS dn_byte,up_byte + dn_byte AS total_byte"
	} else if param.DataType == DataTypeAppId {
		selectStr = "dst_ip AS name,host,sumMerge(bytes_up_view) AS up_byte,sumMerge(bytes_dn_view) AS dn_byte,up_byte + dn_byte AS total_byte"
	}

	tableNameNew := "bigdata_fcas_v2.dws_gen_traffic_hour"
	tableNameOld := "bigdata_fcas.dws_gen_traffic_hour"

	ckV2Db := global.V2ClickhouseDB.Table(tableNameNew)
	ckV1Db := global.V1ClickhouseDB.Table(tableNameOld)

	if param.SrcIp != "" {
		ckV2Db = ckV2Db.Where("src_ip LIKE ?", "%"+param.SrcIp+"%")
		ckV1Db = ckV1Db.Where("src_ip LIKE ?", "%"+param.SrcIp+"%")
	}
	if param.DstIp != "" {
		ckV2Db = ckV2Db.Where("dst_ip = ?", param.DstIp)
		ckV1Db = ckV1Db.Where("dst_ip = ?", param.DstIp)
	}
	if param.AppId != 0 {
		ckV2Db = ckV2Db.Where("app_id = ?", param.AppId)
		ckV1Db = ckV1Db.Where("app_id = ?", param.AppId)
	}
	if param.Limit == 0 {
		param.Limit = 10
	}

	var ckDb *gorm.DB
	timeRangeType := utils2.GetDbTypeByTimeRange(param.StartTime, param.EndTime)
	switch timeRangeType {
	case global.QueryNew:
		ckDb = ckV2Db.Select(selectStr).Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	case global.QueryCrossOld2New:
		ckDb = global.V2ClickhouseDB.Table("(? UNION ALL ?)",
			global.V2ClickhouseDB.Table(tableNameNew).Select("start_time,dst_ip,host,app_id,bytes_up_view,bytes_dn_view").Where("start_time >= ? AND start_time < ?", global.CONFIG.DeploymentDate, param.EndTime),
			global.V2ClickhouseDB.Table(tableNameOld).Select("start_time,dst_ip,host,app_id,bytes_up_view,bytes_dn_view").Where("start_time >= ? AND start_time < ?", param.StartTime, global.CONFIG.DeploymentDate)).
			Select(selectStr)
	case global.QueryOld:
		ckDb = ckV1Db.Select(selectStr).Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	default:
		return result, errors.New("开始、结束时间解析错误")
	}
	err := ckDb.Group("name,host").Limit(param.Limit).Offset(param.Limit * (param.Page - 1)).Order("total_byte DESC").Find(&result).Error
	if err != nil {
		global.Log.Error("获取用户行为分析详情数据失败", zap.Error(err))
		return result, err
	}
	return result, nil
}

func (service UserActionService) ExportData(param traffic.UserActionReqParam) ([]byte, error) {
	fields := []string{"name", "upByte", "dnByte", "totalByte"}
	headers := []string{"名称", "上行流量", "下行流量", "总流量"}
	param.Page = 1
	param.Limit = global.CONFIG.ExportLimit
	pageInfo, err := service.GetUserActionTable(param)
	if err != nil {
		return nil, err
	}
	list := pageInfo.List

	// 将查询结果转换为 []map[string]interface{} 格式
	dataList, ok := list.([]traffic.UserActionTable)
	if !ok {
		global.Log.Error(err.Error())
		return nil, err
	}
	dataMapList := make([]map[string]interface{}, len(dataList))
	for i, dataItem := range dataList {
		dataMapList[i] = map[string]interface{}{
			fields[0]: dataItem.Name,
			fields[1]: dataItem.UpByte,
			fields[2]: dataItem.DnByte,
			fields[3]: dataItem.TotalByte,
		}
	}
	return utils2.ExportToExcel(fields, headers, dataMapList)
}
