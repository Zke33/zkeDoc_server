package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) ImageRouter() {
	app := api.App.ImageApi
	router.POST("image", middleware.JwtAuth, app.ImageUploadView)
	router.GET("images", middleware.JwtAdmin, app.ImageListView)
	router.DELETE("images", middleware.JwtAdmin, app.ImageRemoveView)
}
