package client

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"web2app/global"
	"web2app/model"
	"web2app/model/common/response"
	"web2app/model/request"
	"web2app/utils"
)

type BaseApi struct {
}

// CreateUser 创建用户
// @Tags user
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/create [post]
func (baseApi *BaseApi) CreateUser(c *gin.Context) {
	var userinfo model.User
	_ = c.BindJSON(&userinfo)
	code, b := global.Dc.Get(userinfo.Email)
	if !b {
		response.FailWithDetailed("verification code is err", "fail", c)
		return
	}
	if code.(string) != userinfo.Code {
		response.FailWithDetailed("verification code is err", "fail", c)
		return
	}
	if err := userService.CreateUser(&userinfo); err != nil {
		response.FailWithDetailed(err.Error(), "fail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}

// Login
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /system/login [post]
func (baseApi *BaseApi) Login(c *gin.Context) {
	var l request.Login
	_ = c.BindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &model.User{Email: l.Email, Password: l.PassWord}
		if err, user := userService.Login(u); err != nil {
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			baseApi.tokenNext(c, *user)
		}
	} else {

		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (baseApi *BaseApi) tokenNext(c *gin.Context, user model.User) {
	expiresTime := global.GVA_CONFIG.JWT.ExpiresTime
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(model.BaseClaims{
		Email: user.Email,
		ID:    user.ID,
	}, expiresTime)
	token, err := j.CreateToken(claims)
	token = utils.TokenCrypt(token, time.Now().Format(time.DateOnly))
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	res := model.SysUserRes{Email: user.Email, ID: user.ID, Money: user.Money}
	//tree, _ := menuService.GetMenuTree(user.RoleId)
	response.OkWithDetailed(response.LoginResponse{
		//MemuTree:  tree,
		User:      res,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return

}
