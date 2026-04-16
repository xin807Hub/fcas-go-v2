package traffic

import (
	"errors"
	"fcas_server/global"
	"fcas_server/model/common/chart"
	"fcas_server/model/common/response"
	"fcas_server/model/traffic"
	utils2 "fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type AppTypeService struct{}

var AppTypeNewTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas_v2.dws_app_type_10m",
	global.Interval1hParticle:  "bigdata_fcas_v2.dws_app_type_1h",
	global.Interval1dParticle:  "bigdata_fcas_v2.dws_app_type_1d",
}

var AppTypeOldTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas.dws_app_type_10min",
	global.Interval1hParticle:  "bigdata_fcas.dws_app_type_hour",
	global.Interval1dParticle:  "bigdata_fcas.dws_app_type_hour",
}

func getAppTypeDb(param traffic.AppTypeReqParam) (ckDb *gorm.DB, err error) {
	particle, err := utils2.GetParticleByTimeRange(param.StartTime, param.EndTime)

	tableNameNew := AppTypeNewTableNameMap[particle]
	tableNameOld := AppTypeOldTableNameMap[particle]

	ckV2Db := global.V2ClickhouseDB.Table(tableNameNew)
	ckV1Db := global.V1ClickhouseDB.Table(tableNameOld)

	if len(param.IspNameList) > 0 {
		ckV2Db = ckV2Db.Where("( isp IN (?) OR d_isp IN (?) )", param.IspNameList, param.IspNameList)
		ckV1Db = ckV1Db.Where("( isp IN (?) OR d_isp IN (?) )", param.IspNameList, param.IspNameList)
	}
	if len(param.LinkIdList) > 0 {
		ckV2Db = ckV2Db.Where("link_id IN (?) ", param.LinkIdList)
		ckV1Db = ckV1Db.Where("link_id IN (?) ", param.LinkIdList)
	}
	if param.DstProvince != "" {
		ckV2Db = ckV2Db.Where("dst_province = ?", param.DstProvince)
		ckV1Db = ckV1Db.Where("dst_province = ?", param.DstProvince) // string
	}
	if len(param.UserIdList) > 0 {
		ckV2Db = ckV2Db.Where("(user_id IN (?) OR d_user_id IN (?)) ", param.UserIdList, param.UserIdList)
		ckV1Db = ckV1Db.Where("(user_id IN (?) OR d_user_id IN (?)) ", param.UserIdList, param.UserIdList)
	}
	if len(param.AppTypeIdList) > 0 {
		if len(param.AppTypeIdList) == 1 {
			ckV2Db = ckV2Db.Where("app_type = ?", param.AppTypeIdList[0])
			ckV1Db = ckV1Db.Where("app_type = ?", param.AppTypeIdList[0])
		}
		ckV2Db = ckV2Db.Where("app_type IN (?) ", param.AppTypeIdList)
		ckV1Db = ckV1Db.Where("app_type IN (?) ", param.AppTypeIdList)
	}

	timeRangeType := utils2.GetDbTypeByTimeRange(param.StartTime, param.EndTime)
	switch timeRangeType {
	case global.QueryNew:
		ckDb = ckV2Db.Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	case global.QueryCrossOld2New:
		ckDbNew := ckV2Db
		ckDbOld := ckV1Db
		ckDb = global.V2ClickhouseDB.Table("(? UNION ALL ?)",
			ckDbNew.Table(tableNameNew).
				Select("app_type,start_time,bytes_up_view,bytes_dn_view").
				Where("start_time >= ? AND start_time < ?", global.CONFIG.DeploymentDate, param.EndTime),
			ckDbOld.Table(tableNameOld).
				Select("app_type,start_time,bytes_up_view,bytes_dn_view").
				Where("start_time >= ? AND start_time < ?", param.StartTime, global.CONFIG.DeploymentDate))
	case global.QueryOld:
		ckDb = ckV1Db.Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	default:
		return nil, errors.New("开始、结束时间解析错误")
	}

	selectSql := "app_type,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn, Round(if(isNaN(traffic_up), 0 , traffic_up) * 8 / {{.interval}}, 2) as traffic_up_bps, Round(if(isNaN(traffic_dn), 0 , traffic_dn) * 8 / {{.interval}}, 2) as traffic_dn_bps"
	selectSql = strings.ReplaceAll(selectSql, "{{.interval}}", strconv.Itoa(particle))

	ckDb = ckDb.Select(selectSql).Group("app_type,start_time")
	return ckDb, err
}

func (service AppTypeService) GetAppTypeRankLevel1(param traffic.AppTypeReqParam) (traffic.Level1Data, error) {
	var level1Data = traffic.Level1Data{}
	ckDb, err := getAppTypeDb(param)
	if err != nil {
		return level1Data, err
	}
	totalSpanSeconds, err := getQuerySpanSeconds(param.StartTime, param.EndTime)
	if err != nil {
		return level1Data, err
	}

	level1Data, err = GetGatherData(ckDb, level1Data, "app_type", totalSpanSeconds)
	if err != nil {
		global.Log.Error("获取app大类1级汇总数据错误", zap.Error(err))
		return level1Data, err
	}

	level1Data, err = GetPieData(ckDb, level1Data, "appType")
	if err != nil {
		global.Log.Error("获取app大类1级饼图数据错误", zap.Error(err))
		return level1Data, err
	}

	return level1Data, nil
}

func (service AppTypeService) GetLevel1TableData(param traffic.AppTypeReqParam) (response.PageResult, error) {
	var result response.PageResult
	var level1Tables []traffic.AppTypeLevel1TableData
	ckDb, err := getAppTypeDb(param)
	if err != nil {
		return result, err
	}
	if param.Limit == 0 {
		param.Limit = 10
	}
	totalSpanSeconds, err := getQuerySpanSeconds(param.StartTime, param.EndTime)
	if err != nil {
		return result, err
	}
	selectStr := "app_type, " +
		"dictGet('app_type_dict','name',app_type) AS name," +
		"max(traffic_up_bps) AS max_up_bps, " +
		"max(traffic_dn_bps) AS max_dn_bps, " +
		buildLevel1AverageExpr("sum(traffic_up)", totalSpanSeconds, "avg_up_bps") + ", " +
		buildLevel1AverageExpr("sum(traffic_dn)", totalSpanSeconds, "avg_dn_bps") + ", " +
		"sum(traffic_up) AS up_byte," +
		"sum(traffic_dn) AS dn_byte," +
		"up_byte + dn_byte AS total_byte"
	var total int64
	db := global.V2ClickhouseDB.Table("(?)", ckDb).Select(selectStr).Group("app_type")
	err = global.V2ClickhouseDB.Table("(?)", db).Count(&total).Error
	if err != nil {
		global.Log.Error("", zap.Error(err))
		return result, err
	}

	err = db.Limit(param.Limit).Offset(param.Limit * (param.Page - 1)).Order(" total_byte DESC ").Find(&level1Tables).Error
	if err != nil {
		return result, err
	}

	result.List = level1Tables
	result.TotalCount = total
	result.PageSize = param.Limit
	result.CurrPage = param.Page
	return result, nil
}

// GetAppTypeRankLevel2 2级排名入口
func (service AppTypeService) GetAppTypeRankLevel2(param traffic.AppTypeReqParam) ([]chart.FlowTrendDot, error) {
	ckDb, _ := getAppTypeDb(param)
	trendSeries, err := GetLevel2TrendData(ckDb)
	if err != nil {
		global.Log.Error("获取运营商二级趋势图/表格数据错误", zap.Error(err))
	}
	return trendSeries, err
}

func (service AppTypeService) GetLevel2TableData(param traffic.AppTypeReqParam) (response.PageResult, error) {
	if param.Limit == 0 {
		param.Limit = 10
	}
	var result = response.PageResult{
		PageSize: param.Limit,
		CurrPage: param.Page,
	}
	ckDb, err := getAppTypeDb(param)
	if err != nil {
		return result, errors.New("根据查询条件获取db句柄出错")
	}
	return GetLevel2TableData(ckDb, param.PageInfo)
}

func (service AppTypeService) ExportData(param traffic.AppTypeReqParam) ([]byte, error) {
	fields := []string{"name", "maxDnBps", "maxUpBps", "avgUpBps", "avgDnBps", "upByte", "dnByte", "totalByte", "proportion"}
	headers := []string{"业务大类", "上行峰值(Mbps)", "下行峰值(Mbps)", "上行平均(Mbps)", "下行平均(Mbps)", "上行总量(MB)", "下行总量(MB)", "总流量(MB)", "总量占比"}
	param.Page = 1
	param.Limit = global.CONFIG.ExportLimit
	pageInfo, err := service.GetLevel1TableData(param)
	if err != nil {
		return nil, err
	}
	ckDb, _ := getAppTypeDb(param)
	totalByte, err := GetTotalByte(ckDb)
	if err != nil {
		return nil, err
	}
	list := pageInfo.List

	// 将查询结果转换为 []map[string]interface{} 格式
	dataList, ok := list.([]traffic.AppTypeLevel1TableData)
	if !ok {
		global.Log.Error(err.Error())
		return nil, err
	}
	dataMapList := make([]map[string]interface{}, len(dataList))
	for i, dataItem := range dataList {
		dataMapList[i] = map[string]interface{}{
			fields[0]: dataItem.Name,
			fields[1]: fmt.Sprintf("%.2f", float64(dataItem.MaxUpBps)/1000/1000),
			fields[2]: fmt.Sprintf("%.2f", float64(dataItem.MaxDnBps)/1000/1000),
			fields[3]: fmt.Sprintf("%.2f", float64(dataItem.AvgUpBps)/1000/1000),
			fields[4]: fmt.Sprintf("%.2f", float64(dataItem.AvgDnBps)/1000/1000),
			fields[5]: fmt.Sprintf("%.2f", float64(dataItem.UpByte)/1000/1000),
			fields[6]: fmt.Sprintf("%.2f", float64(dataItem.DnByte)/1000/1000),
			fields[7]: fmt.Sprintf("%.2f", float64(dataItem.TotalByte)/1000/1000),
			fields[8]: fmt.Sprintf("%.4f%%", float64(dataItem.TotalByte)*100/float64(totalByte)),
		}
	}
	return utils2.ExportToExcel(fields, headers, dataMapList)
}
