package flags

import (
	"fmt"
	"gvd_server/global"
)

func Port(port int) {
	global.Config.System.Port = port
	fmt.Println("初始化端口")
}
