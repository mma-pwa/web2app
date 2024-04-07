package client

import (
	"web2app/service"
)

type ApiGroup struct {
	AppApi
	BaseApi
	UserApi
	UpgradeApi
	AppCustomURLApi
	AppUserReportDayApi
	AppUserLoginLogApi
	DomainApi
}

var (
	appService              = service.AppService
	userService             = service.UsersService
	upgradeService          = service.UpgradeService
	appCustomURLService     = service.AppCustomURLService
	appUserLoginLogService  = service.AppUserLoginLogService
	appUserReportDayService = service.AppUserReportDayService
	domainService           = service.DomainService
)
