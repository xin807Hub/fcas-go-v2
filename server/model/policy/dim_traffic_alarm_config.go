package policy

import (
	"fcas_server/global"
	"fcas_server/model/common/request"
	"gorm.io/gorm"
)

type DimTrafficAlarmConfigReq struct {
	request.PageInfo
	Id        int    `json:"id"`
	Ids       []int  `json:"ids"`
	Name      string `json:"name"`
	AppTypeId int    `json:"app_type_id"`
	AppId     int    `json:"app_id"`
}

type DimTrafficAlarmConfig struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	StartTime         string  `json:"start_time"`
	EndTime           string  `json:"end_time"`
	AppTypeId         int     `json:"app_type_id"`
	AppTypeName       string  `json:"app_type_name" gorm:"-"`
	AppId             int     `json:"app_id"`
	AppName           string  `json:"app_name" gorm:"-"`
	LinkIds           string  `json:"link_ids"`
	BaseValue         int     `json:"base_value"`
	IncreaseRatio     float64 `json:"increase_ratio"`
	DecreaseRatio     float64 `json:"decrease_ratio"`
	Email             string  `json:"email"`
	IncreaseBaseValue int     `json:"increase_base_value"`
	DecreaseBaseValue int     `json:"decrease_base_value"`
	Deleted           *int64  `json:"deleted" gorm:"index"`
}

func (d *DimTrafficAlarmConfig) TableName() string {
	return "dim_traffic_alarm_config"
}

func (d *DimTrafficAlarmConfig) AfterFind(db *gorm.DB) (err error) {
	d.AppTypeName = global.AppTypeMap[d.AppTypeId]
	d.AppName = global.AppMap[d.AppId]
	return nil
}
