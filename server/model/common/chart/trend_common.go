package chart

type FlowTrendDot struct { // 趋势图的每个点
	StartTime string  `json:"startTime" gorm:"start_time"` // 日期
	DnBps     float64 `json:"dnBps" gorm:"dn_bps"`
	UpBps     float64 `json:"upBps" gorm:"up_bps"`
	TotalBps  float64 `json:"totalBps" gorm:"total_bps"`
}
