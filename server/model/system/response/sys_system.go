package response

import "fcas_server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
