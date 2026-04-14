package initialize

import (
	"fcas_server/global"
	"fcas_server/utils/addr_set"
	"go.uber.org/zap"
	"os"
)

func AddrSet() {
	addrSet := addr_set.NewAddrSet()

	var addrs []struct {
		IpAddr []string `gorm:"column:ip_address;type:json;json"`
	}
	err := global.ServiceDB.Table("dim_user_info").Pluck("ip_address", &addrs).Error
	if err != nil {
		global.Log.Error("初始化AddrSet失败", zap.Error(err))
		os.Exit(1)
	}

	for _, addr := range addrs {
		addrSet.AddMultipleAddr(addr.IpAddr)
	}

	global.AddrSet = addrSet

	global.Log.Info("初始化AddrSet成功", zap.Any("size", addrSet.Size()))
}
