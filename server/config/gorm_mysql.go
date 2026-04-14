package config

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
	DbName    struct {
		System  string `yaml:"system" mapstructure:"system"`
		Service string `yaml:"service" mapstructure:"service"`
	} `yaml:"dbname" mapstructure:"dbname"`
}

func (m *Mysql) DsnSystem() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName.System + "?" + m.Config
}

func (m *Mysql) DsnService() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName.Service + "?" + m.Config
}
