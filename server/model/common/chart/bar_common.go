package chart

type FlowBarDot struct { // 饼图的每一块
	Name  string  `json:"name" gorm:"name"`   // 名称
	Value float64 `json:"value" gorm:"value"` // 值
}
