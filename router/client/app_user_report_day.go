package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type AppUserReportDayRouter struct {
}

func (r *AppUserReportDayRouter) InitAppUserReportDayRouter(Router *gin.RouterGroup) {
	reportRouter := Router.Group("report")
	reportRouterApi := v1.ApiGroupApp.AppUserReportDayApiGroup.AppUserReportDayApi
	{
		reportRouter.POST("list", reportRouterApi.GetAppUserReportDayList)

	}

}
