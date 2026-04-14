package traffic

import "fcas_server/model/common/request"

type UserActionReqParam struct {
	StartTime string `json:"startTime" form:"startTime" binding:"required"`
	EndTime   string `json:"endTime"   form:"endTime" binding:"required"`
	SrcIp     string `json:"srcIp"     form:"srcIp" binding:"required"`
	DstIp     string `json:"dstIp"     form:"dstIp"`
	AppId     uint64 `json:"appId"     form:"appId"`
	// @Enum dst_ip,app_id
	DataType         string `json:"dataType" form:"dataType" binding:"required"`
	request.PageInfo `binding:"required"`
}

type UserActionTable struct {
	AppId     uint64 `json:"appId,omitempty" gorm:"app_id"`
	Name      string `json:"name" gorm:"name"`
	UpByte    uint64 `json:"upByte" gorm:"up_byte"`
	DnByte    uint64 `json:"dnByte" gorm:"dn_byte"`
	TotalByte uint64 `json:"totalByte" gorm:"total_byte"`
}

type UserActionDetail struct {
	Name      string `json:"name" gorm:"name"`
	Host      string `json:"host" gorm:"host"`
	UpByte    uint64 `json:"upByte" gorm:"up_byte"`
	DnByte    uint64 `json:"dnByte" gorm:"dn_byte"`
	TotalByte uint64 `json:"totalByte" gorm:"total_byte"`
}
