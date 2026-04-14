package model

const (
	RespSucc int = 0
)

type DpiResp struct {
	Code int    `json:"code" form:"code"` // 0: 成功，其他：错误码
	Msg  string `json:"msg" form:"msg"`
}
