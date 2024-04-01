package main

import (
	"gvd_server/core"
	_ "gvd_server/docs"
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
)

// @title 文档项目api文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)
	global.ESClient = core.InitEs()

	option := flags.Parse()
	if option.Run() {
		return
	}

	core.InitAddrDB()
	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
