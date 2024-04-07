package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type UsersRouter struct {
}

// InitUsersRouter 初始化 User 路由信息
func (s *UsersRouter) InitUsersRouter(Router *gin.RouterGroup) {
	usersRouter := Router.Group("user")
	usersRouterWithoutRecord := Router.Group("user")
	var usersApi = v1.ApiGroupApp.UserApiGroup.UserApi
	{

		usersRouter.DELETE("delete", usersApi.DeleteUser)
		usersRouter.PUT("update", usersApi.UpdateUserinfo)
	}
	{
		usersRouterWithoutRecord.GET("info/:id", usersApi.FindUserinfo) // 根据ID获取Users

	}

}
