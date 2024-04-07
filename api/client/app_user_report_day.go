package client

import (
	"github.com/gin-gonic/gin"
	"web2app/model/common/response"
	"web2app/model/request"
)

type AppUserReportDayApi struct {
}

func (appUserReportDayApi *AppUserReportDayApi) GetAppUserReportDayList(c *gin.Context) {
	var pageInfo request.AppUserReportDaySearch
	if err := c.BindJSON(&pageInfo); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
		return
	}
	userId, b := c.Get("userId")
	if b {
		pageInfo.UserID = userId.(string)
	}
	if err, list, total := appUserReportDayService.GetAppUserReportDayList(pageInfo); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "success", c)
	}
}
