package traffic

import (
	"fcas_server/model/common/chart"
	"fcas_server/model/common/request"
)

// CommonTrafficReqParam 流量分析 通用请求参数
// swagger:model
type CommonTrafficReqParam struct {
	// @Enum level1,level2,level3
	RankLevel string `json:"rankLevel" form:"rankLevel" binding:"required"`
	request.PageInfo
	StartTime   string   `json:"startTime" form:"startTime" binding:"required"`
	EndTime     string   `json:"endTime" form:"endTime" binding:"required"`
	IspNameList []string `json:"ispNameList" form:"ispNameList"`
	LinkIdList  []uint32 `json:"linkIdList" form:"linkIdList"`
	DstProvince string   `json:"dstProvince" form:"dstProvince"`
	UserIdList  []uint32 `json:"userIdList" form:"userIdList"`
}

type Level1Data struct {
	PieData    []chart.FlowPiePiece `json:"pieData" form:"pieData"`
	GatherData GatherData           `json:"gatherData" form:"gatherData"`
}

type CommonLevel1TableData struct {
	MaxUpBps  float64 `json:"maxUpBps" gorm:"max_up_bps"`
	MaxDnBps  float64 `json:"maxDnBps" gorm:"max_dn_bps"`
	AvgUpBps  float64 `json:"avgUpBps" gorm:"avg_up_bps"`
	AvgDnBps  float64 `json:"avgDnBps" gorm:"avg_dn_bps"`
	UpByte    uint64  `json:"upByte" gorm:"up_byte"`
	DnByte    uint64  `json:"dnByte" gorm:"dn_byte"`
	TotalByte uint64  `json:"totalByte" gorm:"total_byte"`
}

type GatherData struct {
	//MaxUpBps  float64 `json:"maxUpBps" gorm:"max_up_bps"`
	//MaxDnBps  float64 `json:"maxDnBps" gorm:"max_dn_bps"`
	AvgUpBps  float64 `json:"avgUpBps" gorm:"avg_up_bps"`
	AvgDnBps  float64 `json:"avgDnBps" gorm:"avg_dn_bps"`
	UpByte    uint64  `json:"upByte" gorm:"up_byte"`
	DnByte    uint64  `json:"dnByte" gorm:"dn_byte"`
	TotalByte uint64  `json:"totalByte" gorm:"total_byte"`
}
