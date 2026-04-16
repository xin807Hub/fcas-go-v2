package traffic

import (
	"errors"
	"fcas_server/global"
	"fcas_server/model/common/chart"
	"fcas_server/model/common/response"
	"fcas_server/model/traffic"
	"fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

const (
	NotOversea = 0
	OverSea    = 1
	Unknown    = 2
)

type IspService struct{}

var IspTrafficNewTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas_v2.dws_isp_10m",
	global.Interval1hParticle:  "bigdata_fcas_v2.dws_isp_1h",
	global.Interval1dParticle:  "bigdata_fcas_v2.dws_isp_1d",
}

var IspTrafficOldTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas.dws_isp_10min",
	global.Interval1hParticle:  "bigdata_fcas.dws_isp_hour",
	global.Interval1dParticle:  "bigdata_fcas.dws_isp_hour",
}

var IspOverseaType = map[uint8]string{
	0: "国内",
	1: "国外",
	2: "未知",
}

func getIspDb(param traffic.IspReqParam) (ckDb *gorm.DB, err error) {
	particle, err := utils.GetParticleByTimeRange(param.StartTime, param.EndTime)
	if err != nil {
		global.Log.Error("获取查询颗粒度失败：", zap.Error(err))
		return nil, err
	}

	tableNameNew := IspTrafficNewTableNameMap[particle]
	tableNameOld := IspTrafficOldTableNameMap[particle]

	// v2 新库在无 user 视角时，同一条流量两侧 ISP 都已落到同一行。
	// 这里必须按行选出唯一“对端 ISP”，否则双分支 union 会把同一批流量算两次。
	if len(param.UserIdList) == 0 {
		ckDbV2, buildErr := buildV2IspDbWithoutUserScope(tableNameNew, particle, param)
		if buildErr != nil {
			return nil, buildErr
		}

		queryTimeRangeType := utils.GetDbTypeByTimeRange(param.StartTime, param.EndTime)
		switch queryTimeRangeType {
		case global.QueryNew:
			return ckDbV2.Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime), nil
		case global.QueryCrossOld2New:
			ckDbV2 = ckDbV2.Where("start_time >= ? AND start_time < ?", global.CONFIG.DeploymentDate, param.EndTime)
			ckDbV1, oldErr := buildLegacyIspDb(tableNameOld, particle, param, global.V1ClickhouseDB)
			if oldErr != nil {
				return nil, oldErr
			}
			ckDbV1 = ckDbV1.Where("start_time >= ? AND start_time < ?", param.StartTime, global.CONFIG.DeploymentDate)
			return global.V2ClickhouseDB.Table("(? UNION ALL ?)", ckDbV2, ckDbV1), nil
		case global.QueryOld:
			ckDbV1, oldErr := buildLegacyIspDb(tableNameOld, particle, param, global.V1ClickhouseDB)
			if oldErr != nil {
				return nil, oldErr
			}
			return ckDbV1.Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime), nil
		default:
			return nil, errors.New("开始、结束时间解析错误")
		}
	}

	ckDbV2, err := buildLegacyIspDb(tableNameNew, particle, param, global.V2ClickhouseDB)
	if err != nil {
		return nil, err
	}
	ckDbV1, err := buildLegacyIspDb(tableNameOld, particle, param, global.V1ClickhouseDB)
	if err != nil {
		return nil, err
	}

	queryTimeRangeType := utils.GetDbTypeByTimeRange(param.StartTime, param.EndTime)
	switch queryTimeRangeType {
	case global.QueryNew:
		ckDb = ckDbV2.Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	case global.QueryCrossOld2New:
		ckDbV2 = ckDbV2.Where("start_time >= ? AND start_time < ?", global.CONFIG.DeploymentDate, param.EndTime)
		ckDbV1 = ckDbV1.Where("start_time >= ? AND start_time < ?", param.StartTime, global.CONFIG.DeploymentDate)
		ckDb = global.V2ClickhouseDB.Table("(? UNION ALL ?)", ckDbV2, ckDbV1)
	case global.QueryOld:
		ckDb = ckDbV1.Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime)
	default:
		return nil, errors.New("开始、结束时间解析错误")
	}
	return ckDb, err
}

func buildLegacyIspDb(tableName string, particle int, param traffic.IspReqParam, baseDB *gorm.DB) (*gorm.DB, error) {
	// Keep the ISP direction aligned with v1:
	// - user_id branch uses peer ISP d_isp
	// - d_user_id branch uses peer ISP isp
	subUserDb := applyNonEmptyIspFilter(baseDB.Table(tableName), "d_isp")
	subDUserDb := applyNonEmptyIspFilter(baseDB.Table(tableName), "isp")

	if param.IsOversea != nil {
		ispNames, lookupErr := getIspNamesByOversea(*param.IsOversea)
		if lookupErr != nil {
			return nil, lookupErr
		}
		subUserDb = applyIspNamesFilter(subUserDb, "d_isp", ispNames)
		subDUserDb = applyIspNamesFilter(subDUserDb, "isp", ispNames)
	}

	if len(param.LinkIdList) > 0 {
		subUserDb = subUserDb.Where("link_id IN (?)", param.LinkIdList)
		subDUserDb = subDUserDb.Where("link_id IN (?)", param.LinkIdList)
	}

	if param.DstProvince != "" {
		subUserDb = subUserDb.Where("dst_province = ?", param.DstProvince)
		subDUserDb = subDUserDb.Where("dst_province = ?", param.DstProvince)
	}

	if len(param.UserIdList) > 0 {
		subUserDb = subUserDb.Where("user_id IN (?)", param.UserIdList)
		subDUserDb = subDUserDb.Where("d_user_id IN (?)", param.UserIdList)
	}

	if len(param.IspNameList) > 0 {
		subUserDb = subUserDb.Where("d_isp IN (?)", param.IspNameList)
		subDUserDb = subDUserDb.Where("isp IN (?)", param.IspNameList)
	}

	if param.Isp != "" {
		subUserDb = subUserDb.Where("d_isp = ?", param.Isp)
		subDUserDb = subDUserDb.Where("isp = ?", param.Isp)
	}

	subUserSelectSQL := "d_isp AS isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn"
	subDUserSelectSQL := "isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn"

	subUserDb = subUserDb.Select(subUserSelectSQL).Group("d_isp,start_time")
	subDUserDb = subDUserDb.Select(subDUserSelectSQL).Group("isp,start_time")

	return buildIspTrafficQuery(baseDB, subUserDb, subDUserDb, particle, true), nil
}

func buildV2IspDbWithoutUserScope(tableName string, particle int, param traffic.IspReqParam) (*gorm.DB, error) {
	selectedIspExpr := "if(user_id != 0, d_isp, isp)"
	baseDb := applyNonEmptyIspExprFilter(global.V2ClickhouseDB.Table(tableName), selectedIspExpr)

	if param.IsOversea != nil {
		ispNames, err := getIspNamesByOversea(*param.IsOversea)
		if err != nil {
			return nil, err
		}
		baseDb = applyIspNamesExprFilter(baseDb, selectedIspExpr, ispNames)
	}

	if len(param.LinkIdList) > 0 {
		baseDb = baseDb.Where("link_id IN (?)", param.LinkIdList)
	}

	if param.DstProvince != "" {
		baseDb = baseDb.Where("dst_province = ?", param.DstProvince)
	}

	if len(param.IspNameList) > 0 {
		baseDb = applyIspNamesExprFilter(baseDb, selectedIspExpr, param.IspNameList)
	}

	if param.Isp != "" {
		baseDb = baseDb.Where(fmt.Sprintf("%s = ?", selectedIspExpr), param.Isp)
	}

	materializedQuery := global.V2ClickhouseDB.Table("(?)", baseDb.Select(
		fmt.Sprintf("%s AS selected_isp,start_time,bytes_up_view,bytes_dn_view", selectedIspExpr),
	))
	sourceQuery := materializedQuery.
		Select("selected_isp AS isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn").
		Group("selected_isp,start_time")
	return buildIspTrafficQuery(global.V2ClickhouseDB, sourceQuery, nil, particle, false), nil
}

// GetIspRankLevel1 1级运营商排名入口
func (service IspService) GetIspRankLevel1(param traffic.IspReqParam) (level1Data traffic.Level1Data, err error) {
	ckDb, err := getIspDb(param)
	if err != nil {
		return level1Data, err
	}
	totalSpanSeconds, err := getQuerySpanSeconds(param.StartTime, param.EndTime)
	if err != nil {
		return level1Data, err
	}

	level1Data, err = GetGatherData(ckDb, level1Data, "isp", totalSpanSeconds)
	if err != nil {
		global.Log.Error("获取运营商1级汇总数据错误", zap.Error(err))
		return level1Data, err
	}

	level1Data, err = GetPieData(ckDb, level1Data, "isp")
	if err != nil {
		global.Log.Error("获取运营商1级饼图数据错误", zap.Error(err))
		return level1Data, err
	}
	return level1Data, nil
}

func (service IspService) GetLevel1TableData(param traffic.IspReqParam) (response.PageResult, error) {
	var result response.PageResult
	var ispTables []traffic.IspLevel1TableData
	var total int64
	if param.Limit == 0 {
		param.Limit = 10
	}

	ckDb, err := getIspDb(param)
	if err != nil {
		return result, err
	}
	totalSpanSeconds, err := getQuerySpanSeconds(param.StartTime, param.EndTime)
	if err != nil {
		return result, err
	}
	err = buildIspLevel1TableQuery(ckDb, totalSpanSeconds).
		Count(&total).
		Limit(param.Limit).
		Offset((param.Page - 1) * param.Limit).
		Order("total_byte DESC").
		Find(&ispTables).Error
	if err != nil {
		global.Log.Error("获取运营商1级排名表格分页数据失败", zap.Error(err))
		return result, err
	}

	ispMap, mapErr := getIspMap()
	if mapErr == nil && len(ispTables) > 0 {
		for i := range ispTables {
			if isOversea, ok := ispMap[ispTables[i].Isp]; ok {
				ispTables[i].IsOversea = isOversea
			}
		}
	}

	result.TotalCount = total
	result.List = ispTables
	result.PageSize = param.Limit
	result.CurrPage = param.Page
	return result, nil
}

func getIspMap() (map[string]uint8, error) {
	ispMap := make(map[string]uint8)
	ispList, err := listIspOversea()
	if err != nil {
		return ispMap, err
	}

	for i := range ispList {
		ispMap[ispList[i].Name] = ispList[i].IsOversea
	}
	return ispMap, nil
}

type ispOversea struct {
	Name      string
	IsOversea uint8
}

func listIspOversea() ([]ispOversea, error) {
	var ispList []ispOversea
	err := global.ServiceDB.Table("isp_view").Select("name,is_oversea").Group("name,is_oversea").Find(&ispList).Error
	if err != nil {
		return nil, err
	}
	return ispList, nil
}

func getIspNamesByOversea(isOversea uint8) ([]string, error) {
	ispList, err := listIspOversea()
	if err != nil {
		return nil, err
	}

	ispNames := make([]string, 0, len(ispList))
	for _, item := range ispList {
		if item.IsOversea == isOversea {
			ispNames = append(ispNames, item.Name)
		}
	}
	return ispNames, nil
}

func applyIspNamesFilter(db *gorm.DB, column string, ispNames []string) *gorm.DB {
	if len(ispNames) == 0 {
		return db.Where("1 = 0")
	}
	return db.Where(fmt.Sprintf("%s IN (?)", column), ispNames)
}

func applyIspNamesExprFilter(db *gorm.DB, expr string, ispNames []string) *gorm.DB {
	if len(ispNames) == 0 {
		return db.Where("1 = 0")
	}
	return db.Where(fmt.Sprintf("%s IN (?)", expr), ispNames)
}

func applyNonEmptyIspFilter(db *gorm.DB, column string) *gorm.DB {
	return db.Where(fmt.Sprintf("ifNull(%s, '') != ''", column))
}

func applyNonEmptyIspExprFilter(db *gorm.DB, expr string) *gorm.DB {
	return db.Where(fmt.Sprintf("ifNull(%s, '') != ''", expr))
}

func buildIspTrafficQuery(baseDB *gorm.DB, subUserDb *gorm.DB, subDUserDb *gorm.DB, particle int, useDualBranch bool) *gorm.DB {
	unionSelectSQL := "isp,start_time,traffic_up,traffic_dn"
	aggregateSelectSQL := "" +
		"isp,start_time," +
		"sum(traffic_up) AS grouped_traffic_up," +
		"sum(traffic_dn) AS grouped_traffic_dn," +
		"Round(if(isNaN(sum(traffic_up)), 0, sum(traffic_up)) * 8 / {{.interval}}, 2) AS grouped_traffic_up_bps," +
		"Round(if(isNaN(sum(traffic_dn)), 0, sum(traffic_dn)) * 8 / {{.interval}}, 2) AS grouped_traffic_dn_bps"
	aggregateSelectSQL = strings.ReplaceAll(aggregateSelectSQL, "{{.interval}}", strconv.Itoa(particle))
	finalSelectSQL := "" +
		"isp,start_time," +
		"grouped_traffic_up AS traffic_up," +
		"grouped_traffic_dn AS traffic_dn," +
		"grouped_traffic_up_bps AS traffic_up_bps," +
		"grouped_traffic_dn_bps AS traffic_dn_bps"

	sourceQuery := subUserDb
	if useDualBranch {
		sourceQuery = baseDB.Table("((?) Union Distinct (?))", subUserDb, subDUserDb).Select(unionSelectSQL)
	}

	ckDb := baseDB.Table("(?)", sourceQuery).
		Select(aggregateSelectSQL).
		Group("isp,start_time")
	ckDb = baseDB.Table("(?)", ckDb).Select(finalSelectSQL)
	return applyNonEmptyIspFilter(ckDb, "isp")
}

func buildIspLevel1TableQuery(ckDb *gorm.DB, totalSpanSeconds int) *gorm.DB {
	selectSQL := "isp," +
		"max(traffic_up_bps) AS max_up_bps, " +
		"max(traffic_dn_bps) AS max_dn_bps, " +
		buildLevel1AverageExpr("sum(traffic_up)", totalSpanSeconds, "avg_up_bps") + ", " +
		buildLevel1AverageExpr("sum(traffic_dn)", totalSpanSeconds, "avg_dn_bps") + ", " +
		"sum(traffic_up) AS up_byte," +
		"sum(traffic_dn) AS dn_byte, " +
		"up_byte + dn_byte AS total_byte "

	return applyNonEmptyIspFilter(global.V2ClickhouseDB.Table("(?)", ckDb).Select(selectSQL), "isp").
		Group("isp")
}

// GetIspRankLevel2 2级运营商排名入口
func (service IspService) GetIspRankLevel2(param traffic.IspReqParam) ([]chart.FlowTrendDot, error) {
	ckDb, _ := getIspDb(param)
	trendSeries, err := GetLevel2TrendData(ckDb)
	if err != nil {
		global.Log.Error("获取运营商2级趋势图数据错误", zap.Error(err))
	}
	return trendSeries, err
}

func (service IspService) GetLevel2TableData(param traffic.IspReqParam) (response.PageResult, error) {
	if param.Limit == 0 {
		param.Limit = 10
	}

	result := response.PageResult{
		PageSize: param.Limit,
		CurrPage: param.Page,
	}
	ckDb, _ := getIspDb(param)

	var trendTableData []chart.FlowTrendDot
	var total int64
	selectSQL := "start_time, " +
		"sum(traffic_dn_bps) AS dn_bps, " +
		"sum(traffic_up_bps) AS up_bps, " +
		"(dn_bps + up_bps) AS total_bps"
	selectSQL = strings.ReplaceAll(selectSQL, "{{.interval}}", strconv.Itoa(600))

	err := global.V2ClickhouseDB.Table("(?)", ckDb).Select(selectSQL).
		Group("start_time").
		Count(&total).
		Limit(param.Limit).
		Offset(param.Limit * (param.Page - 1)).
		Order("start_time desc").
		Find(&trendTableData).
		Error
	if err != nil {
		return result, err
	}

	result.TotalCount = total
	result.List = trendTableData
	return result, nil
}

func (service IspService) ExportData(param traffic.IspReqParam) ([]byte, error) {
	fields := []string{"isp", "isOversea", "maxUpBps", "maxDnBps", "avgUpBps", "avgDnBps", "upByte", "dnByte", "totalByte", "proportion"}
	headers := []string{"运营商", "运营商国内外类型", "上行峰值(Mbps)", "下行峰值(Mbps)", "上行平均(Mbps)", "下行平均(Mbps)", "上行总量(MB)", "下行总量(MB)", "总流量(MB)", "总量占比"}
	param.Page = 1
	param.Limit = global.CONFIG.ExportLimit

	pageInfo, err := service.GetLevel1TableData(param)
	if err != nil {
		return nil, err
	}

	ckDb, _ := getIspDb(param)
	totalByte, err := GetTotalByte(ckDb)
	if err != nil {
		return nil, err
	}

	list := pageInfo.List
	dataList, ok := list.([]traffic.IspLevel1TableData)
	if !ok {
		global.Log.Error(err.Error())
		return nil, err
	}

	dataMapList := make([]map[string]interface{}, len(dataList))
	for i, dataItem := range dataList {
		dataMapList[i] = map[string]interface{}{
			fields[0]: dataItem.Isp,
			fields[1]: IspOverseaType[dataItem.IsOversea],
			fields[2]: fmt.Sprintf("%.2f", float64(dataItem.MaxUpBps)/1000/1000),
			fields[3]: fmt.Sprintf("%.2f", float64(dataItem.MaxDnBps)/1000/1000),
			fields[4]: fmt.Sprintf("%.2f", float64(dataItem.AvgUpBps)/1000/1000),
			fields[5]: fmt.Sprintf("%.2f", float64(dataItem.AvgDnBps)/1000/1000),
			fields[6]: fmt.Sprintf("%.2f", float64(dataItem.UpByte)/1000/1000),
			fields[7]: fmt.Sprintf("%.2f", float64(dataItem.DnByte)/1000/1000),
			fields[8]: fmt.Sprintf("%.2f", float64(dataItem.TotalByte)/1000/1000),
			fields[9]: fmt.Sprintf("%.4f%%", float64(dataItem.TotalByte)*100/float64(totalByte)),
		}
	}

	return utils.ExportToExcel(fields, headers, dataMapList)
}
