package config

import "fmt"

type System struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (system System) Addr() string {
	return fmt.Sprintf("%s:%d", system.Ip, system.Port)
}
