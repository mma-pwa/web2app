package client

type RouterGroup struct {
	BaseRouter
	UsersRouter
	AppRouter
	UpgradeRouter
	AppCustomURLRouter
	AppUserLoginLogRouter
	AppUserReportDayRouter
	DomainApiRoute
}
