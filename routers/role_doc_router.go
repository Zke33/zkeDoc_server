package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) RoleDocRouter() {
	app := api.App.RoleDocApi
	r := router.Group("role_docs").Use(middleware.JwtAdmin)
	r.GET(":id", app.RoleDocListView)
	r.POST("", app.RoleDocCreateView)
	r.DELETE("", app.RoleDocRemoveView)
	r.GET("info", app.RoleDocInfoView)
	r.PUT("info", app.RoleDocInfoUpdateView)
	r.PUT("", app.RoleDocUpdateView)

	nr := router.Group("role_docs")
	nr.GET("", app.RoleDocTreeView)
}
