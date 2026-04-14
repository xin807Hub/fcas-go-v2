package object

import (
	"fcas_server/global"
	"fcas_server/model/common/response"
	objectModel "fcas_server/model/object"
	"fcas_server/service/object"
	"fmt"
	"github.com/gin-gonic/gin"
)

type appClassifyApi struct {
	svc *object.AppClassifySvc
}

func NewAppClassifyRouter(rg *gin.RouterGroup) {

	router := rg.Group("object/appClassify")

	api := appClassifyApi{
		svc: object.NewAppClassifySvc(global.Log, global.ServiceDB),
	}

	router.GET("list", api.List)      // /object/appClassify/list
	router.POST("import", api.Import) // /object/appClassify/import

}

// List
// @Tags      appClassify-应用分类与自定义
// @Param     data  body  objectModel.AppClassifyListRequest  true  " "
// @Router    /object/appClassify/list [get]
func (a appClassifyApi) List(c *gin.Context) {
	var req objectModel.AppClassifyListRequest
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

// Import
// @Tags      appClassify-应用分类与自定义
// @Router    /object/appClassify/import [post]
func (a appClassifyApi) Import(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.svc.Import(file); err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}

	response.OkWithMessage("导入成功", c)
}
