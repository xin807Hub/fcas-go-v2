package object

import (
	"fcas_server/model/common/request"
	"gorm.io/gorm"
)

type AppClassifyListRequest struct {
	request.PageInfo

	AppTypeId int `json:"appTypeId,omitempty" form:"appTypeId,omitempty"`
	AppId     int `json:"appId,omitempty" form:"appId,omitempty"`
}

type AppClassify struct {
	Id        int `json:"id"`
	DeletedAt gorm.DeletedAt

	AppTypeId   int    `json:"appTypeId"`
	AppTypeName string `json:"appTypeName"`
	AppId       int    `json:"appId"`
	AppName     string `json:"appName"`
}

func (AppClassify) TableName() string {
	return "dim_app_classify"
}
