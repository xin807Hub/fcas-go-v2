package configuration

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	"fcas_server/model/configuration/req"
	"fcas_server/service/configuration"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type dimUserCrowdGroupApi struct {
	svc *configuration.DimUserCrowdGroupSvc
}

func NewDimUserCrowdGroupRouter(rg *gin.RouterGroup) {

	router := rg.Group("configuration/dimusercrowdgroup").Use(middleware.OperationRecord())

	api := dimUserCrowdGroupApi{
		svc: configuration.NewDimUserCrowdGroupSvc(global.Log, global.ServiceDB),
	}

	router.GET("info/:id", api.Info)                    // /configuration/dimusercrowdgroup/info/:id
	router.GET("list", api.List)                        // /configuration/dimusercrowdgroup/list
	router.POST("save", api.Save)                       // /configuration/dimusercrowdgroup/save
	router.POST("update", api.Update)                   // /configuration/dimusercrowdgroup/update
	router.POST("delete", api.Delete)                   // /configuration/dimusercrowdgroup/delete
	router.GET("getGroupTree/:level", api.GetGroupTree) // /configuration/dimusercrowdgroup/getGroupTree/:level
}

// Info
// @Tags      dimusercrowdgroup
// @Router    /configuration/dimusercrowdgroup/info/:id [get]
func (a dimUserCrowdGroupApi) Info(c *gin.Context) {
	id := c.Param("id")
	item, err := a.svc.GetById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(item, c)
}

// List
// @Tags      dimusercrowdgroup
// @Router    /configuration/dimusercrowdgroup/list [get]
func (a dimUserCrowdGroupApi) List(c *gin.Context) {
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
// @Tags      dimusercrowdgroup
// @Router    /configuration/dimusercrowdgroup/save [post]
func (a dimUserCrowdGroupApi) Save(c *gin.Context) {
	var req req.DimUserCrowdGroupSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.GroupName, req.ID); err != nil {
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
// @Tags      dimusercrowdgroup
// @Router    /configuration/dimusercrowdgroup/update [post]
func (a dimUserCrowdGroupApi) Update(c *gin.Context) {
	var req req.DimUserCrowdGroupSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.svc.ValidateUnique(req.GroupName, req.ID); err != nil {
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
// @Tags      dimusercrowdgroup
// @Router    /configuration/dimusercrowdgroup/delete [post]
func (a dimUserCrowdGroupApi) Delete(c *gin.Context) {
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

// GetGroupTree
// @Tags      dimusercrowdgroup
// @Router    /configuration/dimusercrowdgroup/getGroupTree/:level [get]
func (a dimUserCrowdGroupApi) GetGroupTree(c *gin.Context) {
	param := c.Param("level")
	level, _ := strconv.Atoi(param)
	tree, err := a.svc.GetGroupTree(level)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(tree, c)
}

// Export
// @Tags      dimusercrowdgroup
// @Router    /configuration/dimusercrowdgroup/export [get]
func (a dimUserCrowdGroupApi) Export(c *gin.Context) {

	// TODO 导出功能实现
	//response.OkWithData(item, c)
}
