package routers

import "gvd_server/api"

func (router RouterGroup) SiteRouter() {
	app := api.App.SiteApi
	r := router.Group("site")
	r.GET("", app.SiteDetailView)
	r.PUT("", app.SiteUpdateView)
}
