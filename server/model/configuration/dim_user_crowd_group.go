package configuration

type DimUserCrowdGroup struct {
	ID        int    `json:"id" gorm:"column:id"`
	GroupName string `json:"group_name" gorm:"column:group_name"` // 用户群组名称
	Remark    string `json:"remark" gorm:"column:remark"`         // 备注

	Crowds []DimUserCrowdGroupCrowd `json:"crowds" gorm:"-"`
}

func (m *DimUserCrowdGroup) TableName() string {
	return "dim_user_crowd_group"
}

type DimUserCrowdGroupRelation struct {
	ID      int `json:"id" gorm:"column:id"`
	CrowdID int `json:"crowd_id" gorm:"column:crowd_id"` // 用户群id
	GroupID int `json:"group_id" gorm:"column:group_id"` // 用户群组id
}

func (m *DimUserCrowdGroupRelation) TableName() string {
	return "dim_user_crowd_group_relation"
}

type DimUserCrowdGroupCrowd struct {
	ID        int    `json:"id" gorm:"column:id"`
	CrowdName string `json:"crowdName" gorm:"column:crowd_name"` // 用户群组中的用户群名称
}

type GroupTreeNode struct {
	ID       string          `json:"id"`
	Label    string          `json:"label"`
	Children []GroupTreeNode `json:"children"`
}
