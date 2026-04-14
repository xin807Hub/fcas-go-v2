package traffic

import (
	"fcas_server/model/common/request"
	"fcas_server/model/traffic"
	"fmt"
	"testing"
)

func setupAlarmLogSvc() *AlarmLogSvc {
	return NewAlarmLogSvc(
		setupLog(),
		setupMysql(),
	)
}

func TestAlarmLogSvc_List(t *testing.T) {
	svc := setupAlarmLogSvc()

	params := traffic.AlarmLogListRequest{
		PageInfo: request.PageInfo{
			Page:  1,
			Limit: 10,
		},
	}

	output, total, err := svc.List(params)
	if err != nil {
		return
	}

	fmt.Println("total:", total)
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}
