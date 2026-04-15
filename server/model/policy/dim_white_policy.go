package policy

import (
	"fcas_server/global"
	"fcas_server/model/common/request"
	"gorm.io/gorm"
	"time"
)

const ControlPolicyMsgType = 1284
const WhitePolicyMsgType = 1285

type DimWhitePolicyReq struct {
	request.PageInfo
	Id              int    `json:"id" form:"id"`
	Ids             []int  `json:"ids" form:"ids"`
	PolicyName      string `json:"policyName" form:"policyName"`
	UerType         string `json:"uerType" form:"uerType"`
	UerCrowdGroupId string `json:"uerCrowdGroupId" form:"uerCrowdGroupId"`
	UerCrowdId      string `json:"uerCrowdId" form:"uerCrowdId"`
	UerName         string `json:"uerName" form:"uerName"`
	UerId           string `json:"uerId" form:"uerId"`
	AppTypeId       int    `json:"appTypeId" form:"appTypeId"`
	AppId           int    `json:"appId" form:"appId"`
}

type DimWhitePolicy struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	UserType         int       `json:"user_type"`
	UserCrowdGroupId int       `json:"user_crowd_group_id"`
	UserCrowdId      int       `json:"user_crowd_id"`
	UserId           int       `json:"user_id"`
	UlTos            *int      `json:"ul_tos"`
	DlTos            *int      `json:"dl_tos"`
	AppTypeId        int       `json:"app_type_id"`
	AppTypeName      string    `json:"app_type_name" gorm:"-"`
	AppId            int       `json:"app_id"`
	AppName          string    `json:"app_name" gorm:"-"`
	StartTime        string    `json:"start_time"`
	EndTime          string    `json:"end_time"`
	Remark           string    `json:"remark"`
	CreateTime       time.Time `json:"create_time"`
}

func (d *DimWhitePolicy) TableName() string {
	return "dim_white_policy"
}

func (d *DimWhitePolicy) AfterFind(db *gorm.DB) (err error) {
	d.AppTypeName = global.AppTypeMap[d.AppTypeId]
	d.AppName = global.AppMap[d.AppId]
	return nil
}

type DimWhitePolicyResp struct {
	DimWhitePolicy
	UserCrowdGroupName string `json:"user_crowd_group_name"`
	UserCrowdName      string `json:"user_crowd_name"`
	UserName           string `json:"user_name"`
	AppTypeName        string `json:"app_type_name"`
	AppName            string `json:"app_name"`
}
