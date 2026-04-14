package task

import (
	"fcas_server/global"
	modelConfiguration "fcas_server/model/configuration"
	modelPolicy "fcas_server/model/policy"
	"fcas_server/model/traffic"
	"fcas_server/plugin/email/service"
	"fcas_server/service/policy"
	"fcas_server/utils"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

const (
	PERIOD = 10
	// INCREASE_ALARM_TYPE 上浮告警
	INCREASE_ALARM_TYPE = 1
	// DECREASE_ALARM_TYPE 下降告警
	DECREASE_ALARM_TYPE = 2
)

var dimTrafficAlarmConfigService = policy.DimTrafficAlarmConfigService{}
var emailService = service.EmailService{}

type MonitorUserAlarmTask struct {
}

// RunUserAlarmTask 每10分钟执行一次任务
func RunUserAlarmTask() {
	end := time.Now().Add(-PERIOD * time.Minute)
	start := end.Add(-PERIOD * time.Minute)

	startTime := start.Format(global.DateTimeLayout)
	endTime := end.Format(global.DateTimeLayout)

	lastEnd := start
	lastStart := lastEnd.Add(-PERIOD * time.Minute)
	lastStartTime := lastStart.Format(global.DateTimeLayout)
	lastEndTime := lastEnd.Format(global.DateTimeLayout)

	// 获取有效告警配置
	alarmConfigs := dimTrafficAlarmConfigService.GetAlarmConfigValid()

	// 查询不同app上期link平均流量
	lastAppLinkPeriodTotal := getLastAppLinkPeriod(lastStartTime, lastEndTime, alarmConfigs)

	// 查询告警日志
	logList := listLinkAppByTime(startTime, endTime, alarmConfigs, lastAppLinkPeriodTotal)

	if len(logList) == 0 {
		return
	}

	global.Log.Info("告警日志条数：" + strconv.Itoa(len(logList)))

	mapDimLine := getAllLineMap()

	for i := 0; i < len(logList); i++ {
		logList[i].AppTypeName = global.AppTypeMap[logList[i].AppType]
		logList[i].AppName = global.AppMap[logList[i].AppId]
		logList[i].LinkName = mapDimLine[logList[i].LinkId]
		trafficAlarmLog := traffic.AlarmLog{}
		trafficAlarmLog.AppTypeName = global.AppTypeMap[logList[i].AppType]
		trafficAlarmLog.AppName = global.AppMap[logList[i].AppId]
		trafficAlarmLog.LinkName = mapDimLine[logList[i].LinkId]
		trafficAlarmLog.StartTime = logList[i].StartTime
		trafficAlarmLog.AppId = logList[i].AppId
		trafficAlarmLog.PolicyId = logList[i].PolicyId
		trafficAlarmLog.AppTypeId = logList[i].AppType
		trafficAlarmLog.LinkId = logList[i].LinkId
		trafficAlarmLog.AppTrafficTotalSpeed = logList[i].AppTrafficTotalSpeed
		trafficAlarmLog.LastPeriodTotalSpeed = int(lastAppLinkPeriodTotal[logList[i].AppId][logList[i].LinkId])
		// 实际流量
		trafficAlarmLog = setAlarmType(trafficAlarmLog)
		// 发送邮件
		if len(logList[i].Email) > 0 {
			policyName := logList[i].PolicyName + "策略流量告警"
			body := buildMailContent(logList[i])
			err := emailService.SendEmail(logList[i].Email, policyName, body)
			if err != nil {
				global.Log.Error("告警邮件发送失败：" + err.Error())
			} else {
				global.Log.Info("告警邮件发送成功")
			}
		}
		// 记录告警日志
		if err := global.ServiceDB.Model(&traffic.AlarmLog{}).Save(&trafficAlarmLog).Error; err != nil {
			global.Log.Error("告警日志记录失败：" + err.Error())
		}
	}

}

func listLinkAppByTime(startTime, endTime string, alarmConfig []modelPolicy.DimTrafficAlarmConfig, appLinkMap map[int]map[int]float32) (rest []modelPolicy.DwsLinkApp1Min) {
	for i := 0; i < len(alarmConfig); i++ {
		dto := alarmConfig[i]
		linkIds := strings.Split(dto.LinkIds, ",")
		for _, linkId := range linkIds {
			appId := dto.AppId
			linkIdInt, _ := strconv.ParseInt(linkId, 0, 64)
			if appLinkMap[appId] == nil || appLinkMap[appId][int(linkIdInt)] == 0 {
				continue
			}
			lastTotalSpeed := appLinkMap[appId][int(linkIdInt)]
			maxValue := lastTotalSpeed + (lastTotalSpeed * float32(dto.IncreaseRatio) / 100)
			var minValue float32
			if dto.DecreaseRatio == 0 {
				minValue = 0
			} else {
				minValue = lastTotalSpeed - (lastTotalSpeed * float32(dto.IncreaseRatio) / 100)
			}
			alarmList := getDwsLinkApp1Min(startTime, endTime, strconv.Itoa(dto.AppId), linkId, "", minValue, maxValue)
			for j := 0; j < len(alarmList); j++ {
				alarmList[j].Email = dto.Email
				alarmList[j].PolicyId = dto.Id
				alarmList[j].PolicyName = dto.Name
				if float32(alarmList[j].AppTrafficTotalSpeed) > maxValue {
					alarmList[j].AlarmType = INCREASE_ALARM_TYPE
				} else if float32(alarmList[j].AppTrafficTotalSpeed) < minValue {
					alarmList[j].AlarmType = DECREASE_ALARM_TYPE
				}
			}
			rest = append(rest, alarmList...)
		}
	}
	return rest
}

func getLastAppLinkPeriod(startTime, endTime string, alarmConfig []modelPolicy.DimTrafficAlarmConfig) map[int]map[int]float32 {
	rest := map[int]map[int]float32{}
	appLinkMap := map[int]float32{}
	linkCount := map[int]int{}
	for i := 0; i < len(alarmConfig); i++ {
		dto := alarmConfig[i]
		alarmList := getDwsLinkApp1Min(startTime, endTime, strconv.Itoa(dto.AppId), "", dto.LinkIds, 0, 0)
		if alarmList == nil || len(alarmList) == 0 {
			continue
		}
		for _, dwsLinkApp := range alarmList {
			linkId := dwsLinkApp.LinkId
			speed := appLinkMap[linkId]
			curSpeed := dwsLinkApp.AppTrafficTotalSpeed
			if speed == 0 {
				appLinkMap[linkId] = float32(curSpeed)
				linkCount[linkId] = 1
			} else {
				linkCount[linkId] = linkCount[linkId] + 1
				appLinkMap[linkId] = speed + float32(curSpeed)
			}
		}
		for key, val := range appLinkMap {
			appLinkMap[key] = val / float32(linkCount[key])
		}
		rest[dto.AppId] = appLinkMap
	}
	return rest
}

func getDwsLinkApp1Min(startTime, endTime, appId, linkId, linkIds string, minValue, maxValue float32) (result []modelPolicy.DwsLinkApp1Min) {
	var baseSQL string
	timeRangeType := utils.GetDbTypeByTimeRange(startTime, endTime)
	switch timeRangeType {
	case global.QueryOld: // v1
		baseSQL = baseAppSql(global.CONFIG.ClickHouse.DbName.V1, "dws_link_app_1min", startTime, endTime, appId, linkId, linkIds, minValue, maxValue)
		break
	case global.QueryNew: // v2
		baseSQL = baseAppSql(global.CONFIG.ClickHouse.DbName.V2, "dws_link_app_5min", startTime, endTime, appId, linkId, linkIds, minValue, maxValue)
		break
	case global.QueryCrossOld2New: // v1v2
		v1Sql := baseAppSql(global.CONFIG.ClickHouse.DbName.V1, "dws_link_app_1min", startTime, endTime, appId, linkId, linkIds, minValue, maxValue)
		v2Sql := baseAppSql(global.CONFIG.ClickHouse.DbName.V2, "dws_link_app_5min", startTime, endTime, appId, linkId, linkIds, minValue, maxValue)
		baseSQL = fmt.Sprintf(`(%s) UNION ALL (%s)`, v1Sql, v2Sql)
		break
	default:
		global.Log.Error("时间范围不正确")
	}

	if len(baseSQL) > 0 {
		if err := global.V2ClickhouseDB.Raw(baseSQL).Scan(&result).Error; err != nil {
			return nil
		}
	}
	return result
}

func baseAppSql(dbName, table, startTime, endTime, appId, linkId, linkIds string, minValue, maxValue float32) string {
	sql := `select * from %s.%s where 1=1 `
	// 构造查询条件
	sql += ` and start_time between '` + startTime + `' and '` + endTime + `'`
	sql += ` and app_id = ` + appId
	if len(linkIds) > 0 {
		sql += ` and link_id in (` + linkIds + `) `
	}
	if len(linkId) > 0 {
		sql += ` and link_id = ` + linkId
	}
	if minValue > 0 || maxValue > 0 {
		sql += ` and app_traffic_total_speed not between ` + fmt.Sprintf("%v", minValue) + ` and ` + fmt.Sprintf("%v", maxValue)
	}
	return fmt.Sprintf(sql, dbName, table)
}

func setAlarmType(log traffic.AlarmLog) traffic.AlarmLog {
	appTrafficTotalSpeed := log.AppTrafficTotalSpeed
	lastPeriodTotalSpeed := log.LastPeriodTotalSpeed
	if appTrafficTotalSpeed >= lastPeriodTotalSpeed {
		log.AlarmType = INCREASE_ALARM_TYPE
		ratio := float32((appTrafficTotalSpeed - lastPeriodTotalSpeed) / lastPeriodTotalSpeed)
		log.Ratio = fmt.Sprintf("%.2f", ratio*100)
	} else {
		log.AlarmType = DECREASE_ALARM_TYPE
		ratio := float32((lastPeriodTotalSpeed - appTrafficTotalSpeed) / lastPeriodTotalSpeed)
		log.Ratio = fmt.Sprintf("%.2f", ratio*100)
	}
	return log
}

func buildMailContent(linkApp modelPolicy.DwsLinkApp1Min) string {
	content := "<html>\n" +
		" <body>\n" +
		" \t<div style=\"border:1px solid #E7ECEF;font-size:14px;\">\n" +
		" \t<div style=\"border-bottom:1px solid #E7ECEF;background-color:#FCFCFC;padding:10px;\">流量告警</div>\n" +
		" \t<div style=\"padding:10px;\">\n" +
		"\t\t<p style=\"color:red;\">告警策略：#policyName</p>\n" +
		"\t\t<p>应用大类：#appTypeName</p>\n" +
		"\t\t<p>应用小类：#appName</p>\n" +
		"\t\t<p>链路：#linkName</p>\n" +
		"\t\t<p>告警类型：#alarmType</p>\n" +
		"\t\t<p style=\"color:red;\">实际流速：#appTrafficTotalSpeed</p>\n" +
		"\t</div>\n" +
		" </body>\n" +
		"</html>"
	content = strings.ReplaceAll(content, "#policyName", linkApp.PolicyName)
	content = strings.ReplaceAll(content, "#appTypeName", linkApp.AppTypeName)
	content = strings.ReplaceAll(content, "#appName", linkApp.AppName)
	content = strings.ReplaceAll(content, "#linkName", linkApp.LinkName)
	alarmType := ""
	if linkApp.AlarmType == 1 {
		alarmType = "上浮告警"
	} else {
		alarmType = "下降告警"
	}
	content = strings.ReplaceAll(content, "#alarmType", alarmType)
	content = strings.ReplaceAll(content, "#appTrafficTotalSpeed", utils.ComputeSpeed(linkApp.AppTrafficTotalSpeed))
	return content
}

func getAllLineMap() map[int]string {
	var restDto []modelConfiguration.DimLineInfo
	if err := global.ServiceDB.Model(modelConfiguration.DimLineInfo{}).Find(&restDto).Error; err != nil {
		global.Log.Error("查询DimLineInfo失败", zap.Error(err))
		return nil
	}

	rest := make(map[int]string)
	for i := 0; i < len(restDto); i++ {
		dto := restDto[i]
		rest[dto.LineVlan] = dto.LineName
	}
	return rest
}
