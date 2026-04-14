package traffic

import "fcas_server/model/system/request"

type IspReqParam struct {
	CommonTrafficReqParam
	IsOversea *uint8 `json:"isOversea" form:"isOversea"`
	Isp       string `json:"isp" form:"isp"`
}

type IspLevel1TableData struct {
	Isp       string `json:"isp" gorm:"isp"`
	IsOversea uint8  `json:"isOversea" gorm:"is_oversea"`
	CommonLevel1TableData
}

type IspTableReqParam struct {
	request.PageInfo `binding:"dive,required"`
	StartTime        string   `json:"startTime" form:"startTime" binding:"required"`
	EndTime          string   `json:"endTime" form:"endTime" binding:"required"`
	IspIdList        []uint32 `json:"ispIdList" form:"ispIdList"`
	LinkIdList       []uint32 `json:"linkIdList" form:"linkIdList"`
	ProvinceList     []uint32 `json:"provinceList" form:"provinceList"`
	UserIdList       []uint32 `json:"userIdList" form:"userIdList"`
	IsOversea        *uint8   `json:"isOversea" form:"isOversea"`
	Isp              string   `json:"isp" form:"isp"`
}
