package global

import (
	"fcas_server/utils/addr_set"
	"fcas_server/utils/timer"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"fcas_server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	SystemDB           *gorm.DB
	ServiceDB          *gorm.DB
	V1ClickhouseDB     *gorm.DB
	V2ClickhouseDB     *gorm.DB
	REDIS              redis.UniversalClient
	CONFIG             config.Server
	Viper              *viper.Viper
	Log                *zap.Logger
	Timer              = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{}
	ROUTERS            gin.RoutesInfo
	BlackCache         local_cache.Cache

	AddrSet *addr_set.AddrSet // 用于用户管理功能中的IP地址管理

	AppMap     map[int]string // 用于根据AppId获取AppName
	AppTypeMap map[int]string // 用于根据AppTypeId获取AppTypeName

	Trans ut.Translator
)
