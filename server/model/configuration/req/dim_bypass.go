package req

type DimBypassSaveRequest struct {
	ID         int    `json:"id"`
	OlpId      int    `json:"olpId"` // bypass编号
	BypassName string `json:"bypassName"`
	BypassIp   string `json:"bypassIp"`
	BypassPort int    `json:"bypassPort"`
	Remark     string `json:"remark"`
}
