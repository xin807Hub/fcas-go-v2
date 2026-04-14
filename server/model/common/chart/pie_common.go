package chart

type FlowPiePiece struct { // 饼图的每一块
	Name      string `json:"name"` // 名称
	DnByte    uint64 `json:"dnByte" gorm:"dn_byte"`
	TotalByte uint64 `json:"totalByte" gorm:"total_byte"`
}
