package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("system")
	baseApi := v1.ApiGroupApp.BaseApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.GET("captcha", baseApi.Captcha)
		baseRouter.POST("createUser", baseApi.CreateUser)
		baseRouter.POST("sendEmail", baseApi.SendEmail)
	}
	return baseRouter
}
