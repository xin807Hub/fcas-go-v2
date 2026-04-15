package configuration

import (
	"fcas_server/model/configuration/req"
	"fmt"
	"log"
	"testing"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupLog() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func setupMysql() *gorm.DB {
	host := "192.168.4.180"
	user := "root"
	password := "Mysql@2o20..."
	dbName := "fcas_service"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb3&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	)

	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("连接数据库失败, error=" + err.Error())
	}

	return client
}

func TestDimLineInfoSvc_GetById(t *testing.T) {
	testCases := []struct {
		args string
	}{
		{
			args: "1",
		},
		{
			args: "26",
		},
		{
			args: "30",
		},
	}
	svc := NewDimLineInfoSvc(setupLog(), setupMysql())
	for _, tt := range testCases {
		t.Run(tt.args, func(t *testing.T) {

			result, err := svc.GetById(tt.args)
			if err != nil {
				t.Error(err)
				return
			}

			t.Logf("result=%+v", result)

		})
	}
}

func TestDimLineInfoSvc_Save(t *testing.T) {

	svc := NewDimLineInfoSvc(setupLog(), setupMysql())

	err := svc.Save(req.DimLineInfoSaveRequest{
		LineName: "link1",
		LineNum:  "2",
		LineVlan: 2,
		Remark:   "22222",
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("save success")

}
