package traffic

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	modelTraffic "fcas_server/model/traffic"
	"fcas_server/service/traffic"
	"fcas_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

type AppIdTrafficApi struct {
	service traffic.AppIdService
}

func (a AppIdTrafficApi) Router(Router *gin.RouterGroup) {
	router := Router.Group("traffic/appId").Use(middleware.OperationRecord())
	router.POST("rankData", a.GetAppIdeData)                // 1级/2级排名数据查询
	router.POST("trendTableData", a.GetAppIdTrendTableData) // 1级/2级/3级排名时间趋势表格数据查询
	router.POST("rankTableData", a.GetAppIdRankTableData)   // 2级排名表格数据查询
	router.POST("export", a.Export)                         // 导出数据
}

// GetAppIdeData
// @Tags      AppId小类业务流量分析
// @Summary   AppId小类业务流量查询
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.AppIdReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "趋势图 表格"
// @Router    /traffic/appId/rankData [POST]
func (a AppIdTrafficApi) GetAppIdeData(c *gin.Context) {
	var req modelTraffic.AppIdReqParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	var rankLevelData interface{}
	switch req.RankLevel {
	case global.LevelOne:
		rankLevelData, err = a.service.GetAppIdRankLevel1(req)
	case global.LevelTwo:
		rankLevelData, err = a.service.GetAppIdRankLevel2(req)
	case global.LevelThree:
		rankLevelData, err = a.service.GetAppIdRankLevel3(req)
	default:
		rankLevelData, err = a.service.GetAppIdRankLevel1(req)
	}
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败: "+err.Error(), c)
		return
	}
	response.OkWithDetailed(rankLevelData, "获取成功", c)
}

// GetAppIdTrendTableData
// @Tags      AppId小类业务流量分析
// @Summary   App小类流量查询 - 时间趋势表格数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.AppIdReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "小类维度表格"
// @Router    /traffic/appId/trendTableData [POST]
func (a AppIdTrafficApi) GetAppIdTrendTableData(c *gin.Context) {
	var req modelTraffic.AppIdReqParam
	var pageInfo response.PageResult
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	switch req.RankLevel {
	case global.LevelOne:
		pageInfo, err = a.service.GetLevel1TableData(req)
	case global.LevelTwo:
		pageInfo, err = a.service.GetLevel2Or3TrendTable(req)
	case global.LevelThree:
		pageInfo, err = a.service.GetLevel2Or3TrendTable(req)
	}
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("获取失败: %s", err.Error()), c)
		return
	}
	response.OkWithDetailed(pageInfo, "获取成功", c)
}

// GetAppIdRankTableData
// @Tags      AppId小类业务流量分析
// @Summary   App小类流量查询 - 排名表格数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.AppIdReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "小类维度表格"
// @Router    /traffic/appId/rankTableData [POST]
func (a AppIdTrafficApi) GetAppIdRankTableData(c *gin.Context) {
	var req modelTraffic.AppIdReqParam
	var pageInfo response.PageResult
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	pageInfo, err = a.service.GetLevel2RankTable(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("获取失败: %s", err.Error()), c)
		return
	}
	response.OkWithDetailed(pageInfo, "获取成功", c)
}

// Export
// @Tags      AppId小类业务流量分析
// @Summary   导出小类业务流量分析
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.AppIdReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "小类维度表格"
// @Router    /traffic/appId/export [POST]
func (a AppIdTrafficApi) Export(c *gin.Context) {
	var req modelTraffic.AppIdReqParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	fileBytes, err := a.service.ExportData(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fmt.Sprintf("应用小类排名_%s.xlsx", time.Now().Format("20060102150405"))))) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}
