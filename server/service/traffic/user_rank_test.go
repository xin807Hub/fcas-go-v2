package traffic

import (
	"fcas_server/model/traffic"
	"fmt"
	"testing"
)

func setupUserRankSvc() *UserRankSvc {
	return NewUserRankSvc(
		setupLog(),
		setupMysql(),
		setupClickHouse(),
		"bigdata_fcas",
		"bigdata_fcas_v2",
	)
}

func TestUserRankSvc_Level1Pie(t *testing.T) {
	svc := setupUserRankSvc()

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

func TestUserRankSvc_Level2Trend(t *testing.T) {
	svc := setupUserRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
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

func TestUserRankSvc_Level1Table(t *testing.T) {
	svc := setupUserRankSvc()

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

func TestUserRankSvc_Level2Table(t *testing.T) {
	svc := setupUserRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
		TopN:          14,
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

func TestUserRankSvc_Level3Table(t *testing.T) {
	svc := setupUserRankSvc()

	params := traffic.UserCommonRankParams{
		Particle:      600,
		TimeRangeType: "v1",
		StartTime:     "2024-10-31 00:00:00",
		EndTime:       "2024-11-01 01:00:00",
		SrcIp:         "1.1.1.18",
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
