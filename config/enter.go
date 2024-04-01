package config

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Jwt    Jwt    `yaml:"jwt"`
	Site   Site   `yaml:"site"`
	Es     Es     `yaml:"es"`
}
