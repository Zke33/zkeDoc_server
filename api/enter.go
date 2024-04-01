package api

import (
	"gvd_server/api/data_api"
	"gvd_server/api/doc_api"
	"gvd_server/api/image_api"
	"gvd_server/api/log_api"
	"gvd_server/api/role_api"
	"gvd_server/api/role_doc_api"
	"gvd_server/api/site_api"
	"gvd_server/api/user_api"
)

type Api struct {
	UserApi    user_api.UserApi
	ImageApi   image_api.ImageApi
	LogApi     log_api.LogApi
	SiteApi    site_api.SiteApi
	RoleApi    role_api.RoleApi
	DocApi     doc_api.DocApi
	RoleDocApi role_doc_api.RoleDocApi
	DataApi    data_api.DataApi
}

var App = new(Api)
