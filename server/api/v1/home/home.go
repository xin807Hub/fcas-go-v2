package home

import (
	"fcas_server/global"
	"fcas_server/model/common/response"
	modelHome "fcas_server/model/home"
	"fcas_server/service/home"
	"fcas_server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Api struct {
	service home.Service
}

func (a Api) Router(Router *gin.RouterGroup) {
	router := Router.Group("traffic/dwsTotalTraffic")
	router.POST("home", a.GetHomeData) // 获取首页数据：趋势图+表格
}

// GetHomeData
// @Tags      Home 首页：趋势图、表格
// @Summary   首页的全部数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body  modelHome.ReqParam  true  "开始时间 结束时间 链路"
// @Success   200   {object}  response.Response{data=modelHome.ReqParam,msg=string}  "趋势图 表格"
// @Router    /traffic/dwsTotalTraffic/home [post]
func (a Api) GetHomeData(c *gin.Context) {
	var req modelHome.ReqParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数解析或校验错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	homeData, err := a.service.GetHomeData(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(homeData, "获取成功", c)
}
