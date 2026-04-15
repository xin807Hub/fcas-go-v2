package configuration

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	configuration2 "fcas_server/model/configuration"
	"fcas_server/model/configuration/req"
	"fcas_server/service/configuration"
	"fcas_server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type dimBypassApi struct {
	svc *configuration.DimBypassSvc
}

func NewDimBypassRouter(rg *gin.RouterGroup) {
	router := rg.Group("/configuration/dimbypass").Use(middleware.OperationRecord())

	api := dimBypassApi{
		svc: configuration.NewDimBypassSvc(global.Log, global.ServiceDB),
	}

	router.GET("info/:id", api.Info)                                  // /configuration/dimbypass/info/:id
	router.GET("list", api.List)                                      // /configuration/dimbypass/list
	router.POST("save", api.Save)                                     // /configuration/dimbypass/save
	router.POST("update", api.Update)                                 // /configuration/dimbypass/update
	router.POST("delete", api.Delete)                                 // /configuration/dimbypass/delete
	router.POST("setStatus", api.SetStatus)                           // /configuration/dimbypass/setStatus
	router.POST("validateBypassPassword", api.ValidateBypassPassword) // /configuration/dimbypass/validateBypassPassword
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
	var params req.ListRequest
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := a.svc.List(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(response.PageResult{
		List:       list,
		TotalCount: total,
		CurrPage:   params.Page,
		PageSize:   params.Limit,
	}, c)
}

// Save
// @Tags      dimbypass
// @Router    /configuration/dimbypass/save [post]
func (a dimBypassApi) Save(c *gin.Context) {
	var params req.DimBypassSaveRequest
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(params.BypassName, params.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	if err := a.svc.Save(params); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.Ok(c)
}

// Update
// @Tags      dimbypass
// @Router    /configuration/dimbypass/update [post]
func (a dimBypassApi) Update(c *gin.Context) {
	var params req.DimBypassSaveRequest
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(params.BypassName, params.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	if err := a.svc.Update(params); err != nil {
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

// ValidateBypassPassword
// @Tags      dimbypass
// @Router    /configuration/dimbypass/validateBypassPassword [post]
func (a dimBypassApi) ValidateBypassPassword(c *gin.Context) {
	var params struct {
		BypassPassword string `json:"bypass_password"`
	}
	if err := c.ShouldBind(&params); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 检查bypassPassword是否正确
	currUserId := utils.GetUserID(c)
	validateResult, err := a.svc.ValidateBypassPassword(currUserId, params.BypassPassword)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败: %v", err), c)
		return
	}

	if !validateResult {
		// 校验Bypass密码错误
		response.OkWithData(map[string]any{
			"succ": false,
			"msg":  "Bypass切换密码错误，请重新输入",
		}, c)
		return
	}

	response.OkWithData(map[string]any{
		"succ": true,
	}, c)
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
		response.FailWithMessage(fmt.Sprintf("操作失败: %v", err), c)
		return
	}

	response.Ok(c)
}
