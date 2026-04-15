package configuration

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	"fcas_server/model/configuration/req"
	"fcas_server/service/configuration"
	"fmt"
	"github.com/gin-gonic/gin"
)

type dimUserCrowdApi struct {
	svc *configuration.DimUserCrowdSvc
}

func NewDimUserCrowdRouter(rg *gin.RouterGroup) {

	router := rg.Group("configuration/dimusercrowd").Use(middleware.OperationRecord())

	api := dimUserCrowdApi{
		svc: configuration.NewDimUserCrowdSvc(global.Log, global.ServiceDB),
	}

	router.GET("info/:id", api.Info)  // /configuration/dimusercrowd/info/:id
	router.GET("list", api.List)      // /configuration/dimusercrowd/list
	router.POST("save", api.Save)     // /configuration/dimusercrowd/save
	router.POST("update", api.Update) // /configuration/dimusercrowd/update
	router.POST("delete", api.Delete) // /configuration/dimusercrowd/delete
}

// Info
// @Tags      dimusercrowd
// @Router    /configuration/dimusercrowd/info/:id [get]
func (a dimUserCrowdApi) Info(c *gin.Context) {
	id := c.Param("id")
	item, err := a.svc.GetById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(item, c)
}

// List
// @Tags      dimusercrowd
// @Router    /configuration/dimusercrowd/list [get]
func (a dimUserCrowdApi) List(c *gin.Context) {
	var req req.ListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := a.svc.List(req)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(response.PageResult{
		List:       list,
		TotalCount: total,
		CurrPage:   req.Page,
		PageSize:   req.Limit,
	}, c)
}

// Save
// @Tags      dimusercrowd
// @Router    /configuration/dimusercrowd/save [post]
func (a dimUserCrowdApi) Save(c *gin.Context) {
	var req req.DimUserCrowdSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.CrowdName, req.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	// 入库
	if err := a.svc.Save(req); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.Ok(c)
}

// Update
// @Tags      dimusercrowd
// @Router    /configuration/dimusercrowd/update [post]
func (a dimUserCrowdApi) Update(c *gin.Context) {
	var req req.DimUserCrowdSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.CrowdName, req.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	// 更新
	if err := a.svc.Update(req); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.Ok(c)
}

// Delete
// @Tags      dimusercrowd
// @Router    /configuration/dimusercrowd/delete [post]
func (a dimUserCrowdApi) Delete(c *gin.Context) {
	var ids []int
	if err := c.ShouldBind(&ids); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.svc.Delete(ids); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.Ok(c)
}

// Export
// @Tags      dimusercrowd
// @Router    /configuration/dimusercrowd/export [get]
func (a dimUserCrowdApi) Export(c *gin.Context) {

	// TODO 导出功能实现
	//response.OkWithData(item, c)
}
