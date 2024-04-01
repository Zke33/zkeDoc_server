package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) DocRouter() {
	app := api.App.DocApi
	r := router.Group("docs")
	r.POST("", middleware.JwtAdmin, app.DocCreateView)
	r.GET("info/:id", middleware.JwtAdmin, app.DocInfoView)
	r.PUT(":id", middleware.JwtAdmin, app.DocUpdateView)
	r.GET(":id", app.DocContentView)
	r.POST("pwd", app.DocPwdView)
	r.GET("edit/:id", middleware.JwtAdmin, app.DocEditContentView)
	r.GET("digg/:id", app.DocDiggView)
	r.DELETE(":id", middleware.JwtAdmin, app.DocRemoveView)
	r.GET("search", app.DocSearchView)
}
