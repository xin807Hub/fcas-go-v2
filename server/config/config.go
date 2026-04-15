package config

import (
	"fcas_server/plugin/email/config"
	"time"
)

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
	// oss
	Local    Local      `mapstructure:"local" json:"local" yaml:"local"`
	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`
	// redis
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	// gorm
	Mysql          Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	ClickHouse     ClickHouse `mapstructure:"clickhouse" json:"clickhouse" yaml:"clickhouse"`
	DeploymentDate time.Time  `mapstructure:"deployment-date" json:"deployment-date" yaml:"deployment-date"` // 查询数据库切换临界时间点
	SelectIsp      string     `mapstructure:"select-isp" json:"select-isp" yaml:"select-isp"`                // 大类、小类选出来需要分析的运营商
	Policy         Policy     `mapstructure:"policy" json:"policy" yaml:"policy"`
	ExportLimit    int        `mapstructure:"export-limit" json:"export-limit" yaml:"export-limit"` // 导出数量限制
	SyncApi        bool       `mapstructure:"sync-api" json:"sync-api" yaml:"sync-api"`

	Emain   config.Email `mapstructure:"email" json:"email" yaml:"email"`
	DpiConf DpiConf      `mapstructure:"dpi" json:"dpi" yaml:"dpi"`
	Timer   Timer        `mapstructure:"timer" json:"timer" yaml:"timer"`
}
type DpiConf struct {
	MsgAddr  string `mapstructure:"msg-addr" json:"msg-addr" yaml:"msg-addr"`
	PushAddr string `mapstructure:"push-addr" json:"push-addr" yaml:"push-addr"`
}

type Timer struct {
	ClearDbTask        TaskTimerCnf `mapstructure:"clear-db-task" json:"clear-db-task" yaml:"clear-db-task"`
	UserAlarmTask      TaskTimerCnf `mapstructure:"user-alarm-task" json:"user-alarm-task" yaml:"user-alarm-task"`
	PolicyTask         TaskTimerCnf `mapstructure:"policy-task" json:"policy-task" yaml:"policy-task"`
	PolicyActionIdTask TaskTimerCnf `mapstructure:"policy-action-id-task" json:"policy-action-id-task" yaml:"policy-action-id-task"`
	PolicyLogTask      TaskTimerCnf `mapstructure:"policy-log-task" json:"policy-log-task" yaml:"policy-log-task"`
}

type TaskTimerCnf struct {
	Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Spec   string `mapstructure:"spec" json:"spec" yaml:"spec"`
}
