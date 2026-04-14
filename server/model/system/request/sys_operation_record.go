package request

import (
	"fcas_server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	PageInfo
}
