package home

type DwsLink10minView struct {
	StartTime string  `json:"startTime" gorm:"start_time;comment:话单时间"`
	LinkId    string  `json:"linkId" gorm:"link_id;comment:话单时间"`
	UpByte    uint64  `json:"upByte" gorm:"up_byte;comment:上行总流量"`
	DnByte    uint64  `json:"dnByte" gorm:"dn_byte;comment:上行总流量"`
	TotalByte uint64  `json:"totalByte" gorm:"total_byte;comment:总流量"`
	UpBps     float64 `json:"upBps" gorm:"up_bps;comment:上行10min的平均流速"`
	DnBps     float64 `json:"dnBps" gorm:"dn_bps;comment:下行10min的平均流速"`
	TotalBps  float64 `json:"totalBps" gorm:"total_bps;comment:10min总平均流速"`
}

func (DwsLink10minView) TableName() string {
	return "dws_link_10min_view"
}
