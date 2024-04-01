package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) RoleRouter() {
	app := api.App.RoleApi
	r := router.Group("roles").Use(middleware.JwtAdmin)
	r.GET("", app.RoleListView)
	r.POST("", app.RoleCreateView)
	r.PUT("", app.RoleUpdateView)
	r.DELETE("", app.RoleRemoveView)

}
