package service

import "web2app/api/client"

type ApiGroup struct {
	AppApiGroup              client.ApiGroup
	BaseApiGroup             client.ApiGroup
	UserApiGroup             client.ApiGroup
	UpgradeApiGroup          client.ApiGroup
	AppCustomURLApiGroup     client.ApiGroup
	AppUserLoginLogApiGroup  client.ApiGroup
	AppUserReportDayApiGroup client.ApiGroup
	DomainApiGroup           client.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
