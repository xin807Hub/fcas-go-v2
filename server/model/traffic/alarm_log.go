package traffic

import (
	"fcas_server/model/common/request"
	"time"
)

type AlarmLogListRequest struct {
	request.PageInfo

	StartTime *time.Time `json:"startTime" form:"startTime"`
	EndTime   *time.Time `json:"endTime" form:"endTime"`
	AppTypeId *int       `json:"appTypeId" form:"appTypeId"`
	AppId     *int       `json:"appId" form:"appId"`
	LinkIds   []int      `json:"linkIds" form:"linkIds"`
}

type AlarmLog struct {
	Id                   int       `json:"id"`
	StartTime            time.Time `json:"startTime"`
	PolicyId             int       `json:"policyId"`
	LinkId               int       `json:"linkId"`
	AppId                int       `json:"appId"`
	AppTypeId            int       `json:"appTypeId"`
	AlarmType            int       `json:"alarmType"` // 1-上浮；2-下降
	AppTrafficTotalSpeed int       `json:"appTrafficTotalSpeed"`
	LastPeriodTotalSpeed int       `json:"last_period_total_speed"`

	PolicyName  string `json:"policyName"`
	AppTypeName string `json:"appTypeName"`
	AppName     string `json:"appName"`
	LinkName    string `json:"linkName"`

	// 上下浮动比例
	Ratio string `json:"ratio" gorm:"-"`
}

func (AlarmLog) TableName() string {
	return "dws_traffic_alarm_log"
}
