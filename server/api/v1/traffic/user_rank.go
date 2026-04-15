package traffic

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	trafficModel "fcas_server/model/traffic"
	"fcas_server/service/traffic"
	"fcas_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

type userRankApi struct {
	svc *traffic.UserRankSvc
}

func NewUserRankRouter(rg *gin.RouterGroup) {

	router := rg.Group("traffic/userRank").Use(middleware.OperationRecord())
	dbName := global.CONFIG.ClickHouse.DbName
	api := userRankApi{
		svc: traffic.NewUserRankSvc(global.Log, global.ServiceDB, global.V2ClickhouseDB, dbName.V1, dbName.V2),
	}

	router.POST("level1Pie", api.Level1Pie)     // /traffic/userRank/level1Pie 一级饼图
	router.POST("level1Table", api.Level1Table) // /traffic/userRank/level1Table 一级表格

	router.POST("level2Trend", api.Level2Trend) // /traffic/userRank/level2Trend 二级趋势图
	router.POST("level2Table", api.Level2Table) // /traffic/userRank/level2Table 二级表格

	router.POST("level3Table", api.Level3Table) // /traffic/userRank/level3Table 三级表格

	router.POST("export", api.Export) // /traffic/userRank/export

}

// Level1Pie
// @Tags      UserRank-用户排名分析
// @Summary   一级饼图数据
// @Param     data  body  trafficModel.UserCommonRankParams  true  " "
// @Success   200  {object}  []traffic.RankLevel1Base  " "
// @Router    /traffic/userRank/level1Pie [post]
func (a userRankApi) Level1Pie(c *gin.Context) {
	var params trafficModel.UserCommonRankParams
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := fillParticleAndTimeRangeType(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pie, err := a.svc.Level1Pie(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(pie, c)
}

// Level1Table
// @Tags      UserRank-用户排名分析
// @Summary   一级表格数据
// @Param     data  body  trafficModel.UserCommonRankParams  true  " "
// @Success   200  {object}  []traffic.RankLevel1AggTraffic  " "
// @Router    /traffic/userRank/level1Table [post]
func (a userRankApi) Level1Table(c *gin.Context) {
	var params trafficModel.UserCommonRankParams
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := fillParticleAndTimeRangeType(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	rows, gather, err := a.svc.Level1Table(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(map[string]any{
		"rows":   rows,
		"gather": gather,
	}, c)
}

// Level2Trend
// @Tags      UserRank-用户排名分析
// @Summary   二级趋势图
// @Param     data  body  trafficModel.UserCommonRankParams  true  " "
// @Success   200  {object}  []traffic.RankLevel1Base  " "
// @Router    /traffic/userRank/level2Trend [post]
func (a userRankApi) Level2Trend(c *gin.Context) {
	var params trafficModel.UserCommonRankParams
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := fillParticleAndTimeRangeType(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	output, err := a.svc.Level2Trend(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(output, c)
}

// Level2Table
// @Tags      UserRank-用户排名分析
// @Summary   二级表格
// @Param     data  body  trafficModel.UserCommonRankParams  true  " "
// @Success   200  {object}  []traffic.RankLevel2Table  " "
// @Router    /traffic/userRank/level2Table [post]
func (a userRankApi) Level2Table(c *gin.Context) {
	var params trafficModel.UserCommonRankParams
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := fillParticleAndTimeRangeType(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	output, err := a.svc.Level2Table(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(output, c)
}

// Level3Table
// @Tags      UserRank-用户排名分析
// @Summary   三级表格
// @Param     data  body  trafficModel.UserCommonRankParams  true  " "
// @Success   200  {object}  []traffic.RankLevel3Table  " "
// @Router    /traffic/userRank/level3Table [post]
func (a userRankApi) Level3Table(c *gin.Context) {
	var params trafficModel.UserCommonRankParams
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := fillParticleAndTimeRangeType(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	output, err := a.svc.Level3Table(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(output, c)
}

// Export
// @Tags      UserRank-用户排名分析
// @Param     data  body  trafficModel.UserCommonRankParams  true  " "
// @Router    /traffic/userRank/export [post]
func (a userRankApi) Export(c *gin.Context) {
	var params trafficModel.UserCommonRankParams
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := fillParticleAndTimeRangeType(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	output, err := a.svc.Export(params)
	if err != nil {
		a.svc.Log.Error("导出失败", zap.Any("params", params), zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fmt.Sprintf("用户排名_%s.xlsx", time.Now().Format("20060102150405"))))) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", output)

}

func fillParticleAndTimeRangeType(commonParams *trafficModel.UserCommonRankParams) (err error) {
	if commonParams == nil {
		return fmt.Errorf("commonParams不能为空")
	}

	timeRangeType := utils.GetDbTypeByTimeRange(commonParams.StartTime, commonParams.EndTime)
	switch timeRangeType {
	case global.QueryOld:
		timeRangeType = "v1"
	case global.QueryNew:
		timeRangeType = "v2"
	case global.QueryCrossOld2New:
		timeRangeType = "v1v2"
	default:
		return fmt.Errorf("时间范围不正确")
	}

	commonParams.TimeRangeType = timeRangeType
	commonParams.Particle, err = utils.GetParticleByTimeRange(commonParams.StartTime, commonParams.EndTime)
	if err != nil {
		return err
	}

	return nil
}
