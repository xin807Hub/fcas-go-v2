package response

type PageResult struct {
	PageSize   int         `json:"pageSize"`
	TotalCount int64       `json:"totalCount"`
	CurrPage   int         `json:"currPage"`
	List       interface{} `json:"list"`
}
