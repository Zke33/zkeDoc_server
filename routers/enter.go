package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvd_server/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func Routers() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiGroup := router.Group("api")
	apiGroup.Use(middleware.LogMiddleWare())
	routerGroup := RouterGroup{apiGroup}
	router.Static("/uploads", "uploads")
	routerGroup.UserRouter()
	routerGroup.ImageRouter()
	routerGroup.LogRouter()
	routerGroup.SiteRouter()
	routerGroup.RoleRouter()
	routerGroup.DocRouter()
	routerGroup.RoleDocRouter()
	routerGroup.DataRouter()
	return router
}
