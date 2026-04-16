package traffic

import (
	"fcas_server/global"
	"fcas_server/model/common/chart"
	"fcas_server/model/common/request"
	"fcas_server/model/common/response"
	"fcas_server/model/traffic"
	"fcas_server/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func getQuerySpanSeconds(startTime, endTime string) (int, error) {
	particle, err := utils.GetParticleByTimeRange(startTime, endTime)
	if err != nil {
		return 0, err
	}

	start, err := time.Parse(global.DateTimeLayout, startTime)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse(global.DateTimeLayout, endTime)
	if err != nil {
		return 0, err
	}
	if !end.After(start) {
		return 0, fmt.Errorf("end time must be after start time")
	}

	totalSeconds := int(end.Sub(start).Seconds())
	slotCount := totalSeconds / particle
	if totalSeconds%particle != 0 {
		slotCount++
	}
	return slotCount * particle, nil
}

func buildLevel1AverageExpr(totalExpr string, totalSpanSeconds int, alias string) string {
	return fmt.Sprintf("Round((%s) * 8 / %d, 2) AS %s", totalExpr, totalSpanSeconds, alias)
}

func buildGatherDataQuery(ckDb *gorm.DB, totalSpanSeconds int) *gorm.DB {
	aggregateTx := global.V2ClickhouseDB.Table("(?)", ckDb).
		Select(`sum(traffic_up) AS total_up_byte,
				sum(traffic_dn) AS total_dn_byte`)

	// Wrap one more time so ClickHouse does not expand aliases back into aggregate
	// expressions like sum(total_up_byte) when building avg_up_bps/up_byte.
	return global.V2ClickhouseDB.Table("(?)", aggregateTx).
		Select(fmt.Sprintf(`%s,
				%s,
				total_up_byte AS up_byte,
				total_dn_byte AS dn_byte,
				total_up_byte + total_dn_byte AS total_byte`,
			buildLevel1AverageExpr("total_up_byte", totalSpanSeconds, "avg_up_bps"),
			buildLevel1AverageExpr("total_dn_byte", totalSpanSeconds, "avg_dn_bps")))
}

func GetGatherData(ckDb *gorm.DB, level1Data traffic.Level1Data, _ string, totalSpanSeconds int) (traffic.Level1Data, error) {
	var gatherData traffic.GatherData
	err := buildGatherDataQuery(ckDb, totalSpanSeconds).Find(&gatherData).Error
	if err != nil {
		return level1Data, err
	}
	level1Data.GatherData = gatherData
	return level1Data, nil
}

func GetLevel2TableData(ckDb *gorm.DB, param request.PageInfo) (response.PageResult, error) {
	var result response.PageResult
	var trendTableData []chart.FlowTrendDot
	var total int64
	selectSql := "start_time," +
		"sum(traffic_dn_bps) AS dn_bps," +
		"sum(traffic_up_bps) AS up_bps," +
		"(dn_bps + up_bps) AS total_bps"

	err := global.V2ClickhouseDB.Table("(?)", ckDb).Select(selectSql).
		Group("start_time").
		Count(&total).
		Limit(param.Limit).
		Offset(param.Limit * (param.Page - 1)).
		Order("start_time DESC").
		Find(&trendTableData).
		Error
	if err != nil {
		return result, err
	}

	result.TotalCount = total
	result.List = trendTableData
	return result, nil
}

func GetLevel2TrendData(ckDb *gorm.DB) ([]chart.FlowTrendDot, error) {
	var trendData []chart.FlowTrendDot
	selectSql := "start_time, " +
		"sum(traffic_dn_bps) AS dn_bps, " +
		"sum(traffic_up_bps) AS up_bps, " +
		" dn_bps + up_bps AS total_bps"
	err := global.V2ClickhouseDB.Table("(?)", ckDb).Select(selectSql).Group("start_time").Order("start_time").Find(&trendData).Error
	if err != nil {
		return trendData, err
	}
	return trendData, nil
}

// GetPieData 获取Top N饼图
func GetPieData(ckDb *gorm.DB, level1Data traffic.Level1Data, trafficType string) (traffic.Level1Data, error) {
	var pieData []chart.FlowPiePiece
	var selectSql string
	switch trafficType {
	case "isp":
		selectSql = "isp AS name,sum(traffic_dn) AS dn_byte,sum(traffic_dn+traffic_up) AS total_byte"
	case "appType":
		selectSql = "dictGet('app_type_dict','name',app_type) AS name,sum(traffic_dn) AS dn_byte,sum(traffic_dn+traffic_up) AS total_byte"
	case "appId":
		selectSql = "dictGet('app_id_dict','name',app_id) AS name,sum(traffic_dn) AS dn_byte,sum(traffic_dn+traffic_up) AS total_byte"
	}

	err := global.V2ClickhouseDB.Table("(?)", ckDb).
		Select(selectSql).
		Group("name").
		Order("total_byte desc").
		Limit(10).
		Find(&pieData).Error
	if err != nil {
		return level1Data, err
	}
	var otherDnByte = level1Data.GatherData.DnByte
	var otherTotalByte = level1Data.GatherData.TotalByte
	if len(pieData) > 0 {
		for i := 0; i < len(pieData); i++ {
			otherDnByte = otherDnByte - pieData[i].DnByte
			otherTotalByte = otherTotalByte - pieData[i].TotalByte
		}
	}
	pieData = append(pieData, chart.FlowPiePiece{
		Name:      global.OtherPart,
		DnByte:    otherDnByte,
		TotalByte: otherTotalByte,
	})
	level1Data.PieData = pieData
	return level1Data, nil
}

func GetTotalByte(ckDb *gorm.DB) (uint64, error) {
	var totalByte uint64
	err := global.V2ClickhouseDB.Table("(?)", ckDb).Select("sum(traffic_up) + sum(traffic_dn)").Find(&totalByte).Error
	if err != nil {
		global.Log.Error(err.Error())
	}
	return totalByte, err
}
