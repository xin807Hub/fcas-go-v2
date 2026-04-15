package configuration

import (
	"fcas_server/model/common/request"
	"fcas_server/model/configuration/req"
	"fmt"
	"testing"
)

func TestDimUserCrowdGroupSvc_List(t *testing.T) {
	svc := NewDimUserCrowdGroupSvc(setupLog(), setupMysql())

	result, total, err := svc.List(req.ListRequest{
		PageInfo: request.PageInfo{
			Page:  1,
			Limit: 30,
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
	fmt.Println(total)
}
