package policy

import "time"

type DwsLinkApp1Min struct {
	StartTime            time.Time `json:"start_time"`
	LinkId               int       `json:"link_id"`
	AppType              int       `json:"app_type"`
	AppId                int       `json:"app_id"`
	AppTrafficUp         int       `json:"app_traffic_up"`
	AppTrafficDn         int       `json:"app_traffic_dn"`
	AppTrafficTotal      int       `json:"app_traffic_total"`
	AppTrafficUpSpeed    int       `json:"app_traffic_up_speed"`
	AppTrafficDnSpeed    int       `json:"app_traffic_dn_speed"`
	AppTrafficTotalSpeed int       `json:"app_traffic_total_speed"`

	AppName     string `json:"app_name" gorm:"-"`
	AppTypeName string `json:"app_type_name" gorm:"-"`
	LinkName    string `json:"link_name" gorm:"-"`
	Email       string `json:"email" gorm:"-"`
	AlarmType   int    `json:"alarm_type" gorm:"-"`
	PolicyId    int    `json:"policy_id" gorm:"-"`
	PolicyName  string `json:"policy_name" gorm:"-"`
}

func (DwsLinkApp1Min) TableName() string {
	return "dws_link_app_1min"
}
