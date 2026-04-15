package configuration

import (
	"fcas_server/model/common/request"
	"fcas_server/model/configuration/req"
	"fmt"
	"testing"
)

func TestDimUserCrowdSvc_List(t *testing.T) {
	svc := NewDimUserCrowdSvc(setupLog(), setupMysql())

	result, total, err := svc.List(req.ListRequest{
		PageInfo: request.PageInfo{
			Page:  1,
			Limit: 30,
		},
		Key: "内容",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
	fmt.Println(total)
}
