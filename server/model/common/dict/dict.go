package dict

type DimDict struct {
	Id       uint64    `json:"id" form:"id" gorm:"id"`                              // 字典值
	Name     string    `json:"name" form:"name" gorm:"name"`                        // 字典描述
	Children []DimDict `json:"children,omitempty" form:"children" gorm:"type:json"` // 字典子节点
}
