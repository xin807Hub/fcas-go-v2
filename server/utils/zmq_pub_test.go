package utils

import "testing"

func TestZmqSend(t *testing.T) {
	// 示例调用参数
	policyJson := `{"policy": "example"}`
	sendUrl := "tcp://*:6901"
	msgType := "1284"
	policyDir := "conf/policies"
	// 调用发送方法
	SendMessage(policyJson, sendUrl, msgType, policyDir)
}
