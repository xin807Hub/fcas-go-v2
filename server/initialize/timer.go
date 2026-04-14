package initialize

import (
	"fcas_server/global"
	"fcas_server/task"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Timer() {
	go func() {
		// 每天进行一次清理数据库的任务
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.SystemDB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			global.Log.Error("add timer error:", zap.Error(err))
		}

		// 10分钟执行一次告警定时任务
		_, err = global.Timer.AddTaskByFunc("UserAlarm", "0 0/10 * * * ?", func() {
			task.RunUserAlarmTask()
		}, "用户监控告警任务", option...)
		if err != nil {
			global.Log.Error("add timer error:", zap.Error(err))
		}

		// 每15分钟执行一次定时下发任务
		_, err = global.Timer.AddTaskByFunc("AutoPolicyTask", "0 0/15 * * * ?", func() {
			task.PolicyAutoSendTask()
		}, "过期策略定时解绑以及延期策略定时下发任务", option...)
		if err != nil {
			global.Log.Error("add timer error:", zap.Error(err))
		}

		// 每5s从ftp读取actionId的任务
		_, err = global.Timer.AddTaskByFunc("AutoPolicyTask", "0/5 * * * * ?", func() {
			task.RunPolicyIdActionIdTask()
		}, "每5s从ftp读取actionId的任务", option...)
		if err != nil {
			global.Log.Error("add timer error:", zap.Error(err))
		}

		// 每5s从分流拉取日志的任务
		_, err = global.Timer.AddTaskByFunc("AutoPolicyTask", "2/5 * * * * ?", func() {
			task.RunPolicyTrafficLogTask()
		}, "每5s从分流拉取日志任务", option...)
		if err != nil {
			global.Log.Error("add timer error:", zap.Error(err))
		}
	}()
}
