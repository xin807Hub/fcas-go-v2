package policy

import (
	"fcas_server/model/common/request"
)

type DimControlPolicyReq struct {
	request.PageInfo
	Id              int    `json:"id"`
	Ids             []int  `json:"ids"`
	PolicyName      string `json:"policyName" form:"policyName" `
	UerType         string `json:"uerType" form:"uerType" `
	UerCrowdGroupId string `json:"uerCrowdGroupId" form:"uerCrowdGroupId" `
	UerCrowdId      string `json:"uerCrowdId" form:"uerCrowdId" `
	UerName         string `json:"uerName" form:"uerName" `
	UerId           string `json:"uerId" form:"uerId" `
	AppTypeId       int    `json:"appTypeId" form:"appTypeId" `
	AppId           int    `json:"appId" form:"appId" `
}

type DimControlPolicy struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	UserType         int    `json:"user_type"`
	UserCrowdGroupId int    `json:"user_crowd_group_id"`
	UserCrowdId      int    `json:"user_crowd_id"`
	UserId           int    `json:"user_id"`
	UlFlowRate       int    `json:"ul_flow_rate"`
	DlFlowRate       int    `json:"dl_flow_rate"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	Remark           string `json:"remark"`
	FlowCtrlType     int    `json:"flow_ctrl_type"`
	AppTypeId        int    `json:"app_type_id"`
	AppId            int    `json:"app_id"`
	DstIp            string `json:"dst_ip"`
	DstPort          string `json:"dst_port"`
	PeriodType       int    `json:"period_type"`
	PolicyPeriod     string `json:"policy_period"`
	CreateTime       string `json:"create_time"`
	LinkIds          string `json:"link_ids"`
	Status           string `json:"status"`
}

func (d *DimControlPolicy) TableName() string {
	return "dim_control_policy"
}

type DimControlPolicyResp struct {
	DimControlPolicy
	UserCrowdGroupName string `json:"user_crowd_group_name" gorm:"-"`
	UserCrownName      string `json:"user_crown_name" gorm:"-"`
	UserName           string `json:"user_name" gorm:"-"`
	AppTypeName        string `json:"app_type_name" gorm:"-"`
	AppName            string `json:"app_name" gorm:"-"`

	UpTrafficSpeed int `json:"up_traffic_speed" gorm:"-"`  // 上行总流量
	DnTrafficSpeed int `json:"dn_traffic_speed" gorm:"-"`  // 下行总流量
	UpPassSpeed    int `json:"up_pass_speed" gorm:"-"`     // 上行通过流速
	DnPassSpeed    int `json:"dn_pass_speed" gorm:"-"`     // 下行通过流速
	UpDiscardSpeed int `json:"up_discard_speed" gorm:"-"`  // 上行丢弃流速
	DnDisCardSpeed int `json:"dn_dis_card_speed" gorm:"-"` // 下行丢弃流速
}
