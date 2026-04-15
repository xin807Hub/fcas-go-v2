package traffic

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	modelTraffic "fcas_server/model/traffic"
	"fcas_server/service/traffic"
	"fcas_server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserActionTrafficApi struct {
	service traffic.UserActionService
}

func (a UserActionTrafficApi) Router(Router *gin.RouterGroup) {
	router := Router.Group("traffic/userAction").Use(middleware.OperationRecord())
	router.POST("pageData", a.GetUserActionPageData) // 表格数据查询
	router.POST("detail", a.GetUserActionDetail)     // 详情数据查询
	router.POST("export", a.Export)                  // 导出数据
}

// GetUserActionPageData
// @Tags      UserAction     用户行为分析
// @Summary   用户行为分析查询
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.UserActionReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "目的ip、应用小类 维度表格"
// @Router    /traffic/userAction/pageData [POST]
func (a UserActionTrafficApi) GetUserActionPageData(c *gin.Context) {
	var req modelTraffic.UserActionReqParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	if req.DataType == "" {
		global.Log.Error("请选择一个数据类型", zap.Error(err))
		response.FailWithMessage("请选择一个数据类型", c)
		return
	}
	var pageInfo response.PageResult
	pageInfo, err = a.service.GetUserActionTable(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(pageInfo, "获取成功", c)
}

// GetUserActionDetail
// @Tags      UserAction     用户行为分析
// @Summary   用户行为分析详情查询
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.UserActionReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "详情数据请求"
// @Router    /traffic/userAction/detail [POST]
func (a UserActionTrafficApi) GetUserActionDetail(c *gin.Context) {
	var req modelTraffic.UserActionReqParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	pageInfo, err := a.service.GetUserActionDetail(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(pageInfo, "获取成功", c)
}

// Export
// @Tags      UserAction     用户行为分析
// @Summary   用户行为分析导出
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelTraffic.UserActionReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response  "excel导出"
// @Router    /traffic/userAction/export [POST]
func (a UserActionTrafficApi) Export(c *gin.Context) {
	var req modelTraffic.UserActionReqParam
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
	response.OkWithDetailed(fileBytes, "获取成功", c)
}
