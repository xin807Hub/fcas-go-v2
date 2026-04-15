package policy

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	"fcas_server/model/policy"
	servicePolicy "fcas_server/service/policy"
	"fcas_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DimWhitePolicyApi struct {
	service servicePolicy.DimWhitePolicyService
}

func (a DimWhitePolicyApi) Router(Router *gin.RouterGroup) {
	router := Router.Group("policy/whitePolicy").Use(middleware.OperationRecord())
	router.POST("page", a.PageWhitePolicy)
	router.POST("info", a.GetWhitePolicy)
	router.POST("saveOrUpdate", a.SaveOrUpdate)
	router.POST("delete", a.Delete)
}

// PageWhitePolicy
// @Tags      DimWhitePolicyApi
// @Summary   优先转发策略配置分页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimWhitePolicyReq  true  " "
// @Success   200
// @Router    /policy/whitePolicy/page [POST]
func (a DimWhitePolicyApi) PageWhitePolicy(c *gin.Context) {
	var req policy.DimWhitePolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	result, total, err := a.service.PageWhitePolicy(req)
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

func (a DimWhitePolicyApi) GetWhitePolicy(c *gin.Context) {
	var req policy.DimWhitePolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	result := a.service.GetById(req.Id)
	response.OkWithDetailed(result, "获取成功", c)
}

// SaveOrUpdate
// @Tags      DimWhitePolicyApi
// @Summary   优先转发策略配置新增修改
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimWhitePolicy  true  " "
// @Success   200
// @Router    /policy/whitePolicy/saveOrUpdate [POST]
func (a DimWhitePolicyApi) SaveOrUpdate(c *gin.Context) {
	var req policy.DimWhitePolicy
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	err = a.service.SaveOrUpdateWhitePolicy(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("获取失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// Delete
// @Tags      DimWhitePolicyApi
// @Summary   优先转发策略配置删除
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimWhitePolicyReq.Ids  true  " "
// @Success   200
// @Router    /policy/whitePolicy/delete [GET]
func (a DimWhitePolicyApi) Delete(c *gin.Context) {
	var req policy.DimWhitePolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	err = a.service.DeleteWhitePolicy(req.Ids)
	if err != nil {
		global.Log.Error("操作失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("操作失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}
