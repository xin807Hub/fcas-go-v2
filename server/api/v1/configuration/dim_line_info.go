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

type dimLineInfoApi struct {
	svc *configuration.DimLineInfoSvc
}

func NewDimLineInfoRouter(rg *gin.RouterGroup) {

	router := rg.Group("configuration/dimlineinfo").Use(middleware.OperationRecord())

	api := dimLineInfoApi{
		svc: configuration.NewDimLineInfoSvc(global.Log, global.ServiceDB),
	}

	router.GET("info/:id", api.Info)  // /configuration/dimlineinfo/info/:id
	router.GET("list", api.List)      // /configuration/dimlineinfo/list
	router.POST("save", api.Save)     // /configuration/dimlineinfo/save
	router.POST("update", api.Update) // /configuration/dimlineinfo/update
	router.POST("delete", api.Delete) // /configuration/dimlineinfo/delete
}

// Info
// @Tags      dimlineinfo
// @Router    /configuration/dimbypass/info/:id [get]
func (a dimLineInfoApi) Info(c *gin.Context) {
	id := c.Param("id")
	item, err := a.svc.GetById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(item, c)
}

// List
// @Tags      dimlineinfo
// @Router    /configuration/dimlineinfo/list [get]
func (a dimLineInfoApi) List(c *gin.Context) {
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
// @Tags      dimlineinfo
// @Router    /configuration/dimlineinfo/save [post]
func (a dimLineInfoApi) Save(c *gin.Context) {
	var req req.DimLineInfoSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.LineName, req.LineVlan, req.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	if err := a.svc.Save(req); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.Ok(c)
}

// Update
// @Tags      dimlineinfo
// @Router    /configuration/dimlineinfo/update [post]
func (a dimLineInfoApi) Update(c *gin.Context) {
	var req req.DimLineInfoSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.svc.ValidateUnique(req.LineName, req.LineVlan, req.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	if err := a.svc.Update(req); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.Ok(c)
}

// Delete
// @Tags      dimlineinfo
// @Router    /configuration/dimlineinfo/delete [post]
func (a dimLineInfoApi) Delete(c *gin.Context) {
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
