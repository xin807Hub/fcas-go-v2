package policy

import (
	"fcas_server/global"
	"fcas_server/model/common/response"
	"fcas_server/model/policy"
	servicePolicy "fcas_server/service/policy"
	"fcas_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DimTrafficAlarmConfigApi struct {
	service servicePolicy.DimTrafficAlarmConfigService
}

func (a DimTrafficAlarmConfigApi) Router(Router *gin.RouterGroup) {
	router := Router.Group("policy/alarmConfig")
	router.POST("page", a.PageAlarmConfig)
	router.POST("info", a.GetAlarmConfig)
	router.POST("saveOrUpdate", a.SaveOrUpdate)
	router.POST("delete", a.Delete)
}

// PageAlarmConfig
// @Tags      DimTrafficAlarmConfigApi
// @Summary   业务流量告警分页查询
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimControlPolicyReq  true  " "
// @Success   200
// @Router    /policy/alarmConfig/page [POST]
func (a DimTrafficAlarmConfigApi) PageAlarmConfig(c *gin.Context) {
	var req policy.DimTrafficAlarmConfigReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	result, total, err := a.service.PageDimTrafficAlarmConfig(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("获取失败: %s", err.Error()), c)
		return
	}
	var pageInfo = response.PageResult{
		PageSize:   req.Limit,
		CurrPage:   req.Page,
		TotalCount: total,
		List:       result,
	}
	response.OkWithDetailed(pageInfo, "获取成功", c)
}

func (a DimTrafficAlarmConfigApi) GetAlarmConfig(c *gin.Context) {
	var req policy.DimTrafficAlarmConfigReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	result := a.service.GetById(req)
	response.OkWithDetailed(result, "获取成功", c)
}

// SaveOrUpdate
// @Tags      DimTrafficAlarmConfigApi
// @Summary   业务流量告警新增修改
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimTrafficAlarmConfig  true  " "
// @Success   200
// @Router    /policy/alarmConfig/saveOrUpdate [POST]
func (a DimTrafficAlarmConfigApi) SaveOrUpdate(c *gin.Context) {
	var req policy.DimTrafficAlarmConfig
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	err = a.service.SaveOrUpdateAlarmConfig(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("获取失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// Delete
// @Tags      DimTrafficAlarmConfigApi
// @Summary   业务流量告警删除
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimTrafficAlarmConfigReq.Ids  true  " "
// @Success   200
// @Router    /policy/alarmConfig/delete [GET]
func (a DimTrafficAlarmConfigApi) Delete(c *gin.Context) {
	var req policy.DimTrafficAlarmConfigReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	err = a.service.DeleteAlarmConfig(req.Ids)
	if err != nil {
		global.Log.Error("操作失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("操作失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}
