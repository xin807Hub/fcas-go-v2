package configuration

import (
	"fcas_server/model/common/response"
	"fcas_server/service/configuration"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DimDictApi struct {
	svc configuration.DimDictService
}

func (a DimDictApi) Router(Router *gin.RouterGroup) {
	router := Router.Group("configuration/dimDict")
	router.GET("infoList/:type", a.DictInfoList) // 获取字典下拉框
}

// DictInfoList
// @Tags      DimDict字典下拉框信息
// @Param     type  path string true  "字典类型" Enums(isp,ispSelect,appType,appId,appTypeIdTree)
// @Success   200  {object}  interface{}  " "
// @Router    /configuration/dimDict/infoList/{type} [get]
func (a DimDictApi) DictInfoList(c *gin.Context) {
	dictType := c.Param("type")
	item, err := a.svc.GetDictByType(dictType)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("操作失败, %v", err), c)
		return
	}
	response.OkWithData(item, c)
}
