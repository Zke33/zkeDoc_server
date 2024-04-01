package core

import (
	"gopkg.in/yaml.v2"
	"gvd_server/config"
	"gvd_server/global"
	"log"
	"os"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalln("read yaml err: ", err.Error())
	}
	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		log.Fatalln("解析yaml err：", err.Error())
	}
	return c
}

func SetYaml() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return
	}
	err = os.WriteFile(yamlPath, byteData, 066)
	if err != nil {
		global.Log.Error(err)
		return
	}
}
