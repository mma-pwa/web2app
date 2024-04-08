package client

import (
	"github.com/gin-gonic/gin"
	v1 "web2app/api"
)

type DomainApiRoute struct {
}

func (r *DomainApiRoute) InitDomainApiRoute(Router *gin.RouterGroup) {

	domainRouter := Router.Group("domain")
	domainRouterApi := v1.ApiGroupApp.DomainApiGroup.DomainApi
	{
		domainRouter.GET("list", domainRouterApi.GetDomainList)
	}
}
