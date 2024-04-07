package request

import (
	"web2app/model"
	"web2app/model/common/request"
)

type AppSearch struct {
	request.PageInfo
	model.App
}
