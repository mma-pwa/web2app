package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type AppCustomURLRouter struct{}

func (s *AppCustomURLRouter) InitAppCustomURLRouter(Router *gin.RouterGroup) {
	customRouter := Router.Group("custom")
	customRouterWithoutRecord := Router.Group("custom")
	customRouterApi := v1.ApiGroupApp.AppCustomURLApiGroup.AppCustomURLApi
	{
		customRouter.POST("create", customRouterApi.CreateAppCustomUR)
		customRouter.POST("delete", customRouterApi.DeleteAppCustomURL)
		customRouter.POST("update", customRouterApi.UpdateAppCustomURL)
	}
	{
		customRouterWithoutRecord.POST("list", customRouterApi.GetAppCustomList)
	}
}
