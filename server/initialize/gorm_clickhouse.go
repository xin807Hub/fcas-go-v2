package initialize

import (
	"fcas_server/global"
	"fcas_server/initialize/internal"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func GormClickhouse() {
	c := global.CONFIG.ClickHouse

	openDB := func(dsn string) (*gorm.DB, error) {
		clickhouseConfig := clickhouse.Config{
			DSN:                       dsn,   // DSN data source name
			SkipInitializeWithVersion: false, // 根据版本自动配置
		}
		db, err := gorm.Open(clickhouse.New(clickhouseConfig), &gorm.Config{
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

	v1db, err := openDB(c.DsnV1())
	if err != nil {
		global.Log.Error("clickhouse connect failed", zap.String("db_name", c.DbName.V1), zap.Error(err))
		os.Exit(0)
	}
	v2db, err := openDB(c.DsnV2())
	if err != nil {
		global.Log.Error("clickhouse connect failed", zap.String("db_name", c.DbName.V2), zap.Error(err))
		os.Exit(0)
	}

	global.V1ClickhouseDB = v1db
	global.V2ClickhouseDB = v2db
	global.Log.Info("clickhouse connect success", zap.String("db_name", fmt.Sprintf("%s, %s", c.DbName.V1, c.DbName.V2)))
}
