package router

import (
	"web2app/router/client"
)

type RouterGroup struct {
	Client client.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
