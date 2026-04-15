package utils

import (
	"fcas_server/global"
	"fmt"
	"github.com/pebbe/zmq4"
	"os"
	"path/filepath"
	"time"
)

// SendMessage 发送ZeroMQ消息并备份策略
func SendMessage(jsonStr, url, msgType, policyDir string) {
	// 创建ZMQ上下文
	context, _ := zmq4.NewContext()
	defer func(context *zmq4.Context) {
		err := context.Term()
		if err != nil {
			global.Log.Error(err.Error())
		}
	}(context)

	// 创建PUB socket
	socket, _ := context.NewSocket(zmq4.PUB)
	defer socket.Close()

	// 绑定地址
	if err := socket.Bind(url); err != nil {
		global.Log.Error(fmt.Sprintf("ZeroMQ bind failed: %v", err))
	}

	// 等待连接建立
	time.Sleep(1 * time.Second)

	// 发送消息
	message := []byte(fmt.Sprintf("%s%s", msgType, jsonStr))
	if _, err := socket.SendBytes(message, zmq4.DONTWAIT); err != nil {
		global.Log.Error(fmt.Sprintf("ZeroMQ send failed: %v", err))
	}

	// 备份策略文件
	bakPolicyJson(jsonStr, msgType, policyDir)
}

// bakPolicyJson 备份策略到文件
func bakPolicyJson(str, channel, policyDir string) {
	// 创建目录（如果不存在）
	if err := os.MkdirAll(policyDir, 0755); err != nil {
		panic(fmt.Sprintf("Create dir failed: %v", err))
	}

	// 生成文件名
	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s_%d.json",
		channel,
		timestamp,
		time.Now().UnixNano()/1e6)

	filePath := filepath.Join(policyDir, fileName)

	// 写入文件
	if err := os.WriteFile(filePath, []byte(str), 0644); err != nil {
		global.Log.Error(fmt.Sprintf("Write file failed: %v", err))
	}
}
