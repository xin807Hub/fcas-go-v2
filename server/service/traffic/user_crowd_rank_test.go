package traffic

import (
	"fcas_server/model/traffic"
	"fmt"
	"testing"
)

func setupCrowdRankSvc() *UserCrowdRankSvc {
	return NewUserCrowdRankSvc(
		setupLog(),
		setupMysql(),
		setupClickHouse(),
		"bigdata_fcas",
		"bigdata_fcas_v2",
	)
}

func TestCrowdRankSvc_Level1Pie(t *testing.T) {
	svc := setupCrowdRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
	}

	output, err := svc.Level1Pie(params)
	if err != nil {
		return
	}

	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestCrowdRankSvc_Level2Trend(t *testing.T) {
	svc := setupCrowdRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
		CrowdIdList:   []int64{4},
	}

	output, err := svc.Level2Trend(params)
	if err != nil {
		return
	}

	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestCrowdRankSvc_Level1Table(t *testing.T) {
	svc := setupCrowdRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
	}

	output, _, err := svc.Level1Table(params)
	if err != nil {
		return
	}

	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestCrowdRankSvc_Level2Table(t *testing.T) {
	svc := setupCrowdRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
		TopN:          999,
		CrowdIdList:   []int64{4},
	}

	output, err := svc.Level2Table(params)
	if err != nil {
		return
	}

	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestCrowdRankSvc_Level3Table(t *testing.T) {
	svc := setupCrowdRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
		SrcIp:         "1.159.224.226",
	}

	output, err := svc.Level3Table(params)
	if err != nil {
		return
	}

	fmt.Println("total:", len(output))
	for _, item := range output {
		fmt.Printf("%+v\n", item)
	}
}

func TestCrowdRankSvc_getUserRelations(t *testing.T) {
	svc := setupCrowdRankSvc()

	output, _, err := svc.getRelations(nil)
	if err != nil {
		return
	}

	fmt.Println("total:", len(output))
	for id, relations := range output {
		fmt.Println("id:", id, "relations:", relations)
	}
}

func TestCrowdRankSvc_getNameMap(t *testing.T) {
	svc := setupCrowdRankSvc()

	output, err := svc.getNameMap([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err != nil {
		return
	}

	fmt.Println("total:", len(output))
	for id, name := range output {
		fmt.Println("id:", id, "name:", name)
	}
}
