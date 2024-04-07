package service

import "web2app/service/client"

var (
	AppService              = client.NewAppService()
	JwtService              = client.NewJwtService()
	UsersService            = client.NewUsersService()
	UpgradeService          = client.NewUpgradeService()
	AppCustomURLService     = client.NewAppCustomURLService()
	AppUserLoginLogService  = client.NewAppUserLoginLogService()
	AppUserReportDayService = client.NewAppUserReportDayService()
	DomainService           = client.NewDomainService()
)
