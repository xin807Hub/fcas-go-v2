package initialize

import (
	"fcas_server/global"
	"fcas_server/service/object"
	"go.uber.org/zap"
	"os"
	"sync"
)

func AppMap() {

	svc := object.NewAppClassifySvc(global.Log, global.ServiceDB)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := svc.InitGlobalAppMap(); err != nil {
			global.Log.Error("初始化AppMap失败", zap.Error(err))
			os.Exit(1)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := svc.InitGlobalAppTypeMap(); err != nil {
			global.Log.Error("初始化AppTypeMap失败", zap.Error(err))
			os.Exit(1)
		}
	}()

	wg.Wait()

	global.Log.Info("初始化AppMap成功", zap.Any("appMap_size", len(global.AppMap)), zap.Any("appTypeMap_size", len(global.AppTypeMap)))
}
