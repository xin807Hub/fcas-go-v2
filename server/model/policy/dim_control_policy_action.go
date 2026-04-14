package policy

import "encoding/json"

type DimControlPolicyAction struct {
	Id               int    `json:"id"`
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
		byts, _ := json.Marshal(d)
		return string(byts)
	}
	return ""
}
