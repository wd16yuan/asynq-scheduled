package config

import "time"

type System struct {
	DbType    string        `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	Interval  time.Duration `mapstructure:"interval" json:"interval" yaml:"interval"`
	Retention time.Duration `mapstructure:"retention" json:"retention" yaml:"retention"`
	Retry     int           `mapstructure:"retry" json:"retry" yaml:"retry"`
}
