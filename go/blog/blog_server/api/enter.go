package api

import "blog_server/api/site_api"

type Api struct {
	SiteApi siteApi.SiteApi
}

var App = Api{}
