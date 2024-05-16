package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi
	router.POST("user", middleware.JwtAuth, app.UserCreateView)
	router.POST("login", app.UserLoginView)
	router.PUT("user", middleware.JwtAdmin, app.UserUpdateView)
	router.GET("users", middleware.JwtAdmin, app.UserListView)
	router.DELETE("users", middleware.JwtAdmin, app.UserRemoveView)
	router.GET("logout", middleware.JwtAuth, app.UserLogoutView)
	router.GET("users_info", middleware.JwtAuth, app.UserInfoView)
	router.PUT("users_password", middleware.JwtAuth, app.UserUpdatePasswordView)
	router.PUT("users_info", middleware.JwtAuth, app.UserUpdateInfoView)
}
