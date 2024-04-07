package client

import (
	"github.com/gin-gonic/gin"
	"web2app/model"
	"web2app/model/common/response"
	"web2app/model/request"
	"web2app/utils"
)

type AppCustomURLApi struct {
}

// CreateAppCustomUR
// @Tags app
// @Summary 创建APP
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "创建APP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /app/create [post]
func (appCustomURLApi *AppCustomURLApi) CreateAppCustomUR(c *gin.Context) {
	var appCustomURL model.AppCustomURL
	if err := c.BindJSON(&appCustomURL); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
		return
	}
	if err := appCustomURLService.CreateAppCustomURL(appCustomURL); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteAppCustomURL 删除APP
// @Tags app
// @Summary 删除Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Userinfo true "删除APP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/delete [delete]
func (appCustomURLApi *AppCustomURLApi) DeleteAppCustomURL(c *gin.Context) {
	var app model.AppCustomURL
	_ = c.ShouldBindJSON(&app)
	if err := appCustomURLService.DeleteAppCustomURL(app); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateAppCustomURL 更新app
// @Tags app
// @Summary 更新Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.APP true "更新app"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/update [put]
func (appCustomURLApi *AppCustomURLApi) UpdateAppCustomURL(c *gin.Context) {
	var app model.AppCustomURL
	_ = c.ShouldBindJSON(&app)
	if err := appCustomURLService.UpdateAppCustomURL(app); err != nil {

		response.FailWithDetailed(err.Error(), "更新失败", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// GetAppCustomList 分页获取App列表
// @Tags app
// @Summary 分页获取App列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.AppSearch true "分页获取App列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/list [get]
func (appCustomURLApi *AppCustomURLApi) GetAppCustomList(c *gin.Context) {
	var pageInfo request.AppCustomURLSearch
	if err := c.BindJSON(&pageInfo); err != nil {
		response.FailTokenDetailed(utils.Unauthoriz, "", err.Error(), c)
		return
	}
	userId, b := c.Get("userId")
	if b {
		pageInfo.UserId = userId.(string)
	}
	if err, list, total := appCustomURLService.GetAppCustomList(pageInfo); err != nil {
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
