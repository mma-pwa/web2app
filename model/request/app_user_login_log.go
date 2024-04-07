package request

import "web2app/model/common/request"

type AppUserLoginLogSearch struct {
	request.PageInfo
	UserID string `json:"user_id"`
}
