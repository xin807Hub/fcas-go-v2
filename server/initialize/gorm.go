package initialize

import (
	"os"

	"fcas_server/global"
	"fcas_server/model/system"

	"go.uber.org/zap"
)

func Gorm() {

	GormMysql()
	GormClickhouse()
}

func RegisterTables() {
	db := global.SystemDB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
	)
	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	db.AutoMigrate()

	if err != nil {
		global.Log.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
}
