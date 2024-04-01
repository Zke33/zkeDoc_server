package routers

import "gvd_server/api"

func (router RouterGroup) DataRouter() {
	app := api.App.DataApi
	router.GET("data/sum", app.DataSumApiView)
}
