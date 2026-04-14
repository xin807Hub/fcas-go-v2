package configuration

type DimLineInfo struct {
	ID       int    `json:"id" gorm:"column:id"`              // 线路ID
	LineName string `json:"lineName" gorm:"column:line_name"` // 线路名称
	LineNum  string `json:"lineNum" gorm:"column:line_num"`   // 线路编号
	LineVlan int    `json:"lineVlan" gorm:"column:line_vlan"` // valn
	Remark   string `json:"remark" gorm:"column:remark"`      // 线路备注
}

func (r *DimLineInfo) TableName() string {
	return "dim_line_info"
}
