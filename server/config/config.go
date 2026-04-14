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
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
	// oss
	Local    Local      `mapstructure:"local" json:"local" yaml:"local"`
	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`
	// redis
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	// gorm
	Mysql          Mysql        `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	ClickHouse     ClickHouse   `mapstructure:"clickhouse" json:"clickhouse" yaml:"clickhouse"`
	DeploymentDate time.Time    `mapstructure:"deployment-date" json:"deployment-date" yaml:"deployment-date"` // 查询数据库切换临界时间点
	SelectIsp      string       `mapstructure:"select-isp" json:"select-isp" yaml:"select-isp"`                // 大类、小类选出来需要分析的运营商
	Policy         Policy       `mapstructure:"policy" json:"policy" yaml:"policy"`
	ExportLimit    int          `mapstructure:"export-limit" json:"export-limit" yaml:"export-limit"` // 导出数量限制
	Emain          config.Email `mapstructure:"email" json:"email" yaml:"email"`
	DpiConf        DpiConf      `mapstructure:"dpi" json:"dpi" yaml:"dpi"`
}
type DpiConf struct {
	MsgAddr  string `mapstructure:"msg-addr" json:"msg-addr" yaml:"msg-addr"`    // zmq地址
	PushAddr string `mapstructure:"push-addr" json:"push-addr" yaml:"push-addr"` // zmq-push地址
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"` // 存放casbin模型的相对路径
}
