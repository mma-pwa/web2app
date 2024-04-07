package client

import (
	"github.com/gin-gonic/gin"
	"web2app/model/common/response"
	"web2app/model/request"
)

type AppUserLoginLogApi struct {
}

func (appUserLoginLogApi *AppUserLoginLogApi) GetAppUserLoginLogList(c *gin.Context) {
	var pageInfo request.AppUserLoginLogSearch
	if err := c.BindJSON(&pageInfo); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
		return
	}
	userId, b := c.Get("userId")
	if b {
		pageInfo.UserID = userId.(string)
	}
	if err, list, total := appUserLoginLogService.GetAppUserLoginLogList(pageInfo); err != nil {
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
