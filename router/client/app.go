package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type AppRouter struct{}

func (s *AppRouter) InitAppRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	appRouterWithoutRecord := Router.Group("app")
	appRouterApi := v1.ApiGroupApp.AppApiGroup.AppApi
	{
		appRouter.POST("create", appRouterApi.CreateAPP)
		appRouter.POST("delete", appRouterApi.DeleteApp)
		appRouter.POST("update", appRouterApi.UpdateApp)
	}
	{
		appRouterWithoutRecord.GET("info/:id", appRouterApi.GetAppInfo)
		appRouterWithoutRecord.POST("list", appRouterApi.GetAppInfoList)
		appRouterWithoutRecord.GET("template/list", appRouterApi.GetTemplateImg)
		appRouterWithoutRecord.POST("select/list", appRouterApi.GetAppList)
	}
}
