package main

import (
	"fcas_server/core"
	"fcas_server/global"
	"fcas_server/initialize"
	"fcas_server/plugin/email"
	"fmt"
	_ "go.uber.org/automaxprocs"
	"os"
)

var buildTime string
var buildVer string

func version() {
	verCmd := false

	args := os.Args

	for i := 1; i < len(args); i++ {
		if args[i] == "-v" {
			verCmd = true
		}
	}

	if verCmd {
		fmt.Printf("buildTime: %s\n", buildTime)
		fmt.Printf("buildVer: %s\n", buildVer)
		os.Exit(0)
	}

	return
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=http://192.168.3.243:8001,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate swag init --parseDependency --parseInternal

func main() {
	version()
	initialize.Config()     // 初始化配置
	initialize.Zap()        // 初始化日志库
	initialize.Gorm()       // gorm连接数据库
	initialize.BlackCache() // 加载黑名单缓存
	initialize.Timer()      // 初始化定时器

	initialize.AppMap()    // AppTypeMap,AppMap
	initialize.AddrSet()   // 初始化地址集合
	initialize.TransInit() // 翻译器注册

	email.CreateEmailPlug(global.CONFIG.Emain) // 注册邮件插件

	core.RunWindowsServer()
}
