package policy

import (
	"encoding/json"
	"fcas_server/global"
	"go.uber.org/zap"
)

type DimControlPolicyAction struct {
	Id               int    `json:"id" gorm:"primarykey"` // 主键ID`
	PolicyId         int    `json:"policy_id"`
	VlanId           int    `json:"vlan_id"`
	ShuntIp          string `json:"shunt_ip"`
	UploadActionId   int    `json:"upload_action_id"`
	DownloadActionId int    `json:"download_action_id"`
	UploadDeviceId   int    `json:"upload_device_id"`
}

func (d *DimControlPolicyAction) TableName() string {
	return "dim_control_policy_action"
}

func (d *DimControlPolicyAction) ToString() string {
	if d != nil {
		bytes, err := json.Marshal(d)
		if err != nil {
			global.Log.Error("结构体序列化错误", zap.Error(err))
			return ""
		}
		return string(bytes)
	}
	return ""
}
