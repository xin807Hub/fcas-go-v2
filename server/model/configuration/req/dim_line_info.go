package req

type DimLineInfoSaveRequest struct {
	ID       int    `json:"id"`
	LineName string `json:"lineName"`
	LineNum  string `json:"lineNum"`
	LineVlan int    `json:"lineVlan"`
	Remark   string `json:"remark"`
}
