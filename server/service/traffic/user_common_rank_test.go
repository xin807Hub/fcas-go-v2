package traffic

import (
	"fcas_server/model/traffic"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

const host = "192.168.4.146"

func setupLog() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func setupMysql() *gorm.DB {
	client, err := gorm.Open(mysql.Open(
		fmt.Sprintf("root:123456@tcp(%s:3306)/fcas_service?charset=utf8mb3&parseTime=True&loc=Local", host),
	))
	if err != nil {
		log.Fatalln("连接数据库失败, error=" + err.Error())
	}
	return client
}

func setupClickHouse() *gorm.DB {
	client, err := gorm.Open(clickhouse.Open(
		fmt.Sprintf("clickhouse://default:123456@%s:9000/bigdata_fcas", host),
	))
	if err != nil {
		log.Fatalln("连接数据库失败, error=" + err.Error())
	}
	return client
}

func setupUserCommonRankSvc() *userCommonRankSvc {
	return NewUserCommonRankSvc(
		setupLog(),
		setupMysql(),
		setupClickHouse(),
		"bigdata_fcas",
		"bigdata_fcas_v2",
	)
}

func TestUserCommonRankSvc_level2Table(t *testing.T) {
	svc := setupUserCommonRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v2",
		StartTime:     "2024-11-27 00:00:00",
		EndTime:       "2024-11-29 00:00:00",
		SrcIp:         "194.41.59.218",
		UserIdList:    []int64{17},
		TopN:          15,
	}

	output, err := svc.level2Table(params)
	if err != nil {
		return
	}
	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestUserCommonRankSvc_level3Table(t *testing.T) {
	svc := setupUserCommonRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v2",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-20 01:00:00",
		SrcIp:         "0A05:A8C0:0000:0000:0000:0000:0000:0000",
	}

	output, err := svc.level3Table(params)
	if err != nil {
		return
	}
	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestUserCommonRankSvc_level1AggTraffic(t *testing.T) {
	svc := setupUserCommonRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v2",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-20 01:00:00",
	}

	sql := svc.getLevel1BaseSqlByTimeRange(params)

	var input []*traffic.RankLevel1Base
	if err := svc.ClickHouse.Raw(sql).Scan(&input).Error; err != nil {
		return
	}

	output := svc.level1AggregateTrafficById(input)
	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestUserCommonRankSvc_level1AggTop10(t *testing.T) {
	svc := setupUserCommonRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v2",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-20 01:00:00",
	}

	sql := svc.getLevel1BaseSqlByTimeRange(params)

	var input []*traffic.RankLevel1Base
	if err := svc.ClickHouse.Raw(sql).Scan(&input).Error; err != nil {
		return
	}

	output := svc.level1AggregateByTrafficTotalTop10(input)
	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}

}

func TestUserCommonRankSvc_getLevel1BaseSQLByTimeRange(t *testing.T) {
	svc := setupUserCommonRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-10-31 01:00:00",
		UserIdList:    []int64{1, 3, 5},
		CrowdIdList:   []int64{1, 3, 5},
		GroupIdList:   []int64{1, 3, 5},
		LinkIdList:    []int64{1, 3, 5},
		AppIdList:     []int64{1, 3, 5},
		IspNameList:   []string{"电信", "联通", "移动"},
	}

	sql := svc.getLevel1BaseSqlByTimeRange(params)
	fmt.Println(sql)
}

func TestUserCommonRankSvc_getLevel2TableSqlByTimeRange(t *testing.T) {
	svc := setupUserCommonRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-10-31 01:00:00",
		TopN:          15,
	}

	sql := svc.getLevel2TableSqlByTimeRange(params)
	fmt.Println(sql)
}

func TestUserCommonRankSvc_getLevel3TableSqlByTimeRange(t *testing.T) {
	svc := setupUserCommonRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-10-31 01:00:00",
	}

	sql := svc.getLevel3TableSqlByTimeRange(params)
	fmt.Println(sql)
}

func TestUserCommonRankSvc_aggregate(t *testing.T) {

	baseData := []*traffic.RankLevel1Base{
		{ID: 1, StartTime: "2024-10-31 00:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 1, StartTime: "2024-10-31 01:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 1, StartTime: "2024-10-31 02:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},

		{ID: 2, StartTime: "2024-10-31 00:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 2, StartTime: "2024-10-31 01:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 2, StartTime: "2024-10-31 02:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},

		{ID: 3, StartTime: "2024-10-31 00:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 3, StartTime: "2024-10-31 01:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 3, StartTime: "2024-10-31 02:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},

		{ID: 4, StartTime: "2024-10-31 00:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 4, StartTime: "2024-10-31 01:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
		{ID: 4, StartTime: "2024-10-31 02:00:00", TrafficUp: 2, TrafficDn: 2, TrafficTotal: 4, SpeedUp: 0.1, SpeedDn: 0.1},
	}

	relations := map[int64][]int64{
		1: {1, 2},
		2: {1, 2},
		3: {2},
		4: {1, 2, 3},
	}

	svc := setupUserCommonRankSvc()

	output := svc.aggregate(baseData, relations, false)
	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}
