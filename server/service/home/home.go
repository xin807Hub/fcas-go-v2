package home

import (
	"errors"
	"fcas_server/global"
	"fcas_server/model/common/chart"
	"fcas_server/model/home"
	"fcas_server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct{}

func getDbV1(param home.ReqParam) *gorm.DB {
	ckDb := global.V1ClickhouseDB.
		Table("bigdata_fcas.dws_link_10min").
		Select("start_time,link_id,app_traffic_up_speed*1000*1000 AS up_bps, app_traffic_dn_speed*1000*1000 AS dn_bps").
		Where("start_time >= ? AND start_time < ? ", param.StartTime, param.EndTime)
	if len(param.LinkIdList) > 0 {
		ckDb = ckDb.Where("link_id IN (?)", param.LinkIdList)
	}
	return ckDb
}

func getDbV2(param home.ReqParam) *gorm.DB {
	ckDb := global.V2ClickhouseDB.
		Table("bigdata_fcas_v2.dws_link_10m_view").
		Select("start_time, link_id, up_bps, dn_bps").
		Where("start_time >= ? AND start_time < ? ", param.StartTime, param.EndTime)
	if len(param.LinkIdList) > 0 {
		ckDb = ckDb.Where("link_id IN (?)", param.LinkIdList)
	}
	return ckDb
}

func (homeService Service) GetHomeData(param home.ReqParam) (homeData home.RespHomeData, err error) {
	queryTimeRangeType := utils.GetDbTypeByTimeRange(param.StartTime, param.EndTime)
	var trendSeries []chart.FlowTrendDot
	var linkTables []home.LinkTableData
	var ckDb1 *gorm.DB
	var ckDb2 *gorm.DB

	switch queryTimeRangeType {
	case global.QueryNew:
		ckDb1 = getDbV2(param)
		ckDb2 = getDbV2(param)
	case global.QueryCrossOld2New:
		ckDbNew := getDbV2(param)
		ckDbOld := getDbV1(param)
		ckDb1 = global.V2ClickhouseDB.Table("(? UNION ALL ?)", ckDbNew, ckDbOld)
		ckDb2 = global.V2ClickhouseDB.Table("(? UNION ALL ?)", ckDbNew, ckDbOld)
	case global.QueryOld:
		ckDb1 = global.V2ClickhouseDB.Table("(?)", getDbV1(param))
		ckDb2 = global.V2ClickhouseDB.Table("(?)", getDbV1(param))
	default:
		return home.RespHomeData{}, errors.New("开始时间、结束时间解析错误")
	}

	err = ckDb1.
		Select("start_time,sum(up_bps) AS up_bps,sum(dn_bps) AS dn_bps").
		Group("start_time").
		Order("start_time").
		Find(&trendSeries).Error
	if err != nil {
		global.Log.Error("查询链路趋势图数据错误", zap.Error(err))
		return homeData, err
	}
	homeData.Series = trendSeries

	err = ckDb2.
		Select("dictGet('link_dict','name',link_id) AS link_name,max(up_bps) AS max_up_bps, max(dn_bps) AS max_dn_bps,avg(up_bps) AS avg_up_bps, avg(dn_bps) AS avg_dn_bps").
		Group("link_id").
		Find(&linkTables).Error
	if err != nil {
		global.Log.Error("查询链路表格数据错误", zap.Error(err))
		return homeData, err
	}
	homeData.TableData = linkTables
	return homeData, nil
}
