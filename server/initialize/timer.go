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
		var option []cron.Option
		option = append(option, cron.WithSeconds())

		if global.CONFIG.Timer.ClearDbTask.Enable {
			// 每天进行一次清理数据库的任务
			// 清理DB定时任务
			_, err := global.Timer.AddTaskByFunc("ClearDBTask", global.CONFIG.Timer.ClearDbTask.Spec, func() {
				err := task.ClearTable(global.SystemDB) // 定时任务方法定在task文件包中
				if err != nil {
					fmt.Println("timer error:", err)
				}
			}, "定时清理数据库【日志，黑名单】内容", option...)
			if err != nil {
				global.Log.Error("add timer error:", zap.Error(err))
			}
		}

		if global.CONFIG.Timer.UserAlarmTask.Enable {
			// 10分钟执行一次告警定时任务
			_, err := global.Timer.AddTaskByFunc("UserAlarmTask", global.CONFIG.Timer.UserAlarmTask.Spec, func() {
				task.RunUserAlarmTask()
			}, "用户监控告警任务", option...)
			if err != nil {
				global.Log.Error("add timer error:", zap.Error(err))
			}
		}

		if global.CONFIG.Timer.PolicyTask.Enable {
			// 每15分钟执行一次定时下发任务
			_, err := global.Timer.AddTaskByFunc("PolicyTask", global.CONFIG.Timer.PolicyTask.Spec, func() {
				task.PolicyAutoSendTask()
			}, "过期策略定时解绑以及延期策略定时下发任务", option...)
			if err != nil {
				global.Log.Error("add timer error:", zap.Error(err))
			}
		}

		if global.CONFIG.Timer.PolicyActionIdTask.Enable {
			// 每5s从ftp读取actionId的任务
			_, err := global.Timer.AddTaskByFunc("PolicyActionIdTask", global.CONFIG.Timer.PolicyActionIdTask.Spec, func() {
				task.RunPolicyActionIdTask()
			}, "每5s从ftp读取actionId的任务", option...)
			if err != nil {
				global.Log.Error("add timer error:", zap.Error(err))
			}
		}

		if global.CONFIG.Timer.PolicyLogTask.Enable {
			// 每5s从分流拉取日志的任务
			_, err := global.Timer.AddTaskByFunc("PolicyLogTask", global.CONFIG.Timer.PolicyLogTask.Spec, func() {
				task.RunPolicyTrafficLogTask()
			}, "每5s从分流拉取日志任务", option...)
			if err != nil {
				global.Log.Error("add timer error:", zap.Error(err))
			}
		}
	}()
}
