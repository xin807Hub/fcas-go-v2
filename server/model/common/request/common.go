package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page  int `json:"page" form:"page"`   // 页码
	Limit int `json:"limit" form:"limit"` // 每页大小
}
