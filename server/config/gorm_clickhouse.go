package config

type ClickHouse struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
	DbName    struct {
		V1 string `yaml:"v1" mapstructure:"v1"`
		V2 string `yaml:"v2" mapstructure:"v2"`
	} `yaml:"dbname" mapstructure:"dbname"`
}

func (c *ClickHouse) Dsn() string {
	return "clickhouse://" + c.Username + ":" + c.Password + "@" + c.Path + ":" + c.Port + "/" + c.Dbname + "?" + c.Config
}

func (c *ClickHouse) DsnV1() string {
	return "clickhouse://" + c.Username + ":" + c.Password + "@" + c.Path + ":" + c.Port + "/" + c.DbName.V1 + "?" + c.Config
}

func (c *ClickHouse) DsnV2() string {
	return "clickhouse://" + c.Username + ":" + c.Password + "@" + c.Path + ":" + c.Port + "/" + c.DbName.V2 + "?" + c.Config
}
