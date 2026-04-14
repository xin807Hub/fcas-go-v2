package traffic

import (
	"fcas_server/model/traffic"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserCrowdRankSvc struct {
	Log        *zap.Logger
	Mysql      *gorm.DB
	ClickHouse *gorm.DB

	rankSvc *userCommonRankSvc
}

func NewUserCrowdRankSvc(log *zap.Logger, mysql *gorm.DB, clickhouse *gorm.DB, v1DBName, v2DBName string) *UserCrowdRankSvc {
	return &UserCrowdRankSvc{
		Log:        log.Named("[UserCrowdRank-用户组排名]"),
		Mysql:      mysql,
		ClickHouse: clickhouse,

		rankSvc: NewUserCommonRankSvc(log, mysql, clickhouse, v1DBName, v2DBName),
	}
}

// Level1Pie 返回前端需要的Level1Pie饼图数据
func (svc UserCrowdRankSvc) Level1Pie(params traffic.UserCommonRankParams) ([]*traffic.RankLevel1Base, error) {
	baseData, err := svc.aggregatedLevel1Base(params, true)
	if err != nil {
		svc.Log.Error("查询Level1Pie数据失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}

	output := svc.rankSvc.level1AggregateByTrafficTotalTop10(baseData)

	return output, nil
}

// Level2Trend 返回前端需要的Level2趋势图数据
func (svc UserCrowdRankSvc) Level2Trend(params traffic.UserCommonRankParams) ([]*traffic.RankLevel1Base, error) {
	baseData, err := svc.aggregatedLevel1Base(params, false)
	if err != nil {
		svc.Log.Error("查询Level2Trend数据失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}
	return baseData, nil
}

// Level1Table 返回前端需要的Level1表格数据
func (svc UserCrowdRankSvc) Level1Table(params traffic.UserCommonRankParams) ([]*traffic.RankLevel1AggTraffic, *traffic.RankLevel1AggTraffic, error) {
	baseData, err := svc.aggregatedLevel1Base(params, false)
	if err != nil {
		svc.Log.Error("查询Level1Table数据失败", zap.Any("params", params), zap.Error(err))
		return nil, nil, err
	}

	rows := svc.rankSvc.level1AggregateTrafficById(baseData)
	gather := svc.rankSvc.gather(rows)

	return rows, gather, nil
}

// Level2Table 返回前端需要的Level2表格数据
func (svc UserCrowdRankSvc) Level2Table(params traffic.UserCommonRankParams) ([]*traffic.RankLevel2Table, error) {
	_, userIds, err := svc.getRelations(params.CrowdIdList)
	if err != nil {
		svc.Log.Error("获取用户与用户组关系失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}
	params.UserIdList = userIds

	output, err := svc.rankSvc.level2Table(params)
	if err != nil {
		svc.Log.Error("查询Level2Table数据失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}
	return output, nil
}

// Level3Table 返回前端需要的Level3表格数据
func (svc UserCrowdRankSvc) Level3Table(params traffic.UserCommonRankParams) ([]*traffic.RankLevel3Table, error) {
	_, userIds, err := svc.getRelations(params.CrowdIdList)
	if err != nil {
		svc.Log.Error("获取用户与用户组关系失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}
	params.UserIdList = userIds

	output, err := svc.rankSvc.level3Table(params)
	if err != nil {
		svc.Log.Error("查询Level3Table数据失败", zap.Any("params", params), zap.Error(err))
		return nil, err
	}
	return output, nil
}

func (svc UserCrowdRankSvc) Export(params traffic.UserCommonRankParams) ([]byte, error) {
	return svc.rankSvc.export(params, svc.Level1Table)
}

// aggregatedLevel1Base 返回聚合后的Level1Base数据, byId表示是否只根据ID聚合，否则根据ID和StartTime进行聚合
func (svc UserCrowdRankSvc) aggregatedLevel1Base(params traffic.UserCommonRankParams, onlyByID bool) ([]*traffic.RankLevel1Base, error) {
	relations, userIds, err := svc.getRelations(params.CrowdIdList)
	if err != nil {
		return nil, fmt.Errorf("获取用户与用户组关系失败: %w", err)
	}
	params.UserIdList = userIds

	sql := svc.rankSvc.getLevel1BaseSqlByTimeRange(params)
	var input []*traffic.RankLevel1Base
	if err := svc.ClickHouse.Raw(sql).Scan(&input).Error; err != nil {
		return nil, fmt.Errorf("查询Level1Base数据失败: %w", err)
	}

	aggregated := svc.rankSvc.aggregate(input, relations, onlyByID)

	// 填充Name
	if err := svc.rankSvc.fillName(aggregated, svc.getNameMap); err != nil {
		svc.Log.Error("填充name失败", zap.Error(err))
	}
	return aggregated, nil
}

// getRelations 获取用户与所属用户组的关系 {user_id: [crowd_ids]}
func (svc UserCrowdRankSvc) getRelations(ids []int64) (map[int64][]int64, []int64, error) {
	sql := `
			SELECT u.id             AS user_id,
       			   JSON_ARRAYAGG(c.id) AS crowd_ids
			FROM dim_user_info u
					 JOIN
				 dim_user_crowd_relation uc ON u.id = uc.user_id
					 JOIN
				 dim_user_crowd c ON uc.crowd_id = c.id
			%s
			GROUP BY u.id
			ORDER BY u.id
			`

	var whereSQL string
	if len(ids) != 0 {
		toSQL, _, _ := goqu.From().Where(goqu.L("c.id").In(ids)).ToSQL()
		whereSQL = toSQL[len("SELECT *"):]
	}
	sql = fmt.Sprintf(sql, whereSQL)

	var relations []struct {
		UserId   int64   `gorm:"column:user_id"`
		CrowdIds []int64 `gorm:"column:crowd_ids;type:json;json"`
	}

	if err := svc.Mysql.Raw(sql).Scan(&relations).Error; err != nil {
		return nil, nil, fmt.Errorf("获取用户与用户组关系失败: %w", err)
	}

	output := make(map[int64][]int64, len(relations))
	for _, item := range relations {
		output[item.UserId] = item.CrowdIds
	}

	var userIds []int64
	if len(ids) != 0 {
		userIds = lo.Keys(output)
	}

	return output, userIds, nil
}

func (svc UserCrowdRankSvc) getNameMap(ids []int64) (map[int64]string, error) {
	type field struct {
		Id   int64  `gorm:"column:id"`
		Name string `gorm:"column:crowd_name"`
	}

	var records []*field
	err := svc.Mysql.Table("dim_user_crowd").Select("id, crowd_name").Where("id IN (?)", ids).Scan(&records).Error
	if err != nil {
		return nil, fmt.Errorf("根据用户组ID查询用户组名失败: %w", err)
	}

	var m = make(map[int64]string, len(records))
	for _, item := range records {
		m[item.Id] = item.Name
	}

	m[0] = "未知"
	m[-1] = "其他"

	return m, nil
}
