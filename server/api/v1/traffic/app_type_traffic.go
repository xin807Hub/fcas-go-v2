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

type AppTypeTrafficApi struct {
	service traffic.AppTypeService
}

func (a AppTypeTrafficApi) Router(Router *gin.RouterGroup) {
	router := Router.Group("traffic/appType").Use(middleware.OperationRecord())
	router.POST("rankData", a.GetAppTypeData)           // 1级/2级排名数据查询
	router.POST("rankTableData", a.GetAppTypeTableData) // 1级/2级排名数据查询
	router.POST("export", a.Export)                     // 导出数据
}

// GetAppTypeData
// @Tags      AppType大类业务流量分析
// @Summary   app大类流量查询
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.AppTypeReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "饼图 大类维度表格 汇总数据 趋势图 时间维度表格"
// @Router    /traffic/appType/rankData [POST]
func (a AppTypeTrafficApi) GetAppTypeData(c *gin.Context) {
	var req modelTraffic.AppTypeReqParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	var ispData interface{}
	switch req.RankLevel {
	case global.LevelOne:
		ispData, err = a.service.GetAppTypeRankLevel1(req)
	case global.LevelTwo:
		ispData, err = a.service.GetAppTypeRankLevel2(req)
	default:
		ispData, err = a.service.GetAppTypeRankLevel1(req)
	}

	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(ispData, "获取成功", c)
}

// GetAppTypeTableData
// @Tags      AppType大类业务流量分析
// @Summary   app大类流量表格查询
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.AppTypeReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "饼图 大类维度表格 汇总数据 趋势图 时间维度表格"
// @Router    /traffic/appType/rankTableData [POST]
func (a AppTypeTrafficApi) GetAppTypeTableData(c *gin.Context) {
	var req modelTraffic.AppTypeReqParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	var resultPage response.PageResult
	switch req.RankLevel {
	case global.LevelOne:
		resultPage, err = a.service.GetLevel1TableData(req)
	case global.LevelTwo:
		resultPage, err = a.service.GetLevel2TableData(req)
	default:
		resultPage, err = a.service.GetLevel1TableData(req)
	}
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(resultPage, "获取成功", c)
}

// Export
// @Tags      AppType大类业务流量分析
// @Summary   大类业务流量导出
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.AppTypeReqParam  true  "开始时间 结束时间 链路"
// @Success   200 {object}  response.Response  "饼图、趋势图"
// @Router    /traffic/appType/export [POST]
func (a AppTypeTrafficApi) Export(c *gin.Context) {
	var req modelTraffic.AppTypeReqParam
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
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fmt.Sprintf("应用大类排名_%s.xlsx", time.Now().Format("20060102150405"))))) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", fileBytes)
}
