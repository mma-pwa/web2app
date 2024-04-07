package client

import (
	"github.com/gin-gonic/gin"
	"web2app/model/common/response"
)

type DomainApi struct {
}

func (domainApi *DomainApi) GetDomainList(c *gin.Context) {

	if err, list := domainService.GetDomainList(); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithDetailed(list, "success", c)
	}
}
