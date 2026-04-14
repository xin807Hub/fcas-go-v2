package home

import (
	"fcas_server/model/common/chart"
)

type ReqParam struct {
	LinkIdList []uint32 `json:"linkIdList" form:"linkIdList"`
	StartTime  string   `json:"startTime" form:"startTime" binding:"required"`
	EndTime    string   `json:"endTime" form:"endTime" binding:"required"`
}

type RespHomeData struct {
	Series    []chart.FlowTrendDot `json:"series" form:"series"`
	TableData []LinkTableData      `json:"tables" form:"tables"`
}

type LinkTableData struct {
	LinkName string  `json:"linkName" gorm:"link_name"`
	MaxUpBps float64 `json:"maxUpBps" gorm:"max_up_bps"`
	MaxDnBps float64 `json:"maxDnBps" gorm:"max_dn_bps"`
	AvgUpBps float64 `json:"avgUpBps" gorm:"avg_up_bps"`
	AvgDnBps float64 `json:"avgDnBps" gorm:"avg_dn_bps"`
}
