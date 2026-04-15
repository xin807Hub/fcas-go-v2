package configuration

import (
	"fcas_server/global"
	"fcas_server/middleware"
	"fcas_server/model/common/response"
	"fcas_server/model/configuration/req"
	"fcas_server/service/configuration"
	"fcas_server/utils/addr_set"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type dimUserInfoApi struct {
	svc *configuration.DimUserInfoSvc
}

func NewDimUserInfoRouter(rg *gin.RouterGroup) {
	router := rg.Group("configuration/dimuserinfo").Use(middleware.OperationRecord())

	api := dimUserInfoApi{
		svc: configuration.NewDimUserInfoSvc(global.Log, global.ServiceDB, global.AddrSet),
	}

	router.GET("info/:id", api.Info)  // /configuration/dimuserinfo/info/:id
	router.GET("list", api.List)      // /configuration/dimuserinfo/list
	router.POST("save", api.Save)     // /configuration/dimuserinfo/save
	router.POST("update", api.Update) // /configuration/dimuserinfo/update
	router.POST("delete", api.Delete) // /configuration/dimuserinfo/delete
	router.GET("export", api.Export)  // /configuration/dimuserinfo/export
}

// Info
// @Tags      dimuserinfo
// @Router    /configuration/dimuserinfo/info/:id [get]
func (a dimUserInfoApi) Info(c *gin.Context) {
	id := c.Param("id")
	item, err := a.svc.GetById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithData(item, c)
}

// List
// @Tags      dimuserinfo
// @Router    /configuration/dimuserinfo/list [get]
func (a dimUserInfoApi) List(c *gin.Context) {
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
// @Tags      dimuserinfo
// @Router    /configuration/dimuserinfo/save [post]
func (a dimUserInfoApi) Save(c *gin.Context) {
	var req req.DimUserInfoSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.svc.Log.Info("saving", zap.Any("params", req))

	// 校验用户名
	if err := a.svc.ValidateUnique(req.UserName, req.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	// 校验IP地址
	if err := addr_set.ValidateAddr(req.IpAddress...); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 入库
	if err := a.svc.Save(req); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}
	a.svc.Log.Info("创建成功", zap.Any("AddrSet", global.AddrSet.String()), zap.Any("size", global.AddrSet.Size()))

	response.Ok(c)
}

// Update
// @Tags      dimuserinfo
// @Router    /configuration/dimuserinfo/update [post]
func (a dimUserInfoApi) Update(c *gin.Context) {
	var req req.DimUserInfoSaveRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.svc.Log.Info("saving", zap.Any("params", req))

	// 校验用户名
	if err := a.svc.ValidateUnique(req.UserName, req.ID); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	// 校验IP地址
	if err := addr_set.ValidateAddr(req.IpAddress...); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 更新
	if err := a.svc.Update(req); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}
	a.svc.Log.Info("更新成功", zap.Any("AddrSet", global.AddrSet.String()), zap.Any("size", global.AddrSet.Size()))

	response.Ok(c)
}

// Delete
// @Tags      dimuserinfo
// @Router    /configuration/dimuserinfo/delete [post]
func (a dimUserInfoApi) Delete(c *gin.Context) {
	var ids []int
	if err := c.ShouldBind(&ids); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.svc.Delete(ids); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}
	a.svc.Log.Info("删除成功", zap.Any("AddrSet", global.AddrSet.String()), zap.Any("size", global.AddrSet.Size()))

	response.Ok(c)
}

// Export
// @Tags      dimuserinfo
// @Router    /configuration/dimuserinfo/export [get]
func (a dimUserInfoApi) Export(c *gin.Context) {
	var req req.ListRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	output, err := a.svc.Export(req)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fmt.Sprintf("用户信息_%s.xlsx", time.Now().Format("20060102150405")))) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", output)
}
