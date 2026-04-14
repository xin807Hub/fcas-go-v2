package initialize

import (
	"fcas_server/service/system"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"fcas_server/global"
	"fcas_server/utils"
)

func BlackCache() {
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

	// 从db加载jwt数据
	if global.SystemDB != nil {
		system.LoadAll()
	}

}
