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

type DimControlPolicyApi struct {
	service servicePolicy.DimControlPolicyService
}

func (a DimControlPolicyApi) Router(Router *gin.RouterGroup) {
	router := Router.Group("policy/controlPolicy").Use(middleware.OperationRecord())
	router.POST("page", a.PageControlPolicy)
	router.POST("info", a.GetControlPolicy)
	router.POST("saveOrUpdate", a.SaveOrUpdate)
	router.POST("delete", a.Delete)
}

// PageControlPolicy
// @Tags      DimControlPolicyApi
// @Summary   管控策略配置分页查询
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimControlPolicyReq  true  " "
// @Success   200
// @Router    /policy/controlPolicy/page [POST]
func (a DimControlPolicyApi) PageControlPolicy(c *gin.Context) {
	var req policy.DimControlPolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	result, total, err := a.service.PageControlPolicy(req)
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

func (a DimControlPolicyApi) GetControlPolicy(c *gin.Context) {
	var req policy.DimControlPolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	result := a.service.InfoControlPolicy(req.Id)
	response.OkWithDetailed(result, "获取成功", c)
}

// SaveOrUpdate
// @Tags      DimControlPolicyApi
// @Summary   管控策略配置新增修改
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimControlPolicy  true  " "
// @Success   200
// @Router    /policy/controlPolicy/saveOrUpdate [POST]
func (a DimControlPolicyApi) SaveOrUpdate(c *gin.Context) {
	var req policy.DimControlPolicy
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	err = a.service.SaveOrUpdateControlPolicy(req)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("获取失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

// Delete
// @Tags      DimControlPolicyApi
// @Summary   管控策略配置删除
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body policy.DimControlPolicyReq.Ids  true  " "
// @Success   200
// @Router    /policy/controlPolicy/delete [GET]
func (a DimControlPolicyApi) Delete(c *gin.Context) {
	var req policy.DimControlPolicyReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error("参数绑定错误", zap.Error(err))
		response.FailWithMessage(utils.TranslateErr(err), c)
		return
	}
	err = a.service.DeleteControlPolicy(req.Ids)
	if err != nil {
		global.Log.Error("操作失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("操作失败: %s", err.Error()), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}
