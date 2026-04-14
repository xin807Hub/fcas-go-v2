package configuration

import (
	"fcas_server/global"
	"fcas_server/model/common/response"
	configuration2 "fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fcas_server/service/configuration"
	"fmt"
	"github.com/gin-gonic/gin"
)

type dimBypassApi struct {
	svc *configuration.DimBypassSvc
}

func NewDimBypassRouter(rg *gin.RouterGroup) {
	router := rg.Group("/configuration/dimbypass")

	api := dimBypassApi{
		svc: configuration.NewDimBypassSvc(global.Log, global.ServiceDB),
	}

	router.GET("info/:id", api.Info)        // /configuration/dimbypass/info/:id
	router.GET("list", api.List)            // /configuration/dimbypass/list
	router.POST("save", api.Save)           // /configuration/dimbypass/save
	router.POST("update", api.Update)       // /configuration/dimbypass/update
	router.POST("delete", api.Delete)       // /configuration/dimbypass/delete
	router.POST("setStatus", api.SetStatus) // /configuration/dimbypass/setStatus
}

// Info
// @Tags      dimbypass
// @Router    /configuration/dimbypass/info/:id [get]
func (a dimBypassApi) Info(c *gin.Context) {
	id := c.Param("id")
	item, err := a.svc.GetById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(item, c)
}

// List
// @Tags      dimbypass
// @Router    /configuration/dimbypass/list [get]
func (a dimBypassApi) List(c *gin.Context) {
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
// @Tags      dimbypass
// @Router    /configuration/dimbypass/save [post]
func (a dimBypassApi) Save(c *gin.Context) {
	var req req.DimBypassSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.BypassName, req.ID); err != nil {
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
// @Tags      dimbypass
// @Router    /configuration/dimbypass/update [post]
func (a dimBypassApi) Update(c *gin.Context) {
	var req req.DimBypassSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.BypassName, req.ID); err != nil {
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
// @Tags      dimbypass
// @Router    /configuration/dimbypass/delete [post]
func (a dimBypassApi) Delete(c *gin.Context) {
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

// SetStatus
// @Tags      dimbypass
// @Router    /configuration/dimbypass/setStatus [post]
func (a dimBypassApi) SetStatus(c *gin.Context) {
	var bypass configuration2.DimBypass
	if err := c.ShouldBind(&bypass); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.svc.SetBypassStatus(bypass); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.Ok(c)
}
