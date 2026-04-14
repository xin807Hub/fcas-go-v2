package traffic

import (
	"fcas_server/global"
	"fcas_server/model/traffic"
	"fcas_server/utils"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math"
	"sort"
	"strings"
)

var UserNewTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_user_10m",
	global.Interval1hParticle:  "dws_user_1h",
	global.Interval1dParticle:  "dws_user_1d",
}
var UserOldTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_user_10min",
	global.Interval1hParticle:  "dws_user_hour",
	global.Interval1dParticle:  "dws_user_hour",
}
var UserDbMap = map[string]map[int]string{
	global.DatabaseV1: UserOldTableNameMap,
	global.DatabaseV2: UserNewTableNameMap,
}

var SrcNewTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_srcip_top_10m",
	global.Interval1hParticle:  "dws_srcip_top_1h",
	global.Interval1dParticle:  "dws_srcip_top_1d",
}
var SrcOldTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_srcip_top_10min",
	global.Interval1hParticle:  "dws_srcip_top_hour",
	global.Interval1dParticle:  "dws_srcip_top_hour",
}
var SrcDbMap = map[string]map[int]string{
	global.DatabaseV1: SrcOldTableNameMap,
	global.DatabaseV2: SrcNewTableNameMap,
}

var SrcNoParamNewTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_srcip_top_10m_no_params",
	global.Interval1hParticle:  "dws_srcip_top_1h_no_params",
	global.Interval1dParticle:  "dws_srcip_top_1d_no_params",
}
var SrcNoParamOldTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_srcip_top_10min_no_params",
	global.Interval1hParticle:  "dws_srcip_top_hour_no_params",
	global.Interval1dParticle:  "dws_srcip_top_hour_no_params",
}
var SrcNoParamDbMap = map[string]map[int]string{
	global.DatabaseV1: SrcNoParamOldTableNameMap,
	global.DatabaseV2: SrcNoParamNewTableNameMap,
}

var DstNewTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_dstip_top_10m",
	global.Interval1hParticle:  "dws_dstip_top_1h",
	global.Interval1dParticle:  "dws_dstip_top_1d",
}
var DstOldTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_dstip_top_10min",
	global.Interval1hParticle:  "dws_dstip_top_hour",
	global.Interval1dParticle:  "dws_dstip_top_hour",
}
var DstDbMap = map[string]map[int]string{
	global.DatabaseV1: DstOldTableNameMap,
	global.DatabaseV2: DstNewTableNameMap,
}

var DstNoParamNewTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_dstip_top_10m_no_params",
	global.Interval1hParticle:  "dws_dstip_top_1h_no_params",
	global.Interval1dParticle:  "dws_dstip_top_1d_no_params",
}
var DstNoParamOldTableNameMap = map[int]string{
	global.Interval10mParticle: "dws_dstip_top_10min_no_params",
	global.Interval1hParticle:  "dws_dstip_top_hour_no_params",
	global.Interval1dParticle:  "dws_dstip_top_hour_no_params",
}
var DstNoParamDbMap = map[string]map[int]string{
	global.DatabaseV1: DstNoParamOldTableNameMap,
	global.DatabaseV2: DstNoParamNewTableNameMap,
}

type userCommonRankSvc struct {
	Log        *zap.Logger
	Mysql      *gorm.DB
	ClickHouse *gorm.DB

	V1DBName string
	V2DBName string
}

func NewUserCommonRankSvc(log *zap.Logger, mysql *gorm.DB, clickhouse *gorm.DB, v1DBName, v2DBName string) *userCommonRankSvc {
	return &userCommonRankSvc{
		Log:        log,
		Mysql:      mysql,
		ClickHouse: clickhouse,
		V1DBName:   v1DBName,
		V2DBName:   v2DBName,
	}
}

// level2Table 返回前端需要的Level2表格数据
func (svc userCommonRankSvc) level2Table(params traffic.UserCommonRankParams) ([]*traffic.RankLevel2Table, error) {
	sql := svc.getLevel2TableSqlByTimeRange(params)

	var output []*traffic.RankLevel2Table
	if err := svc.ClickHouse.Raw(sql).Scan(&output).Error; err != nil {
		return nil, err
	}

	// 计算总量占比
	svc.calcTotalProportion(output)

	return output, nil
}

// level3Table 返回前端需要的Level3表格数据
func (svc userCommonRankSvc) level3Table(params traffic.UserCommonRankParams) ([]*traffic.RankLevel3Table, error) {
	sql := svc.getLevel3TableSqlByTimeRange(params)

	var output []*traffic.RankLevel3Table
	if err := svc.ClickHouse.Raw(sql).Scan(&output).Error; err != nil {
		return nil, err
	}

	// 计算总流速
	svc.calcTotalSpeed(output)
	return output, nil
}

// level1AggregateTrafficById 根据ID聚合流量
func (svc userCommonRankSvc) level1AggregateTrafficById(input []*traffic.RankLevel1Base) []*traffic.RankLevel1AggTraffic {
	aggregation := make(map[int64]*traffic.RankLevel1AggTraffic)

	// 聚合数据
	for _, item := range input {

		if _, ok := aggregation[item.ID]; !ok {
			aggregation[item.ID] = &traffic.RankLevel1AggTraffic{
				ID:   item.ID,
				Name: item.Name,
			}
		}

		aggregation[item.ID].TrafficUp += item.TrafficUp
		aggregation[item.ID].TrafficDn += item.TrafficDn
		aggregation[item.ID].TrafficTotal += item.TrafficTotal
		aggregation[item.ID].MaxSpeedUp = math.Max(aggregation[item.ID].MaxSpeedUp, item.SpeedUp)
		aggregation[item.ID].MaxSpeedDn = math.Max(aggregation[item.ID].MaxSpeedDn, item.SpeedDn)
		aggregation[item.ID].AvgSpeedUp += item.SpeedUp
		aggregation[item.ID].AvgSpeedDn += item.SpeedDn
		aggregation[item.ID].Count += 1
	}

	// 计算平均值并转化为切片
	output := make([]*traffic.RankLevel1AggTraffic, 0, len(aggregation))
	for _, agg := range aggregation {
		agg.AvgSpeedUp /= float64(agg.Count)
		agg.AvgSpeedDn /= float64(agg.Count)
		output = append(output, agg)
	}

	// 按 traffic_total 降序排序
	sort.Slice(output, func(i, j int) bool { return output[i].TrafficTotal > output[j].TrafficTotal })

	return output
}

// level1AggregateByTrafficTotalTop10 根据 traffic_total 排序并取前10，其他归为一类
func (svc userCommonRankSvc) level1AggregateByTrafficTotalTop10(input []*traffic.RankLevel1Base) []*traffic.RankLevel1Base {
	// 聚合相同ID的数据
	aggregateMap := make(map[int64]*traffic.RankLevel1Base)
	for _, item := range input {
		if _, ok := aggregateMap[item.ID]; !ok {
			aggregateMap[item.ID] = &traffic.RankLevel1Base{ID: item.ID, Name: item.Name}
		}
		aggregateMap[item.ID].TrafficUp += item.TrafficUp
		aggregateMap[item.ID].TrafficDn += item.TrafficDn
		aggregateMap[item.ID].TrafficTotal += item.TrafficTotal
	}
	aggregated := lo.Values(aggregateMap)

	// 按 traffic_total 降序排序
	sort.Slice(aggregated, func(i, j int) bool { return aggregated[i].TrafficTotal > aggregated[j].TrafficTotal })
	if len(aggregated) <= 10 {
		return aggregated
	}

	// 取前10
	top10 := aggregated[:10]
	others := &traffic.RankLevel1Base{ID: -1, Name: "其他"}
	for _, item := range aggregated[10:] {
		others.TrafficUp += item.TrafficUp
		others.TrafficDn += item.TrafficDn
		others.TrafficTotal += item.TrafficTotal
	}

	return append(top10, others)
}

// getLevel1BaseSqlByTimeRange 根据时间范围构造Level1基础数据SQL
func (svc userCommonRankSvc) getLevel1BaseSqlByTimeRange(params traffic.UserCommonRankParams) string {
	var baseSQL string

	switch params.TimeRangeType {
	case "v1":
		baseSQL = svc.level1BaseSQL(params, svc.V1DBName)

	case "v2":
		baseSQL = svc.level1BaseSQL(params, svc.V2DBName)

	case "v1v2":
		v1BaseSQL := svc.level1BaseSQL(params, svc.V1DBName)
		v2BaseSQL := svc.level1BaseSQL(params, svc.V2DBName)

		baseSQL = fmt.Sprintf(`(%s) UNION ALL (%s)`, v1BaseSQL, v2BaseSQL)
	}

	sql := fmt.Sprintf(`
			SELECT user_id,
				   start_time,
				   sum(traffic_up)                       	   AS traffic_up,
				   sum(traffic_dn)                             AS traffic_dn,
				   plus(traffic_up, traffic_dn) 	   		   AS traffic_total,
				   divide(multiply(traffic_up, 8), {particle}) AS speed_up,
				   divide(multiply(traffic_dn, 8), {particle}) AS speed_dn
			FROM (%s)
			GROUP BY user_id, start_time
			`, baseSQL)

	sql = strings.ReplaceAll(sql, "{particle}", fmt.Sprint(params.Particle))

	return sql
}

// getLevel2TableSqlByTimeRange 根据时间范围构造Level2表格SQL
func (svc userCommonRankSvc) getLevel2TableSqlByTimeRange(params traffic.UserCommonRankParams) string {
	var baseSQL string

	switch params.TimeRangeType {
	case "v1":
		baseSQL = svc.level2TableBaseSQL(params, svc.V1DBName)

	case "v2":
		baseSQL = svc.level2TableBaseSQL(params, svc.V2DBName)

	case "v1v2":
		v1BaseSQL := svc.level2TableBaseSQL(params, svc.V1DBName)
		v2BaseSQL := svc.level2TableBaseSQL(params, svc.V2DBName)

		baseSQL = fmt.Sprintf(`(%s) UNION ALL (%s)`, v1BaseSQL, v2BaseSQL)
	}

	sql := fmt.Sprintf(`
			 SELECT src_ip,
					sum(traffic_up)    AS traffic_up,
					sum(traffic_dn)    AS traffic_dn,
					sum(traffic_total) AS traffic_total,
					avg(speed_up)      AS avg_speed_up,
					avg(speed_dn)      AS avg_speed_dn
			 FROM (%s)
			 GROUP BY src_ip
			 ORDER BY traffic_total DESC
			 LIMIT %d
			 `, baseSQL, params.TopN)

	/*
		SELECT src_ip,
		       sum(traffic_dn)    AS traffic_dn,
		       sum(traffic_up)    AS traffic_up,
		       sum(traffic_total) AS traffic_total
		FROM (
		         SELECT src_ip                       AS src_ip,
		                sumMerge(bytes_up_view)      AS traffic_up,
		                sumMerge(bytes_dn_view)      AS traffic_dn,
		                plus(traffic_up, traffic_dn) AS traffic_total
		         FROM dws_srcip_top_hour
		         -- WHERE 条件
		         WHERE start_time >= '2024-10-31 00:00:00'
		           AND start_time < '2024-10-31 01:00:00'

		         GROUP BY src_ip
		         ORDER BY traffic_total DESC
		         LIMIT 20

		         UNION
		         DISTINCT

		         SELECT dst_ip                       AS src_ip,
		                sumMerge(bytes_up_view)      AS traffic_up,
		                sumMerge(bytes_dn_view)      AS traffic_dn,
		                plus(traffic_up, traffic_dn) AS traffic_total
		         FROM dws_dstip_top_hour

		         -- WHERE 条件
		         WHERE start_time >= '2024-10-31 00:00:00'
		           AND start_time < '2024-10-31 01:00:00'

		         GROUP BY dst_ip
		         ORDER BY traffic_total DESC
		         LIMIT 20
		         )
		GROUP BY src_ip
		ORDER BY traffic_total DESC
		LIMIT 20
	*/

	return sql
}

// getLevel3TableSqlByTimeRange 根据时间范围构造Level3表格SQL
func (svc userCommonRankSvc) getLevel3TableSqlByTimeRange(params traffic.UserCommonRankParams) string {
	var baseSQL string

	switch params.TimeRangeType {
	case "v1":
		baseSQL = svc.level3TableBaseSQL(params, svc.V1DBName)

	case "v2":
		baseSQL = svc.level3TableBaseSQL(params, svc.V2DBName)

	case "v1v2":
		v1BaseSQL := svc.level3TableBaseSQL(params, svc.V1DBName)
		v2BaseSQL := svc.level3TableBaseSQL(params, svc.V2DBName)

		baseSQL = fmt.Sprintf(`(%s) UNION ALL (%s)`, v1BaseSQL, v2BaseSQL)
	}

	sql := fmt.Sprintf(`
			SELECT start_time,
				   src_ip,
				   divide(multiply(sum(traffic_up), 8), {particle}) AS speed_up,
				   divide(multiply(sum(traffic_dn), 8), {particle}) AS speed_dn
			FROM (%s)
			GROUP BY start_time, src_ip
			ORDER BY start_time
			`, baseSQL)

	sql = strings.ReplaceAll(sql, "{particle}", fmt.Sprint(params.Particle)) // 替换计算速率的颗粒度

	/*
		SELECT start_time,
		       src_ip,
		       divide(multiply(sum(traffic_up), 8), 600) AS speed_up,
		       divide(multiply(sum(traffic_dn), 8), 600) AS speed_dn
		FROM (
		         SELECT start_time,
		                src_ip,
		                sumMerge(bytes_dn_view) AS traffic_dn,
		                sumMerge(bytes_up_view) AS traffic_up
		         FROM bigdata_fcas.dws_srcip_top_10min
		         -- WHERE 条件
		         WHERE (("start_time" >= '2024-10-31 00:00:00') AND ("start_time" < '2024-10-31 01:00:00'))
		         GROUP BY start_time, src_ip

		         UNION
		         DISTINCT

		         SELECT start_time,
		                dst_ip                  AS src_ip,
		                sumMerge(bytes_dn_view) AS traffic_dn,
		                sumMerge(bytes_up_view) AS traffic_up
		         FROM bigdata_fcas.dws_dstip_top_10min
		         -- WHERE 条件
		         WHERE (("start_time" >= '2024-10-31 00:00:00') AND ("start_time" < '2024-10-31 01:00:00'))
		         GROUP BY start_time, dst_ip
		         )
		GROUP BY start_time, src_ip
		ORDER BY start_time
	*/

	return sql
}

// level1BaseSQL 构造Level1基础数据SQL
func (svc userCommonRankSvc) level1BaseSQL(params traffic.UserCommonRankParams, dbName string) string {
	sql := `
			SELECT user_id,
				   start_time,
				   sum(traffic_up)                       		AS traffic_up,
				   sum(traffic_dn)                       		AS traffic_dn
			FROM (
					 SELECT user_id,
							start_time,
							sumMerge(bytes_up_view) as traffic_up,
							sumMerge(bytes_dn_view) as traffic_dn
					 FROM {table}
					 %s
					 GROUP BY user_id, start_time
			
					 UNION
					 DISTINCT
			
					 SELECT d_user_id               AS user_id,
							start_time,
							sumMerge(bytes_up_view) as traffic_up,
							sumMerge(bytes_dn_view) as traffic_dn
					 FROM {table}
					 %s
					 GROUP BY d_user_id, start_time
					 )
			GROUP BY user_id, start_time
			`

	table := UserDbMap[dbName][params.Particle]

	sql = strings.ReplaceAll(sql, "{table}", fmt.Sprintf("%s.%s", dbName, table))

	// 构造WHERE条件
	builder := goqu.From().Select()

	if params.StartTime != "" && params.EndTime != "" {
		builder = builder.Where(goqu.And(
			goqu.C("start_time").Gte(params.StartTime),
			goqu.C("start_time").Lt(params.EndTime),
		))
	}

	// 通用Where
	if len(params.LinkIdList) != 0 {
		builder = builder.Where(goqu.C("link_id").In(params.LinkIdList))
	}

	// 通用Where
	if len(params.AppIdList) != 0 {
		builder = builder.Where(goqu.C("app_id").In(params.AppIdList))
	}

	// 剩下条件需要区分原和目的
	srcB := builder
	dstB := builder
	if len(params.UserIdList) != 0 {
		srcB = srcB.Where(goqu.C("user_id").In(params.UserIdList))
		dstB = dstB.Where(goqu.C("d_user_id").In(params.UserIdList))
	}

	// 运营商
	if len(params.IspNameList) != 0 {
		srcB = srcB.Where(goqu.C("isp").In(params.IspNameList))
		dstB = dstB.Where(goqu.C("d_isp").In(params.IspNameList))
	}

	srcWhereSQL, _, _ := srcB.ToSQL()
	dstWhereSQL, _, _ := dstB.ToSQL()
	srcWhereSQL = strings.TrimSpace(srcWhereSQL[len("SELECT *"):])
	dstWhereSQL = strings.TrimSpace(dstWhereSQL[len("SELECT *"):])

	sql = fmt.Sprintf(sql, srcWhereSQL, dstWhereSQL)

	return sql
}

// level1BaseSQL 构造Level2表格基础SQL
func (svc userCommonRankSvc) level2TableBaseSQL(params traffic.UserCommonRankParams, dbName string) string {
	sql := `
		  SELECT src_ip                                		 AS src_ip,
				 sumMerge(bytes_up_view)               	 	 AS traffic_up,
				 sumMerge(bytes_dn_view)               		 AS traffic_dn,
				 plus(traffic_up, traffic_dn)          		 AS traffic_total,
				 divide(multiply(traffic_up, 8), {particle}) AS speed_up,
				 divide(multiply(traffic_dn, 8), {particle}) AS speed_dn
		  FROM {srcTable}
		  -- WHERE 条件
		  %s

		  GROUP BY src_ip, start_time
		  ORDER BY traffic_total DESC
		  LIMIT {topN}

		  UNION
		  DISTINCT

		  SELECT dst_ip                                		 AS src_ip,
				 sumMerge(bytes_up_view)               		 AS traffic_up,
				 sumMerge(bytes_dn_view)               		 AS traffic_dn,
				 plus(traffic_up, traffic_dn)          		 AS traffic_total,
				 divide(multiply(traffic_up, 8), {particle}) AS speed_up,
				 divide(multiply(traffic_dn, 8), {particle}) AS speed_dn
		  FROM {dstTable}
		  -- WHERE 条件
		  %s

		  GROUP BY dst_ip, start_time
		  ORDER BY traffic_total DESC
		  LIMIT {topN}
		`

	sql = strings.ReplaceAll(sql, "{particle}", fmt.Sprint(params.Particle)) // 替换计算速率的颗粒度

	// 根据颗粒度选择表
	var srcTable, dstTable string
	srcTable = SrcNoParamDbMap[dbName][params.Particle]
	dstTable = DstNoParamDbMap[dbName][params.Particle]

	sql = strings.ReplaceAll(sql, "{srcTable}", fmt.Sprintf("%s.%s", dbName, srcTable)) // 替换源IP流量表名
	sql = strings.ReplaceAll(sql, "{dstTable}", fmt.Sprintf("%s.%s", dbName, dstTable)) // 替换目的IP流量表名
	sql = strings.ReplaceAll(sql, "{topN}", fmt.Sprint(params.TopN))                    // 替换Limit数量

	// 构造WHERE条件
	whereBuilder := goqu.From().Select()

	if params.StartTime != "" && params.EndTime != "" {
		whereBuilder = whereBuilder.Where(goqu.And(
			goqu.C("start_time").Gte(params.StartTime),
			goqu.C("start_time").Lt(params.EndTime),
		))
	}

	// link_id,通用Where
	if len(params.LinkIdList) != 0 {
		whereBuilder = whereBuilder.Where(goqu.C("link_id").In(params.UserIdList))
	}

	// app_id,通用Where
	if len(params.AppIdList) != 0 {
		whereBuilder = whereBuilder.Where(goqu.C("app_id").In(params.AppIdList))
	}

	// 剩下条件需要区分原和目的
	srcWhereBuilder := whereBuilder
	dstWhereBuilder := whereBuilder
	if len(params.UserIdList) != 0 {
		srcWhereBuilder = srcWhereBuilder.Where(goqu.C("user_id").In(params.UserIdList))
		dstWhereBuilder = dstWhereBuilder.Where(goqu.C("d_user_id").In(params.UserIdList))
	}

	// 运营商
	if len(params.IspNameList) != 0 {
		srcWhereBuilder = srcWhereBuilder.Where(goqu.C("d_isp").In(params.IspNameList))
		dstWhereBuilder = dstWhereBuilder.Where(goqu.C("isp").In(params.IspNameList))
	}

	if params.SrcIp != "" {
		srcWhereBuilder = srcWhereBuilder.Where(goqu.C("src_ip").Eq(params.SrcIp))
		dstWhereBuilder = dstWhereBuilder.Where(goqu.C("dst_ip").Eq(params.SrcIp))
	}

	srcWhereSQL, _, _ := srcWhereBuilder.ToSQL()
	dstWhereSQL, _, _ := dstWhereBuilder.ToSQL()
	srcWhereSQL = strings.TrimSpace(srcWhereSQL[len("SELECT *"):])
	dstWhereSQL = strings.TrimSpace(dstWhereSQL[len("SELECT *"):])
	sql = fmt.Sprintf(sql, srcWhereSQL, dstWhereSQL)

	return sql
}

// level3TableBaseSQL 构造Level3表格基础SQL
func (svc userCommonRankSvc) level3TableBaseSQL(params traffic.UserCommonRankParams, dbName string) string {
	sql := `
			SELECT start_time,
				   src_ip,
				   sumMerge(bytes_dn_view) AS traffic_dn,
				   sumMerge(bytes_up_view) AS traffic_up
			FROM {srcTable}
			-- WHERE 条件
			%s
			GROUP BY start_time, src_ip
			
			UNION 
			DISTINCT
			
			SELECT start_time,
				   dst_ip 				   AS src_ip,
				   sumMerge(bytes_dn_view) AS traffic_dn,
				   sumMerge(bytes_up_view) AS traffic_up
			FROM {dstTable}
			-- WHERE 条件
			%s
			GROUP BY start_time, dst_ip
			`

	var srcTable, dstTable string
	srcTable = SrcDbMap[dbName][params.Particle]
	dstTable = DstDbMap[dbName][params.Particle]

	sql = strings.ReplaceAll(sql, "{srcTable}", fmt.Sprintf("%s.%s", dbName, srcTable))
	sql = strings.ReplaceAll(sql, "{dstTable}", fmt.Sprintf("%s.%s", dbName, dstTable))

	// 构造WHERE条件
	whereBuilder := goqu.From().Select()
	if params.StartTime != "" && params.EndTime != "" {
		whereBuilder = whereBuilder.Where(goqu.And(
			goqu.C("start_time").Gte(params.StartTime),
			goqu.C("start_time").Lt(params.EndTime),
		))
	}

	if len(params.LinkIdList) != 0 {
		whereBuilder = whereBuilder.Where(goqu.C("link_id").In(params.LinkIdList))
	}

	if len(params.AppIdList) != 0 {
		whereBuilder = whereBuilder.Where(goqu.C("app_id").In(params.AppIdList))
	}

	// 剩下条件需要区分原和目的
	srcWhereBuilder := whereBuilder
	dstWhereBuilder := whereBuilder

	if params.SrcIp != "" {
		srcWhereBuilder = srcWhereBuilder.Where(goqu.C("src_ip").Eq(params.SrcIp))
		dstWhereBuilder = dstWhereBuilder.Where(goqu.C("dst_ip").Eq(params.SrcIp))
	}

	// 运营商
	if len(params.IspNameList) != 0 {
		srcWhereBuilder = srcWhereBuilder.Where(goqu.C("d_isp").In(params.IspNameList))
		dstWhereBuilder = dstWhereBuilder.Where(goqu.C("isp").In(params.IspNameList))
	}

	srcWhereSQL, _, _ := srcWhereBuilder.ToSQL()
	dstWhereSQL, _, _ := dstWhereBuilder.ToSQL()
	srcWhereSQL = strings.TrimSpace(srcWhereSQL[len("SELECT *"):])
	dstWhereSQL = strings.TrimSpace(dstWhereSQL[len("SELECT *"):])
	sql = fmt.Sprintf(sql, srcWhereSQL, dstWhereSQL)

	return sql
}

// aggregate 聚合数据，根据id和时间聚合，onlyById 表示是否只聚合id，否则根据id和start_time聚合，用于crowd和group
func (svc userCommonRankSvc) aggregate(input []*traffic.RankLevel1Base, relations map[int64][]int64, onlyById bool) []*traffic.RankLevel1Base {
	// 定义一个 map 来存储聚合后的结果
	aggMap := make(map[string]*traffic.RankLevel1Base)

	// 遍历用户进行聚合计算
	for _, item := range input {
		// 获取当前用户所属的组or群列表，
		// 若该用户在 userRelations 中有对应的组or群，则进入 if 语句块。
		if relationIds, exists := relations[item.ID]; exists {
			// 遍历当前用户所属的关系中的id，因为同一个用户可能属于多个组or群，需要在每个组or群上累加该用户的流量。
			for _, id := range relationIds {
				// 检查 aggregatedMap 中是否已经存在该组or群的数据，根据id和时间作为key
				var key string
				if onlyById {
					key = fmt.Sprintf("%d", id)
				} else {
					key = fmt.Sprintf("%d_%s", id, item.StartTime)
				}
				if _, ok := aggMap[key]; !ok {
					// 若没有，则初始化数据
					aggMap[key] = &traffic.RankLevel1Base{ID: id, StartTime: item.StartTime}
				}
				// 聚合流量数据
				aggMap[key].TrafficUp += item.TrafficUp
				aggMap[key].TrafficDn += item.TrafficDn
				aggMap[key].TrafficTotal += item.TrafficTotal
				aggMap[key].SpeedUp += item.SpeedUp
				aggMap[key].SpeedDn += item.SpeedDn
			}
		}
	}

	output := lo.Values(aggMap)
	sort.Slice(output, func(i, j int) bool {
		if output[i].ID == output[j].ID {
			return output[i].StartTime < output[j].StartTime
		}
		return output[i].ID < output[j].ID
	})

	return output
}

// fillName 填充名称
func (svc userCommonRankSvc) fillName(input []*traffic.RankLevel1Base, nameMapFunc func([]int64) (map[int64]string, error)) error {
	// 获取id列表
	ids := lo.Map(input, func(item *traffic.RankLevel1Base, _ int) int64 { return item.ID })
	nameMap, err := nameMapFunc(ids)
	if err != nil {
		return err
	}

	// 填充Name
	lo.ForEach(input, func(item *traffic.RankLevel1Base, _ int) { item.Name = nameMap[item.ID] })

	return nil
}

// export 导出
func (svc userCommonRankSvc) export(params traffic.UserCommonRankParams, fn func(params traffic.UserCommonRankParams) ([]*traffic.RankLevel1AggTraffic, *traffic.RankLevel1AggTraffic, error)) ([]byte, error) {
	records, _, err := fn(params)
	if err != nil {
		return nil, err
	}

	headers := []string{"用户", "上行峰值(Mbps)", "下行峰值(Mbps)", "上行平均(Mbps)", "下行平均(Mbps)", "上行总量(MB)", "下行总量(MB)", "总流量(MB)", "总量占比"}
	fields := []string{"name", "maxSpeedUp", "maxSpeedDn", "avgSpeedUp", "avgSpeedDn", "trafficUp", "trafficDn", "trafficTotal", "ratio"}

	totalAll := lo.SumBy(records, func(item *traffic.RankLevel1AggTraffic) uint64 { return item.TrafficTotal })
	totalAllFloat := float64(totalAll) / 1000.0 / 1000.0

	output := make([]map[string]any, 0, len(records))
	for _, r := range records {
		m := make(map[string]any, 9)
		m[fields[0]] = r.Name
		m[fields[1]] = fmt.Sprintf("%.2f", (r.MaxSpeedUp)/1000.0/1000.0)
		m[fields[2]] = fmt.Sprintf("%.2f", (r.MaxSpeedDn)/1000.0/1000.0)
		m[fields[3]] = fmt.Sprintf("%.2f", (r.AvgSpeedUp)/1000.0/1000.0)
		m[fields[4]] = fmt.Sprintf("%.2f", (r.AvgSpeedDn)/1000.0/1000.0)
		m[fields[5]] = fmt.Sprintf("%.2f", (float64(r.TrafficUp))/1000.0/1000.0)
		m[fields[6]] = fmt.Sprintf("%.2f", (float64(r.TrafficDn))/1000.0/1000.0)

		trafficTotal := float64(r.TrafficTotal) / 1000.0 / 1000.0
		m[fields[7]] = fmt.Sprintf("%.2f", trafficTotal)
		m[fields[8]] = fmt.Sprintf("%.2f", (trafficTotal/totalAllFloat)*100.0)

		output = append(output, m)
	}

	return utils.ExportToExcel(fields, headers, output)
}

// gather 汇总
func (svc userCommonRankSvc) gather(input []*traffic.RankLevel1AggTraffic) *traffic.RankLevel1AggTraffic {
	var gather traffic.RankLevel1AggTraffic
	for _, item := range input {
		gather.TrafficUp += item.TrafficUp
		gather.TrafficDn += item.TrafficDn
		gather.TrafficTotal += item.TrafficTotal

		gather.AvgSpeedUp += item.AvgSpeedUp
		gather.AvgSpeedDn += item.AvgSpeedDn
	}

	return &gather
}

// calcTotalProportion 计算并填充总量占比
func (svc userCommonRankSvc) calcTotalProportion(input []*traffic.RankLevel2Table) {
	total := lo.SumBy(input, func(item *traffic.RankLevel2Table) float64 {
		return float64(item.TrafficTotal)
	})

	if total == 0 {
		return
	}

	for _, item := range input {
		item.TotalProportion = float64(item.TrafficTotal) / total
	}
}

// calcTotalSpeed 计算并填充总流速
func (svc userCommonRankSvc) calcTotalSpeed(input []*traffic.RankLevel3Table) {
	for _, item := range input {
		item.TotalSpeed = item.SpeedDn + item.SpeedUp
	}
}
