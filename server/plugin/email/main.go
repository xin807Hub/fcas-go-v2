package email

import (
	"fcas_server/plugin/email/config"
	"fcas_server/plugin/email/global"
	"fcas_server/plugin/email/router"
	"github.com/gin-gonic/gin"
)

type EmailPlugin struct{}

func CreateEmailPlug(config config.Email) *EmailPlugin {
	global.GlobalConfig.To = config.To
	global.GlobalConfig.From = config.From
	global.GlobalConfig.Host = config.Host
	global.GlobalConfig.Secret = config.Secret
	global.GlobalConfig.Nickname = config.Nickname
	global.GlobalConfig.Port = config.Port
	global.GlobalConfig.IsSSL = config.IsSSL
	return &EmailPlugin{}
}

func (*EmailPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (*EmailPlugin) RouterPath() string {
	return "email"
}
