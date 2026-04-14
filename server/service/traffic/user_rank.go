package traffic

import (
	"fcas_server/model/traffic"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sort"
)

type UserRankSvc struct {
	Log        *zap.Logger
	Mysql      *gorm.DB
	ClickHouse *gorm.DB
	rankSvc    *userCommonRankSvc
}

func NewUserRankSvc(log *zap.Logger, mysql *gorm.DB, clickhouse *gorm.DB, v1DBName, v2DBName string) *UserRankSvc {
	return &UserRankSvc{
		Log:        log.Named("[UserRank-用户排名]"),
		Mysql:      mysql,
		ClickHouse: clickhouse,

		rankSvc: NewUserCommonRankSvc(log, mysql, clickhouse, v1DBName, v2DBName),
	}
}

// Level1Pie 返回前端需要的Level1Top10饼图数据（使用窗口函数计算出前10个用户的数据）
func (svc UserRankSvc) Level1Pie(params traffic.UserCommonRankParams) ([]*traffic.RankLevel1Base, error) {
	baseSQL := svc.rankSvc.getLevel1BaseSqlByTimeRange(params)

	var input []*traffic.RankLevel1Base
	if err := svc.ClickHouse.Raw(baseSQL).Scan(&input).Error; err != nil {
		svc.Log.Error("查询Level1Pie数据失败", zap.String("sql", baseSQL), zap.Error(err))
		return nil, err
	}

	output := svc.rankSvc.level1AggregateByTrafficTotalTop10(input)

	if err := svc.rankSvc.fillName(output, svc.getNameMap); err != nil {
		svc.Log.Error("填充name失败", zap.Error(err))
	}
	return output, nil
}

// Level2Trend 返回前端需要的Level2趋势图数据
func (svc UserRankSvc) Level2Trend(params traffic.UserCommonRankParams) ([]*traffic.RankLevel1Base, error) {
	sql := svc.rankSvc.getLevel1BaseSqlByTimeRange(params)
	var output []*traffic.RankLevel1Base
	if err := svc.ClickHouse.Raw(sql).Scan(&output).Error; err != nil {
		svc.Log.Error("查询Level2Trend数据失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}

	sort.Slice(output, func(i, j int) bool { return output[i].StartTime < output[j].StartTime })

	if err := svc.rankSvc.fillName(output, svc.getNameMap); err != nil {
		svc.Log.Error("填充name失败", zap.Error(err))
	}
	return output, nil
}

// Level1Table 返回前端需要的Level1表格数据
func (svc UserRankSvc) Level1Table(params traffic.UserCommonRankParams) (rows []*traffic.RankLevel1AggTraffic, gather *traffic.RankLevel1AggTraffic, err error) {
	sql := svc.rankSvc.getLevel1BaseSqlByTimeRange(params)
	var input []*traffic.RankLevel1Base
	if err := svc.ClickHouse.Raw(sql).Scan(&input).Error; err != nil {
		svc.Log.Error("查询Level1Table数据失败", zap.String("sql", sql), zap.Error(err))
		return nil, nil, err
	}

	if err := svc.rankSvc.fillName(input, svc.getNameMap); err != nil {
		svc.Log.Error("填充name失败", zap.Error(err))
	}

	rows = svc.rankSvc.level1AggregateTrafficById(input)
	gather = svc.rankSvc.gather(rows)

	return rows, gather, nil
}

// Level2Table 返回前端需要的Level2表格数据
func (svc UserRankSvc) Level2Table(params traffic.UserCommonRankParams) ([]*traffic.RankLevel2Table, error) {
	output, err := svc.rankSvc.level2Table(params)
	if err != nil {
		svc.Log.Error("查询Level2Table数据失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}
	return output, nil
}

// Level3Table 返回前端需要的Level3表格数据
func (svc UserRankSvc) Level3Table(params traffic.UserCommonRankParams) ([]*traffic.RankLevel3Table, error) {
	output, err := svc.rankSvc.level3Table(params)
	if err != nil {
		svc.Log.Error("查询Level3Table数据失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}
	return output, nil
}

func (svc UserRankSvc) Export(params traffic.UserCommonRankParams) ([]byte, error) {
	return svc.rankSvc.export(params, svc.Level1Table)
}

func (svc UserRankSvc) getNameMap(ids []int64) (map[int64]string, error) {
	type Item struct {
		Id   int64  `gorm:"column:id"`
		Name string `gorm:"column:user_name"`
	}

	var records []*Item
	err := svc.Mysql.Table("dim_user_info").Select("id, user_name").Where("id IN (?)", ids).Scan(&records).Error
	if err != nil {
		return nil, fmt.Errorf("根据用户ID查询用户名失败: %w", err)
	}

	var m = make(map[int64]string, len(records))
	for _, item := range records {
		m[item.Id] = item.Name
	}

	m[0] = "未知"
	m[-1] = "其他"

	return m, nil
}
