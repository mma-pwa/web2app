package request

import (
	"web2app/model/common/request"
)

type AppCustomURLSearch struct {
	request.PageInfo
	Status int32  `json:"status"`
	UserId string `json:"user_id"`
}
