package configuration

import (
	"encoding/json"
	"fcas_server/utils"
	"fmt"
	"net/http"
	"time"
)

const (
	DurationGetStatus = 2 * time.Second // 获取状态延迟
	DurationSetStatus = 5 * time.Second // 设置状态延迟
)

type DimBypass struct {
	Id         int    `json:"id" gorm:"column:id"`
	OlpId      int    `json:"olpId" gorm:"column:olp_id"`           // bypass编号
	BypassName string `json:"bypassName" gorm:"column:bypass_name"` // 分流器名称
	BypassIp   string `json:"bypassIp" gorm:"column:bypass_ip"`     // 分流器IP
	BypassPort int    `json:"bypassPort" gorm:"column:bypass_port"` // 分流器交互端口
	Remark     string `json:"remark" gorm:"column:remark"`          // 备注
	Status     int    `json:"status" gorm:"-"`                      // 状态 0-停用 1-启用
}

func (r *DimBypass) TableName() string {
	return "dim_bypass"
}

func (r *DimBypass) GetStatus() error {
	// 通过HTTP请求获取状态
	/*
		curl -X POST "http://192.168.1.1:8080/olpStatus-get" \
		     -H "Content-Type: application/json" \
		     -d '{"olpId": "12345"}'
	*/
	urlStr := fmt.Sprintf("http://%s:%d/olpStatus-get", r.BypassIp, r.BypassPort)
	reqBody := map[string]any{"olpId": r.OlpId}
	respBody, err := utils.HttpRequestWithTimeout(urlStr, http.MethodPost, nil, nil, reqBody, DurationGetStatus)
	if err != nil {
		return fmt.Errorf("failed to get status of bypass %s, error: %w", r.BypassIp, err)
	}

	var data struct {
		Code int `json:"code"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return fmt.Errorf("failed to unmarshal response body of bypass %s, error: %w", r.BypassIp, err)
	}

	r.Status = data.Code

	return nil
}

func (r *DimBypass) SetStatus(status int) error {
	// 通过HTTP请求设置状态
	/*
		curl -X POST "http://192.168.1.1:8080/olpStatus-set" \
		     -H "Content-Type: application/json" \
		     -d '{"olpId": "12345", "status": 1}'
	*/
	urlStr := fmt.Sprintf("http://%s:%d/olpStatus-set", r.BypassIp, r.BypassPort)
	reqBody := map[string]any{
		"olpId":  r.OlpId,
		"status": status,
	}
	respBody, err := utils.HttpRequestWithTimeout(urlStr, http.MethodPost, nil, nil, reqBody, DurationSetStatus)
	if err != nil {
		return fmt.Errorf("failed to set status of bypass %s, error: %w", r.BypassIp, err)
	}

	var data struct {
		Code int `json:"code"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return fmt.Errorf("failed to unmarshal response body of bypass %s, error: %w", r.BypassIp, err)
	}

	if data.Code != 1 {
		return fmt.Errorf("failed to set status of bypass %s, code: %d", r.BypassIp, data.Code)
	}
	return nil
}
