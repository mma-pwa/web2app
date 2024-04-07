package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type AppUserLoginLogRouter struct {
}

func (r *AppUserLoginLogRouter) InitAppUserLoginLogRouter(Router *gin.RouterGroup) {
	userLoginLogRouter := Router.Group("log")
	userLoginLogRouterApi := v1.ApiGroupApp.AppUserLoginLogApiGroup.AppUserLoginLogApi
	{
		userLoginLogRouter.POST("user_login/list", userLoginLogRouterApi.GetAppUserLoginLogList)

	}

}
