package configuration

type DimUserCrowd struct {
	ID        int    `json:"id" gorm:"column:id"`
	CrowdName string `json:"crowdName" gorm:"column:crowd_name"` // 用户群名称
	Remark    string `json:"remark" gorm:"column:remark"`        // 备注

	Users []DimUserCrowdUser `json:"users" gorm:"json;type:json;column:users;->"`
}

func (m *DimUserCrowd) TableName() string {
	return "dim_user_crowd"
}

type DimUserCrowdRelation struct {
	ID      int `json:"id" gorm:"column:id"`
	UserID  int `json:"userId" gorm:"column:user_id"`   // 用户ID
	CrowdID int `json:"crowdId" gorm:"column:crowd_id"` // 用户群ID
}

func (m *DimUserCrowdRelation) TableName() string {
	return "dim_user_crowd_relation"
}

type DimUserCrowdUser struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"` // 用户群中的用户名称
}
