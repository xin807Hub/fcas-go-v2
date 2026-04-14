package traffic

import (
	"fcas_server/model/common/chart"
	"fcas_server/model/common/response"
)

const (
	RankTypeDstIP = "dstIp"
	RankTypeHost  = "host"
)

// AppIdReqParam 小类业务流量分析 通用请求参数
// swagger:model
type AppIdReqParam struct {
	CommonTrafficReqParam
	AppIdList []uint32 `json:"appIdList" form:"appIdList"`
	// topN 排名参数
	TopN uint8 `json:"topN" form:"topN"`
	// @Enum dstIp,host
	RankType   string `json:"rankType" form:"rankType"`
	DstIpParam string `json:"dstIpParam" form:"dstIpParam"`
	HostParam  string `json:"hostParam" form:"hostParam"`
}

type AppIdLevel1TableData struct {
	AppId     uint32  `json:"appId" gorm:"app_id"`
	Name      string  `json:"name" gorm:"name"`
	MaxUpBps  float64 `json:"maxUpBps" gorm:"max_up_bps"`
	MaxDnBps  float64 `json:"maxDnBps" gorm:"max_dn_bps"`
	AvgUpBps  float64 `json:"avgUpBps" gorm:"avg_up_bps"`
	AvgDnBps  float64 `json:"avgDnBps" gorm:"avg_dn_bps"`
	UpByte    uint64  `json:"upByte" gorm:"up_byte"`
	DnByte    uint64  `json:"dnByte" gorm:"dn_byte"`
	TotalByte uint64  `json:"totalByte" gorm:"total_byte"`
}

// AppIdLevel2RankTable 小类业务流量排名表格（目的ip和host）
type AppIdLevel2RankTable struct {
	Name      string  `json:"name" gorm:"name"`
	DstAreaId uint32  `json:"dstAreaId,omitempty" gorm:"dst_area_id"`
	Location  string  `json:"location,omitempty" gorm:"location"`
	UpByte    float64 `json:"upByte" gorm:"up_byte"`
	DnByte    float64 `json:"dnByte" gorm:"dn_byte"`
	AvgUpBps  float64 `json:"avgUpBps" gorm:"avg_up_bps"`
	AvgDnBps  float64 `json:"avgDnBps" gorm:"avg_dn_bps"`
}

type AppIdLevel2Data struct {
	TrendData    []chart.FlowTrendDot `json:"trendData" form:"trendData"`
	DstIpRankBar []chart.FlowBarDot   `json:"dstIpRankBar" form:"dstIpRankBar"`
	HostRankBar  []chart.FlowBarDot   `json:"hostRankBar" form:"hostRankBar"`
}

type Level3Data struct {
	TrendData  []chart.FlowTrendDot `json:"pieData" form:"pieData"`
	PageResult response.PageResult  `json:"pageResult" form:"pageResult"`
}

type BizIpAddress struct {
	Id       int32  `json:"id"`
	IpStart  string `json:"ip_start"`
	IpEnd    string `json:"ip_end"`
	Country  string `json:"country"`
	Unknown  string `json:"unknown"`
	Province string `json:"province"`
	City     string `json:"city"`
	Isp      string `json:"isp"`
	IspId    int32  `json:"isp_id"`
}

func (b BizIpAddress) TableName() string {
	return "biz_ip_address"
}
