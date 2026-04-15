package configuration

type DimUserInfo struct {
	ID           int      `json:"id" gorm:"column:id"`
	UserName     string   `json:"userName" gorm:"column:user_name"`                  // 用户名称
	UserType     int      `json:"userType" gorm:"column:user_type"`                  // 用户类型 0-普通用户 1-监测用户
	Remark       string   `json:"remark" gorm:"column:remark"`                       // 备注
	IpAddress    []string `json:"ipAddress" gorm:"column:ip_address;type:json;json"` // IP地址
	IpAddressNum string   `json:"ipAddressNum" gorm:"column:ip_address_num"`         // IP地址数量
}

func (r *DimUserInfo) TableName() string {
	return "dim_user_info"
}
