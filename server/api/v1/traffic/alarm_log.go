package traffic

import (
	"fcas_server/global"
	"fcas_server/model/common/response"
	trafficModel "fcas_server/model/traffic"
	"fcas_server/service/traffic"
	"fmt"
	"github.com/gin-gonic/gin"
)

type alarmLogApi struct {
	svc *traffic.AlarmLogSvc
}

func NewAlarmLogRouter(rg *gin.RouterGroup) {

	router := rg.Group("traffic/alarmLog")

	api := alarmLogApi{
		svc: traffic.NewAlarmLogSvc(global.Log, global.ServiceDB),
	}

	router.POST("list", api.List) // /traffic/alarmLog/list

}

// List
// @Tags      alarmLog
// @Param     data  body  trafficModel.AlarmLogListRequest  true  " "
// @Router    /traffic/alarmLog/list [post]
func (a alarmLogApi) List(c *gin.Context) {
	var req trafficModel.AlarmLogListRequest
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
