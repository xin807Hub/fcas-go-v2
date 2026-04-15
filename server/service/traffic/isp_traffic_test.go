package traffic

import (
	"testing"
	"time"

	"fcas_server/global"
	modelTraffic "fcas_server/model/traffic"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestGetIspDbMatchesV1DirectionAndFiltersEmptyISP(t *testing.T) {
	setupDryRunDBs(t)

	query, err := getIspDb(buildIspReqParam())
	require.NoError(t, err)

	stmt := query.Session(&gorm.Session{DryRun: true}).Find(&[]map[string]interface{}{}).Statement
	sql := stmt.SQL.String()

	require.Contains(t, sql, "SELECT d_isp AS isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn")
	require.Contains(t, sql, "ifNull(d_isp, '') != ''")
	require.Contains(t, sql, "user_id IN (?)")
	require.Contains(t, sql, "d_isp = ?")
	require.Contains(t, sql, "SELECT isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn")
	require.Contains(t, sql, "ifNull(isp, '') != ''")
	require.Contains(t, sql, "d_user_id IN (?)")
	require.Contains(t, sql, "isp = ?")
	require.Contains(t, sql, "Union Distinct")
	require.Contains(t, sql, "sum(traffic_up) AS grouped_traffic_up")
	require.Contains(t, sql, "sum(traffic_dn) AS grouped_traffic_dn")
	require.Contains(t, sql, "grouped_traffic_up AS traffic_up")
	require.Contains(t, sql, "grouped_traffic_dn AS traffic_dn")
	require.Contains(t, sql, "ifNull(isp, '') != ''")
	require.Contains(t, sql, "GROUP BY isp,start_time")
}

func TestBuildIspLevel1TableQueryFiltersEmptyISP(t *testing.T) {
	setupDryRunDBs(t)

	ckDb, err := getIspDb(buildIspReqParam())
	require.NoError(t, err)

	stmt := buildIspLevel1TableQuery(ckDb).
		Limit(20).
		Find(&[]modelTraffic.IspLevel1TableData{}).
		Statement
	sql := stmt.SQL.String()

	require.Contains(t, sql, "max(traffic_up_bps) AS max_up_bps")
	require.Contains(t, sql, "avg(traffic_up_bps) AS avg_up_bps")
	require.Contains(t, sql, "sum(traffic_up) AS up_byte")
	require.Contains(t, sql, "ifNull(isp, '') != ''")
	require.Contains(t, sql, "GROUP BY")
	require.Contains(t, sql, "isp")
}

func TestGetIspDbUsesConditionalIspForV2WithoutUserScope(t *testing.T) {
	setupDryRunDBs(t)

	param := buildIspReqParam()
	param.UserIdList = nil
	param.Isp = ""

	query, err := getIspDb(param)
	require.NoError(t, err)

	stmt := query.Session(&gorm.Session{DryRun: true}).Find(&[]map[string]interface{}{}).Statement
	sql := stmt.SQL.String()

	require.Contains(t, sql, "if(user_id != 0, d_isp, isp) AS selected_isp,start_time,bytes_up_view,bytes_dn_view")
	require.Contains(t, sql, "SELECT selected_isp AS isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn")
	require.Contains(t, sql, "ifNull(if(user_id != 0, d_isp, isp), '') != ''")
	require.NotContains(t, sql, "Union Distinct")
	require.NotContains(t, sql, "SELECT d_isp AS isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn")
	require.NotContains(t, sql, "SELECT isp,start_time,sumMerge(bytes_up_view) AS traffic_up,sumMerge(bytes_dn_view) AS traffic_dn")
}

func setupDryRunDBs(t *testing.T) {
	t.Helper()

	prevV1 := global.V1ClickhouseDB
	prevV2 := global.V2ClickhouseDB
	prevSvc := global.ServiceDB
	prevLog := global.Log
	prevCfg := global.CONFIG
	t.Cleanup(func() {
		global.V1ClickhouseDB = prevV1
		global.V2ClickhouseDB = prevV2
		global.ServiceDB = prevSvc
		global.Log = prevLog
		global.CONFIG = prevCfg
	})

	global.V1ClickhouseDB = newDryRunDB(t)
	global.V2ClickhouseDB = newDryRunDB(t)
	global.ServiceDB = newDryRunDB(t)
	global.Log = zap.NewNop()
	global.CONFIG.DeploymentDate = time.Now().Add(-24 * time.Hour)
}

func buildIspReqParam() modelTraffic.IspReqParam {
	endTime := time.Now().Truncate(time.Second)
	startTime := endTime.Add(-1 * time.Hour)

	return modelTraffic.IspReqParam{
		CommonTrafficReqParam: modelTraffic.CommonTrafficReqParam{
			StartTime:  startTime.Format(global.DateTimeLayout),
			EndTime:    endTime.Format(global.DateTimeLayout),
			UserIdList: []uint32{101},
		},
		Isp: "test-isp",
	}
}

func newDryRunDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		DryRun: true,
	})
	require.NoError(t, err)
	return db
}
