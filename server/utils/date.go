package utils

import (
	"errors"
	"fcas_server/global"
	"go.uber.org/zap"
	"time"
)

// getHoursBetween 计算从 startTime 到 endTime 的小时差值
func getHoursBetween(startTime, endTime string) (int64, error) {
	// 将 startTime 和 endTime 从字符串转换为 time.Time
	startTimeParsed, err := time.Parse(global.DateTimeLayout, startTime)
	if err != nil {
		global.Log.Error("failed to parse start time: %w", zap.Error(err))
		return 0, err
	}

	endTimeParsed, err := time.Parse(global.DateTimeLayout, endTime)
	if err != nil {
		global.Log.Error("failed to parse end time: %w", zap.Error(err))
		return 0, err
	}

	// 计算两个时间之间的差值
	duration := endTimeParsed.Sub(startTimeParsed)
	return int64(duration.Hours()), nil
}

// GetHoursBetween 计算从 startTime 到 endTime 的小时差值
func GetHoursBetween(startTime, endTime string) (int64, error) {
	// 将 startTime 和 endTime 从字符串转换为 time.Time
	startTimeParsed, err := time.Parse(global.DateTimeLayout, startTime)
	if err != nil {
		global.Log.Error("failed to parse start time: %w", zap.Error(err))
		return 0, err
	}

	endTimeParsed, err := time.Parse(global.DateTimeLayout, endTime)
	if err != nil {
		global.Log.Error("failed to parse end time: %w", zap.Error(err))
		return 0, err
	}

	// 计算两个时间之间的差值
	duration := endTimeParsed.Sub(startTimeParsed)
	return int64(duration.Hours()), nil
}

// GetDaysBetween 计算从 startTime 到 endTime 的天数差值
func GetDaysBetween(startTime, endTime string) (int, error) {
	hours, err := getHoursBetween(startTime, endTime)
	if err != nil {
		global.Log.Error("failed to parse end time: %w", zap.Error(err))
		return 0, err
	}
	// 将差值转换为天数
	days := int(hours / 24)
	return days, nil
}

// GetParticleByTimeRange 根据开始时间与结束时间获取颗粒度
func GetParticleByTimeRange(startTime, endTime string) (int, error) {
	start2end, err := GetDaysBetween(startTime, endTime) // 开始时间和结束时间的范围（1h,1d,2d)
	if err != nil {
		return 0, err
	}
	start2now, err := GetDaysBetween(startTime, time.Now().Format(global.DateTimeLayout)) // 开始时间距离现在的范围（7,31）
	if err != nil {
		return 0, err
	}

	if start2now <= global.SevenDays {
		if start2end <= global.OneDays {
			return global.Interval10mParticle, nil
		} else if start2end > global.OneDays && start2end <= global.TwoDays {
			return global.Interval1hParticle, nil
		} else {
			return global.Interval1dParticle, nil
		}
	} else if start2now > global.SevenDays && start2now <= global.OneMonth {
		start2endHours, e := getHoursBetween(startTime, endTime)
		if e != nil {
			global.Log.Error("计算开始时间与结束时间之间相差的小时数失败", zap.Error(e))
			return 0, e
		}
		if start2endHours < 1 {
			return 0, errors.New("包含7天以外的数据，查询的时间范围必须大于1小时")
		}
		if start2end <= global.TwoDays {
			return global.Interval1hParticle, nil
		} else {
			return global.Interval1dParticle, nil
		}
	} else {
		if start2end < global.OneDays {
			return 0, errors.New("包含31天以外的数据，查询的时间范围必须大于1天")
		}
		if start2end > global.OneMonth {
			return 0, errors.New("包含31天以外的数据，查询的时间范围不得大于31天")
		}
		return global.Interval1dParticle, nil
	}
}

// GetDbTypeByTimeRange 查询时间对比部署时间，选择合适的数据库进行查询
func GetDbTypeByTimeRange(startTime, endTime string) string {
	deploymentDate := global.CONFIG.DeploymentDate

	// 解析 startTime
	start, err := time.Parse(global.DateTimeLayout, startTime)
	if err != nil {
		global.Log.Error("failed to parse start time", zap.Error(err))
		return ""
	}

	// 解析 endTime
	end, err := time.Parse(global.DateTimeLayout, endTime)
	if err != nil {
		global.Log.Error("failed to parse start time: %w", zap.Error(err))
		return ""
	}

	// 比较时间并返回相应的状态
	if start.After(deploymentDate) {
		return global.QueryNew
	} else if start.Before(deploymentDate) && end.After(deploymentDate) {
		return global.QueryCrossOld2New
	} else if deploymentDate.After(end) {
		return global.QueryOld
	} else {
		return global.QueryNew
	}
}
