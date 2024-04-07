package request

import "web2app/model/common/request"

type AppUserReportDaySearch struct {
	request.PageInfo
	UserID string
}
