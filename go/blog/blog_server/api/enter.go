package api

import (
	siteApi "blog_server/api/site_api"
	logApi "blog_server/api/log_api"
)


type Api struct {
	SiteApi siteApi.SiteApi
	LogApi  logApi.LogApi
}

var App = Api{}
