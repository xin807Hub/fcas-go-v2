package initialize

import (
	"fcas_server/global"
	"fcas_server/initialize/internal"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// GormMysql 初始化Mysql数据库
func GormMysql() {
	c := global.CONFIG.Mysql

	openDB := func(dsn string) (*gorm.DB, error) {
		mysqlConfig := mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         191,   // string 类型字段的默认长度
			SkipInitializeWithVersion: false, // 根据版本自动配置
		}
		db, err := gorm.Open(
			mysql.New(mysqlConfig),
			&gorm.Config{
				Logger: logger.New(internal.NewWriter(c.GeneralDB, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
					SlowThreshold: 200 * time.Millisecond,
					LogLevel:      c.GeneralDB.LogLevel(),
					Colorful:      true,
				}),
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   c.Prefix,
					SingularTable: c.Singular,
				},
				DisableForeignKeyConstraintWhenMigrating: true,
			})
		if err != nil {
			return nil, err
		}

		db.InstanceSet("gorm:table_options", "ENGINE="+c.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		return db, nil
	}

	systemDb, err := openDB(c.DsnSystem())
	if err != nil {
		global.Log.Error("mysql connect failed", zap.String("db_name", c.DbName.System), zap.Error(err))
		os.Exit(0)
	}
	serviceDb, err := openDB(c.DsnService())
	if err != nil {
		global.Log.Error("mysql connect failed", zap.String("db_name", c.DbName.Service), zap.Error(err))
		os.Exit(0)
	}
	global.SystemDB = systemDb
	global.ServiceDB = serviceDb
	global.Log.Info("mysql connect success", zap.String("db_name", fmt.Sprintf("%s, %s", c.DbName.System, c.DbName.Service)))
}
