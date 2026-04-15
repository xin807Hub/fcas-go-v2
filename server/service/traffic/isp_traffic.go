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

	subSrcDbV2 := global.V2ClickhouseDB.Table(tableNameNew)
	subDstDbV2 := global.V2ClickhouseDB.Table(tableNameNew)
	subSrcDbV1 := global.V1ClickhouseDB.Table(tableNameOld)
	subDstDbV1 := global.V1ClickhouseDB.Table(tableNameOld)

	if param.IsOversea != nil {
		if *param.IsOversea == OverSea { // 查询国外运营商
			subSrcDbV2 = subSrcDbV2.Where("isp = '国外'")
			subDstDbV2 = subDstDbV2.Where("d_isp = '国外'")
			subSrcDbV1 = subSrcDbV1.Where("isp = '国外'")
			subDstDbV1 = subDstDbV1.Where("d_isp = '国外'")
		} else { // 国内
			subSrcDbV2 = subSrcDbV2.Where("isp != '国外'")
			subDstDbV2 = subDstDbV2.Where("d_isp != '国外'")
			subSrcDbV1 = subSrcDbV1.Where("isp != '国外'")
			subDstDbV1 = subDstDbV1.Where("d_isp != '国外'")
		}
	}

	if len(param.LinkIdList) > 0 {
		subSrcDbV2 = subSrcDbV2.Where("link_id IN (?)", param.LinkIdList)
		subDstDbV2 = subDstDbV2.Where("link_id IN (?)", param.LinkIdList)
		subSrcDbV1 = subSrcDbV1.Where("link_id IN (?)", param.LinkIdList)
		subDstDbV1 = subDstDbV1.Where("link_id IN (?)", param.LinkIdList)
	}

	if param.DstProvince != "" {
		subSrcDbV2 = subSrcDbV2.Where("dst_province = ?", param.DstProvince)
		subDstDbV2 = subDstDbV2.Where("dst_province = ?", param.DstProvince)
		subSrcDbV1 = subSrcDbV1.Where("dst_province = ?", param.DstProvince)
		subDstDbV1 = subDstDbV1.Where("dst_province = ?", param.DstProvince)
	}

	if len(param.UserIdList) > 0 {
		subSrcDbV2 = subSrcDbV2.Where("user_id IN (?)", param.UserIdList)
		subDstDbV2 = subDstDbV2.Where("d_user_id IN (?)", param.UserIdList)
		subSrcDbV1 = subSrcDbV1.Where("user_id IN (?)", param.UserIdList)
		subDstDbV1 = subDstDbV1.Where("d_user_id IN (?)", param.UserIdList)
	}
	if param.Isp != "" { // 2级排名传参
		subSrcDbV2 = subSrcDbV2.Where("isp = ?", param.Isp)
		subDstDbV2 = subDstDbV2.Where("d_isp = ?", param.Isp)
		subSrcDbV1 = subSrcDbV1.Where("isp = ?", param.Isp)
		subDstDbV1 = subDstDbV1.Where("d_isp = ?", param.Isp)
	}

	queryTimeRangeType := utils.GetDbTypeByTimeRange(param.StartTime, param.EndTime)

	subSrcSelectSql := "isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn, Round(if(isNaN(traffic_up), 0 , traffic_up) * 8 / {{.interval}}, 2) as traffic_up_bps, Round(if(isNaN(traffic_dn), 0 , traffic_dn) * 8 / {{.interval}}, 2) as traffic_dn_bps"
	subSrcSelectSql = strings.ReplaceAll(subSrcSelectSql, "{{.interval}}", strconv.Itoa(particle))
	subDstSelectSql := "d_isp AS isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn, Round(if(isNaN(traffic_up), 0 , traffic_up) * 8 / {{.interval}}, 2) as traffic_up_bps, Round(if(isNaN(traffic_dn), 0 , traffic_dn) * 8 / {{.interval}}, 2) as traffic_dn_bps"
	subDstSelectSql = strings.ReplaceAll(subDstSelectSql, "{{.interval}}", strconv.Itoa(particle))
	GroupBySql := "isp,start_time"

	subSrcDbV2 = subSrcDbV2.Select(subSrcSelectSql).Group(GroupBySql)
	subDstDbV2 = subDstDbV2.Select(subDstSelectSql).Group(GroupBySql)
	subSrcDbV1 = subSrcDbV1.Select(subSrcSelectSql).Group(GroupBySql)
	subDstDbV1 = subDstDbV1.Select(subDstSelectSql).Group(GroupBySql)

	selectSql := "isp,start_time,traffic_up,traffic_dn,traffic_up_bps,traffic_dn_bps"

	ckDbV2 := global.V2ClickhouseDB.
		Table("((?) Union Distinct (?))", subSrcDbV2, subDstDbV2).
		Select(selectSql)

	ckDbV1 := global.V1ClickhouseDB.
		Table("((?) Union Distinct (?))", subSrcDbV1, subDstDbV1).
		Select(selectSql)

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

// GetIspRankLevel1 1级运营商排名入口
func (service IspService) GetIspRankLevel1(param traffic.IspReqParam) (level1Data traffic.Level1Data, err error) {
	ckDb, _ := getIspDb(param)

	level1Data, err = GetGatherData(ckDb, level1Data, "isp")
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
	ckDb, _ := getIspDb(param)

	selectSql := "isp," +
		"max(traffic_up_bps) AS max_up_bps, " +
		"max(traffic_dn_bps) AS max_dn_bps, " +
		"avg(traffic_up_bps) AS avg_up_bps, " +
		"avg(traffic_dn_bps) AS avg_dn_bps, " +
		"sum(traffic_up) AS up_byte," +
		"sum(traffic_dn) AS dn_byte, " +
		"up_byte + dn_byte AS total_byte "

	err := global.V2ClickhouseDB.Table("(?)", ckDb).
		Select(selectSql).
		Group("isp").
		Count(&total).
		Limit(param.Limit).Offset((param.Page - 1) * param.Limit).Order("total_byte DESC").Find(&ispTables).Error
	if err != nil {
		global.Log.Error("获取运营商1级排名表格分页数据失败", zap.Error(err))
		return result, err
	}

	ispMap, mapErr := getIspMap() // 获取运营商-国内外类型键值对
	if mapErr == nil {            // 没有err就赋值
		if len(ispTables) > 0 {
			for i := range ispTables {
				if isOversea, ok := ispMap[ispTables[i].Isp]; ok {
					ispTables[i].IsOversea = isOversea
				}
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
	var ispMap = make(map[string]uint8)
	type ispOversea struct {
		Name      string
		IsOversea uint8
	}
	var ispList []ispOversea
	err := global.ServiceDB.Table("isp_view").Select("name,is_oversea").Group("name,is_oversea").Find(&ispList).Error
	if err != nil {
		return ispMap, err
	}

	for i := range ispList {
		ispMap[ispList[i].Name] = ispList[i].IsOversea
	}
	return ispMap, nil
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
	var result = response.PageResult{
		PageSize: param.Limit,
		CurrPage: param.Page,
	}
	ckDb, _ := getIspDb(param)
	var trendTableData []chart.FlowTrendDot
	var total int64
	selectSql := "start_time, " +
		"sum(traffic_dn_bps) AS dn_bps, " +
		"sum(traffic_up_bps) AS up_bps, " +
		"(dn_bps + up_bps)  AS total_bps"
	selectSql = strings.ReplaceAll(selectSql, "{{.interval}}", strconv.Itoa(600))

	err := global.V2ClickhouseDB.Table("(?)", ckDb).Select(selectSql).
		Group("start_time").
		Count(&total).
		Limit(param.Limit).
		Offset(param.Limit * (param.Page - 1)).Order("start_time desc").
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

	// 将查询结果转换为 []map[string]interface{} 格式
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
