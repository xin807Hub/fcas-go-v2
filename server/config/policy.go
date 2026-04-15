package config

type Policy struct {
	SendUrl   string `mapstructure:"sendUrl" json:"sendUrl" yaml:"sendUrl"`
	Dir       string `mapstructure:"dir" json:"dir" yaml:"dir"`
	Ftp       Ftp    `mapstructure:"ftp" json:"ftp" yaml:"ftp"`
	ShuntPort int    `mapstructure:"shunt-port" json:"shunt-port" yaml:"shunt-port"`
}

type Ftp struct {
	IP       string `mapstructure:"ip" json:"ip" yaml:"ip"`
	UserName string `mapstructure:"userName" json:"userName" yaml:"userName"`
	UserPwd  string `mapstructure:"userPwd" json:"userPwd" yaml:"userPwd"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
}
