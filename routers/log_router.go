package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) LogRouter() {
	app := api.App.LogApi
	r := router.Group("logs").Use(middleware.LogMiddleWare())
	r.GET("", app.LogListView)
	r.GET("read", app.LogReadView)
	r.DELETE("", app.LogRemoveView)
}
