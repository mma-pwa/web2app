package client

import (
	"github.com/gin-gonic/gin"
	"web2app/model"
	"web2app/model/common/response"
	"web2app/model/request"
)

type AppApi struct {
}

// CreateAPP
// @Tags app
// @Summary 创建APP
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "创建APP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /app/create [post]
func (appApi *AppApi) CreateAPP(c *gin.Context) {
	var app model.App
	if err := c.BindJSON(&app); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
		return
	}
	userId, b := c.Get("userId")
	if b {
		app.UserID = userId.(string)
	}
	if err := appService.CreateAPP(app); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// DeleteApp 删除APP
// @Tags app
// @Summary 删除Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Userinfo true "删除APP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/delete [delete]
func (appApi *AppApi) DeleteApp(c *gin.Context) {
	var app model.App
	_ = c.ShouldBindJSON(&app)
	if err := appService.DeleteApp(app); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// UpdateApp 更新app
// @Tags app
// @Summary 更新Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.APP true "更新app"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/update [put]
func (appApi *AppApi) UpdateApp(c *gin.Context) {
	var app model.App
	_ = c.ShouldBindJSON(&app)
	if err := appService.UpdateApp(app); err != nil {

		response.FailWithDetailed(err.Error(), "更新失败", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// GetAppInfo 用id查询app
// @Tags app
// @Summary 用id查询Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query id string true "用id查询app"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/info/:id [get]
func (appApi *AppApi) GetAppInfo(c *gin.Context) {
	id := c.Param("id")
	if err, appinfo := appService.GetAppInfo(id); err != nil {

		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithDetailed(appinfo, "success", c)
	}
}

// GetAppInfoList 分页获取App列表
// @Tags app
// @Summary 分页获取App列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.AppSearch true "分页获取App列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/list [get]
func (appApi *AppApi) GetAppInfoList(c *gin.Context) {
	var pageInfo request.AppSearch
	if err := c.BindJSON(&pageInfo); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
		return
	}
	userId, b := c.Get("userId")
	if b {
		pageInfo.UserID = userId.(string)
	}
	if err, list, total := appService.GetAppInfoList(pageInfo); err != nil {
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

func (appApi *AppApi) GetTemplateImg(c *gin.Context) {
	type TempImg struct {
		Id     int    `json:"id"`
		ImgUrl string `json:"img_url"`
	}
	imgList := make([]TempImg, 0)
	imgList = append(imgList, TempImg{Id: 1, ImgUrl: "/img/template/img1.png"})
	imgList = append(imgList, TempImg{Id: 2, ImgUrl: "/img/template/img2.png"})
	imgList = append(imgList, TempImg{Id: 3, ImgUrl: "/img/template/img2.png"})
	imgList = append(imgList, TempImg{Id: 4, ImgUrl: "/img/template/img2.png"})
	imgList = append(imgList, TempImg{Id: 5, ImgUrl: "/img/template/img2.png"})
	imgList = append(imgList, TempImg{Id: 6, ImgUrl: "/img/template/img2.png"})
	response.OkWithDetailed(imgList, "success", c)
}

func (appApi *AppApi) GetAppList(c *gin.Context) {

	userId, _ := c.Get("userId")

	if err, list := appService.GetAppList(userId.(string)); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithDetailed(list, "success", c)
	}
}
