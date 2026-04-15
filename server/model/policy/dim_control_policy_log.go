package policy

import "time"

type DimControlPolicyLog struct {
	Id         int       `json:"id" gorm:"primarykey"`
	RecordTime time.Time `json:"record_time"`
	PolicyId   int       `json:"policy_id"`
	UpTraffic  int       `json:"up_traffic"`
	DnTraffic  int       `json:"dn_traffic"`
	UpPass     int       `json:"up_pass"`
	DnPass     int       `json:"dn_pass"`
	UpDiscard  int       `json:"up_discard"`
	DnDiscard  int       `json:"dn_discard"`
}

func (d *DimControlPolicyLog) TableName() string {
	return "dim_control_policy_log"
}
