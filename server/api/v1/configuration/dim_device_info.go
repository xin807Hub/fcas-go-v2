package configuration

import (
	"fcas_server/global"
	"fcas_server/model/common/response"
	"fcas_server/model/configuration/req"
	"fcas_server/service/configuration"
	"fmt"
	"github.com/gin-gonic/gin"
)

type dimDeviceInfoApi struct {
	svc *configuration.DimDeviceInfoSvc
}

func NewDimDeviceInfoRouter(rg *gin.RouterGroup) {
	router := rg.Group("configuration/dimdeviceinfo")

	api := dimDeviceInfoApi{
		svc: configuration.NewDimDeviceInfoSvc(global.Log, global.ServiceDB),
	}

	router.GET("info/:id", api.Info)  // /configuration/dimdeviceinfo/info/:id
	router.GET("list", api.List)      // /configuration/dimdeviceinfo/list
	router.POST("save", api.Save)     // /configuration/dimdeviceinfo/save
	router.POST("update", api.Update) // /configuration/dimdeviceinfo/update
	router.POST("delete", api.Delete) // /configuration/dimdeviceinfo/delete
}

// Info
// @Tags      dimdeviceinfo
// @Summary   分页文件列表
// @Router    /configuration/dimdeviceinfo/info/:id [get]
func (a dimDeviceInfoApi) Info(c *gin.Context) {
	id := c.Param("id")
	item, err := a.svc.GetById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(item, c)
}

// List
// @Tags      dimdeviceinfo
// @Summary   分页文件列表
// @Router    /configuration/dimdeviceinfo/list [get]
func (a dimDeviceInfoApi) List(c *gin.Context) {
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
// @Tags      dimdeviceinfo
// @Summary   创建
// @Router    /configuration/dimdeviceinfo/save [post]
func (a dimDeviceInfoApi) Save(c *gin.Context) {
	var req req.DimDeviceInfoSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.DeviceName, req.ID); err != nil {
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
// @Tags      dimdeviceinfo
// @Summary   更新
// @Router    /configuration/dimdeviceinfo/update [post]
func (a dimDeviceInfoApi) Update(c *gin.Context) {
	var req req.DimDeviceInfoSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.DeviceName, req.ID); err != nil {
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
// @Tags      dimdeviceinfo
// @Summary   删除
// @Router    /configuration/dimdeviceinfo/delete [post]
func (a dimDeviceInfoApi) Delete(c *gin.Context) {
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
