package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type UpgradeRouter struct {
}

func (s *UpgradeRouter) InitUpgradeRouter(Router *gin.RouterGroup) {
	bcWithoutRecord := Router.Group("upgrade")
	var bcApi = v1.ApiGroupApp.UpgradeApiGroup.UpgradeApi
	{
		bcWithoutRecord.POST("upload/img", bcApi.UploadUpgradeFile)

	}
}
