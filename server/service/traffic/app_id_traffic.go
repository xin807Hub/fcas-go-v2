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

type AppIdService struct{}

var AppIdNewTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas_v2.dws_user_10m",
	global.Interval1hParticle:  "bigdata_fcas_v2.dws_user_1h",
	global.Interval1dParticle:  "bigdata_fcas_v2.dws_user_1d",
}
var AppIdOldTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas.dws_user_10min",
	global.Interval1hParticle:  "bigdata_fcas.dws_user_hour",
	global.Interval1dParticle:  "bigdata_fcas.dws_user_hour",
}

var SeverNewTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas_v2.dws_server_10m",
	global.Interval1hParticle:  "bigdata_fcas_v2.dws_server_1h",
	global.Interval1dParticle:  "bigdata_fcas_v2.dws_server_1d",
}
var SeverOldTableNameMap = map[int]string{
	global.Interval10mParticle: "bigdata_fcas.dws_server_10min",
	global.Interval1hParticle:  "bigdata_fcas.dws_server_hour",
	global.Interval1dParticle:  "bigdata_fcas.dws_server_hour",
}

func getAppIdDb(param traffic.AppIdReqParam) (ckDb *gorm.DB, err error) {
	particle, err := utils2.GetParticleByTimeRange(param.StartTime, param.EndTime)
	if err != nil {
		return nil, err
	}
	tableNameNew := AppIdNewTableNameMap[particle]
	tableNameOld := AppIdOldTableNameMap[particle]

	selectSql := "app_id,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn, Round(if(isNaN(traffic_up), 0 , traffic_up) * 8 / {{.interval}}, 2) as traffic_up_bps, Round(if(isNaN(traffic_dn), 0 , traffic_dn) * 8 / {{.interval}}, 2) as traffic_dn_bps"
	groupSql := "app_id,start_time"

	if param.RankLevel == global.LevelTwo || param.RankLevel == global.LevelThree {
		if param.RankLevel == global.LevelTwo && len(param.AppIdList) == 0 {
			return ckDb, errors.New("请求2级排名数据时，[appIdList]不得为空")
		}
		tableNameNew = SeverNewTableNameMap[particle]
		tableNameOld = SeverOldTableNameMap[particle]

		if param.RankType == traffic.RankTypeDstIP {
			selectSql = "app_id,dst_ip,dst_area_id,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn,Round(if(isNaN(traffic_up), 0 , traffic_up) * 8 / {{.interval}}, 2) as traffic_up_bps, Round(if(isNaN(traffic_dn), 0 , traffic_dn) * 8 / {{.interval}}, 2) as traffic_dn_bps"
			groupSql = "app_id,dst_ip,dst_area_id,start_time"
		} else if param.RankType == traffic.RankTypeHost {
			selectSql = "app_id,host,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn,Round(if(isNaN(traffic_up), 0 , traffic_up) * 8 / {{.interval}}, 2) as traffic_up_bps, Round(if(isNaN(traffic_dn), 0 , traffic_dn) * 8 / {{.interval}}, 2) as traffic_dn_bps"
			groupSql = "app_id,host,start_time"
		}
	}

	selectSql = strings.ReplaceAll(selectSql, "{{.interval}}", strconv.Itoa(particle))

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
		ckV1Db = ckV1Db.Where("dst_province = ?", param.DstProvince)
	}
	if len(param.UserIdList) > 0 {
		ckV2Db = ckV2Db.Where("(user_id IN (?) OR d_user_id IN (?)) ", param.UserIdList, param.UserIdList)
		ckV1Db = ckV1Db.Where("(user_id IN (?) OR d_user_id IN (?)) ", param.UserIdList, param.UserIdList)
	}
	if len(param.AppIdList) > 0 {
		if len(param.AppIdList) == 1 {
			ckV2Db = ckV2Db.Where("app_id = ?", param.AppIdList[0])
			ckV1Db = ckV1Db.Where("app_id = ?", param.AppIdList[0])
		} else {
			ckV2Db = ckV2Db.Where("app_id IN (?) ", param.AppIdList)
			ckV1Db = ckV1Db.Where("app_id IN (?) ", param.AppIdList)
		}
	}
	if param.RankLevel == global.LevelThree {
		if param.DstIpParam != "" {
			ckV2Db = ckV2Db.Where("dst_ip = ?", param.DstIpParam)
			ckV1Db = ckV1Db.Where("dst_ip = ?", param.DstIpParam)
		}
		if param.HostParam != "" {
			ckV2Db = ckV2Db.Where("host = ?", param.HostParam)
			ckV1Db = ckV1Db.Where("host = ?", param.HostParam)
		}
	}
	if param.RankType == traffic.RankTypeHost {
		ckV2Db = ckV2Db.Where("host != '' ")
		ckV1Db = ckV1Db.Where("host != '' ")
	}

	timeRangeType := utils2.GetDbTypeByTimeRange(param.StartTime, param.EndTime)
	switch timeRangeType {
	case global.QueryNew:
		ckDb = ckV2Db.
			Select(selectSql).
			Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime).
			Group(groupSql)
	case global.QueryCrossOld2New:
		ckDbNew := ckV2Db
		ckDbOld := ckV1Db
		ckDb = global.V2ClickhouseDB.Table("(? Union Distinct ?)",
			ckDbNew.Table(tableNameNew).Select(selectSql).Where("start_time >= ? AND start_time < ?", global.CONFIG.DeploymentDate, param.EndTime).Group(groupSql),
			ckDbOld.Table(tableNameOld).Select(selectSql).Where("start_time >= ? AND start_time < ?", param.StartTime, global.CONFIG.DeploymentDate).Group(groupSql))
	case global.QueryOld:
		ckDb = ckV1Db.
			Select(selectSql).
			Where("start_time >= ? AND start_time < ?", param.StartTime, param.EndTime).
			Group(groupSql)
	default:
		return nil, errors.New("开始、结束时间解析错误")
	}

	return ckDb, err
}

// GetAppIdRankLevel1 1级小类排名入口
func (service AppIdService) GetAppIdRankLevel1(param traffic.AppIdReqParam) (traffic.Level1Data, error) {
	var level1Data = traffic.Level1Data{}
	ckDb, err := getAppIdDb(param)

	level1Data, err = GetGatherData(ckDb, level1Data, "app_id")
	if err != nil {
		global.Log.Error("获取app大类一级汇总数据错误", zap.Error(err))
		return level1Data, err
	}

	level1Data, err = GetPieData(ckDb, level1Data, "appId")
	if err != nil {
		global.Log.Error("获取app大类一级饼图数据错误", zap.Error(err))
		return level1Data, err
	}

	return level1Data, nil
}

// GetLevel1TableData 获取应用小类表（分页）
func (service AppIdService) GetLevel1TableData(param traffic.AppIdReqParam) (response.PageResult, error) {
	var result response.PageResult
	var level1Tables []traffic.AppIdLevel1TableData
	var total int64
	ckDb, err := getAppIdDb(param)

	selectStr := "app_id," +
		"dictGet('app_id_dict','name',app_id) AS name," +
		"max(traffic_up_bps) AS max_up_bps, " +
		"max(traffic_dn_bps) AS max_dn_bps, " +
		"avg(traffic_up_bps) AS avg_up_bps, " +
		"avg(traffic_dn_bps) AS avg_dn_bps, " +
		"sum(traffic_up) AS up_byte," +
		"sum(traffic_dn) AS dn_byte," +
		"up_byte + dn_byte AS total_byte"

	db := global.V2ClickhouseDB.Table("(?)", ckDb).Select(selectStr).Group("app_id")
	err = global.V2ClickhouseDB.Table("(?)", db).Count(&total).Error
	if err != nil {
		global.Log.Error("获取总数失败", zap.Error(err))
		return result, err
	}
	err = db.Limit(param.Limit).Offset(param.Limit * (param.Page - 1)).Order(" total_byte DESC ").Find(&level1Tables).Error
	if err != nil {
		global.Log.Error("获取分页数据失败", zap.Error(err))
		return result, err
	}

	result.List = level1Tables
	result.TotalCount = total
	result.PageSize = param.Limit
	result.CurrPage = param.Page
	return result, nil
}

// GetLevel2Or3TrendTable 2级：获取某个应用小类的趋势图下的表格数据 3级：获取某应用小类下的目的ip或者域名的表格数据
func (service AppIdService) GetLevel2Or3TrendTable(param traffic.AppIdReqParam) (response.PageResult, error) {
	if param.Limit == 0 {
		param.Limit = 10
	}
	var result = response.PageResult{
		PageSize: param.Limit,
		CurrPage: param.Page,
	}
	if param.RankLevel == global.LevelTwo && len(param.AppIdList) == 0 {
		return result, errors.New("请求2级排名数据时，[appIdList]不得为空")
	}
	if param.RankLevel == global.LevelThree && param.DstIpParam == "" && param.HostParam == "" {
		return result, errors.New("请求3级排名数据时，[dstIpParam]或[hostParam]不得为空")
	}
	ckDb, err := getAppIdDb(param)
	if err != nil {
		return result, errors.New("根据查询条件获取db句柄出错")
	}
	return GetLevel2TableData(ckDb, param.PageInfo)
}

// GetAppIdRankLevel2 2级小类排名入口
func (service AppIdService) GetAppIdRankLevel2(param traffic.AppIdReqParam) (traffic.AppIdLevel2Data, error) {
	var result traffic.AppIdLevel2Data

	ckDb, err := getAppIdDb(param)
	if err != nil {
		return result, err
	}

	trendSeries, err := GetLevel2TrendData(ckDb)
	if err != nil {
		global.Log.Error("获取应用小类二级趋势图/表格数据错误", zap.Error(err))
		return result, err
	}
	result.TrendData = trendSeries

	if param.TopN == 0 {
		param.TopN = 20
	}

	param.RankType = traffic.RankTypeDstIP
	ckDb, err = getAppIdDb(param)
	if err != nil {
		return result, err
	}
	dstIpRankData, err := getLevel2RankData(ckDb, param.TopN, traffic.RankTypeDstIP)
	if err != nil {
		global.Log.Error("获取应用小类2级目的ip排名数据错误", zap.Error(err))
		return result, err
	}
	result.DstIpRankBar = dstIpRankData

	param.RankType = traffic.RankTypeHost
	ckDb, err = getAppIdDb(param)
	if err != nil {
		return result, err
	}
	hostRankData, err := getLevel2RankData(ckDb, param.TopN, traffic.RankTypeHost)
	if err != nil {
		global.Log.Error("获取应用小类2级host排名数据错误", zap.Error(err))
		return result, err
	}
	result.HostRankBar = hostRankData

	return result, nil
}

// 2级：获取该小类的目的ip和host排名数据 (使用均值bps进行排名)
func getLevel2RankData(ckDb *gorm.DB, topN uint8, rankType string) ([]chart.FlowBarDot, error) {
	var rankBarData []chart.FlowBarDot
	var selectStr string
	if rankType == traffic.RankTypeDstIP {
		selectStr = "dst_ip AS name,sum(traffic_up_bps+traffic_dn_bps) AS value"
	} else if rankType == traffic.RankTypeHost {
		selectStr = "host AS name,sum(traffic_up_bps+traffic_dn_bps) AS value"
	}
	err := global.V2ClickhouseDB.Table("(?)", ckDb).
		Select(selectStr).
		Group("name").
		Order("value DESC ").
		Limit(int(topN)).
		Find(&rankBarData).Error
	if err != nil {
		return rankBarData, err
	}
	return rankBarData, nil
}

// GetLevel2RankTable 2级：获取该小类应用的目的ip或host维度的排名表格数据（带分页)
func (service AppIdService) GetLevel2RankTable(param traffic.AppIdReqParam) (response.PageResult, error) {
	if param.Limit == 0 {
		param.Limit = 10
	}
	var result = response.PageResult{
		PageSize: param.Limit,
		CurrPage: param.Page,
	}
	if param.RankType == "" {
		return result, errors.New("rankType不得为空")
	}
	ckDb, err := getAppIdDb(param)
	if err != nil {
		global.Log.Error(err.Error(), zap.Error(err))
		return result, err
	}
	var level2RankTableData []traffic.AppIdLevel2RankTable
	var total int64
	var selectStr = "dst_ip AS name," +
		"if(dst_area_id=0,'未知', dictGet('ip_address_dict','location',dst_area_id)) as location," +
		"sum(traffic_up) AS up_byte," +
		"sum(traffic_dn) AS dn_byte," +
		"(up_byte + dn_byte) AS total_byte," +
		"avg(traffic_up_bps) AS avg_up_bps, " +
		"avg(traffic_dn_bps) AS avg_dn_bps"
	var groupBySql = "name,dst_area_id"

	if param.RankType == traffic.RankTypeHost {
		selectStr = "host AS name," +
			"sum(traffic_up) AS up_byte," +
			"sum(traffic_dn) AS dn_byte," +
			"(up_byte + dn_byte) AS total_byte," +
			"avg(traffic_up_bps) AS avg_up_bps, " +
			"avg(traffic_dn_bps) AS avg_dn_bps"
		groupBySql = "name"
	}

	db := global.V2ClickhouseDB.Table("(?)", ckDb).
		Select(selectStr).
		Group(groupBySql)
	err = global.V2ClickhouseDB.Table("(?)", db).Count(&total).Error
	if err != nil {
		global.Log.Error(err.Error())
		return result, err
	}
	err = db.Limit(param.Limit).Offset(param.Limit * (param.Page - 1)).Order("total_byte DESC").Find(&level2RankTableData).Error
	if err != nil {
		global.Log.Error(err.Error())
		return result, err
	}

	//if param.RankType == traffic.RankTypeDstIP && total > 0 {
	//	for i, item := range level2RankTableData {
	//		var ipAddrInfo traffic.BizIpAddress
	//		err = global.ServiceDB.Model(&traffic.BizIpAddress{}).Where("id", item.DstAreaId).First(&ipAddrInfo).Error
	//		if err != nil {
	//			continue
	//		}
	//		location := strings.Join([]string{ipAddrInfo.Country, ipAddrInfo.Province, ipAddrInfo.City, ipAddrInfo.Isp}, ",")
	//		location = strings.ReplaceAll(location, "0", "未知")
	//		level2RankTableData[i].Location = location
	//	}
	//}

	result.List = level2RankTableData
	result.TotalCount = total
	return result, nil
}

// GetAppIdRankLevel3 获取3级小类排名的趋势图
func (service AppIdService) GetAppIdRankLevel3(param traffic.AppIdReqParam) ([]chart.FlowTrendDot, error) {
	if param.RankLevel != global.LevelThree {
		return nil, errors.New("RankLevel 应当为第3级")
	}
	if param.HostParam == "" && param.DstIpParam == "" {
		return nil, errors.New("hostParam 以及 dstIpParam 不得同时为空")
	}
	ckDb, err := getAppIdDb(param)
	trendSeries, err := GetLevel2TrendData(ckDb)
	if err != nil {
		global.Log.Error("获取应用小类3级趋势图数据错误", zap.Error(err))
	}
	return trendSeries, err
}

// ExportData 导出
func (service AppIdService) ExportData(param traffic.AppIdReqParam) ([]byte, error) {
	fields := []string{"name", "maxDnBps", "maxUpBps", "avgUpBps", "avgDnBps", "upByte", "dnByte", "totalByte", "proportion"}
	headers := []string{"业务小类", "上行峰值(Mbps)", "下行峰值(Mbps)", "上行平均(Mbps)", "下行平均(Mbps)", "上行总量(MB)", "下行总量(MB)", "总流量(MB)", "总量占比"}
	param.Page = 1
	param.Limit = global.CONFIG.ExportLimit
	pageInfo, err := service.GetLevel1TableData(param)
	if err != nil {
		return nil, err
	}
	ckDb, _ := getAppIdDb(param)
	totalByte, err := GetTotalByte(ckDb)
	if err != nil {
		return nil, err
	}
	list := pageInfo.List

	// 将查询结果转换为 []map[string]interface{} 格式
	dataList, ok := list.([]traffic.AppIdLevel1TableData)
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
