package config

type System struct {
	Concurrency int `mapstructure:"concurrency" json:"concurrency" yaml:"concurrency"`
}
