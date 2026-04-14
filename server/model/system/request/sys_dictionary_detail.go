package request

import (
	"fcas_server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	PageInfo
}
