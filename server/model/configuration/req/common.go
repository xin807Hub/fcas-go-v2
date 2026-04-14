package req

import "fcas_server/model/common/request"

type ListRequest struct {
	request.PageInfo
	Key string `json:"key" form:"key"`
}
