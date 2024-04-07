package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web2app/global"
	"web2app/model"
	"web2app/model/common/response"
)

type UserApi struct {
}

// DeleteUser 删除Userinfo
// @Tags Userinfo
// @Summary 删除Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Userinfo true "删除Userinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userinfo/deleteUserinfo [delete]
func (userApi *UserApi) DeleteUser(c *gin.Context) {
	var userinfo model.User
	_ = c.ShouldBindJSON(&userinfo)
	if err := userService.DeleteUser(&userinfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateUserinfo 更新Userinfo
// @Tags Userinfo
// @Summary 更新Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Userinfo true "更新Userinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userinfo/updateUserinfo [put]
func (userApi *UserApi) UpdateUserinfo(c *gin.Context) {
	var userinfo model.User
	_ = c.ShouldBindJSON(&userinfo)
	if err, _ := userService.UpdateUser(&userinfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserinfo 用id查询Userinfo
// @Tags Userinfo
// @Summary 用id查询Userinfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Userinfo true "用id查询Userinfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userinfo/findUserinfo [get]
func (userApi *UserApi) FindUserinfo(c *gin.Context) {
	id := c.Param("id")
	if err, reuserinfo := userService.GetUser(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserinfo": reuserinfo}, c)
	}
}
