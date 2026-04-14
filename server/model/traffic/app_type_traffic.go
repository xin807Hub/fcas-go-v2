package traffic

type AppTypeReqParam struct {
	CommonTrafficReqParam
	AppTypeIdList []uint32 `json:"appTypeIdList" form:"appTypeIdList"`
}

type AppTypeLevel1TableData struct {
	AppType uint32 `json:"appType" gorm:"app_type"`
	Name    string `json:"name" gorm:"name"`
	CommonLevel1TableData
}
