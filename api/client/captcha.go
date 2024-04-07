package client

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"time"
	"web2app/global"
	"web2app/model"
	"web2app/model/common/response"
	"web2app/utils"
)

var store = base64Captcha.DefaultMemStore

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /system/captcha [get]
func (baseApi *BaseApi) Captcha(c *gin.Context) {

	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)

	if id, b64s, _, err := cp.Generate(); err != nil {
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(model.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

func (baseApi *BaseApi) SendEmail(c *gin.Context) {
	var request model.SendEmailRequest
	if err := c.BindJSON(&request); err != nil {
		response.FailTokenDetailed(utils.Unauthoriz, "", err.Error(), c)
		return
	}
	code := utils.GenerateCode()
	global.Dc.Set(request.Email, code, time.Minute*10)
	body := fmt.Sprintf("`Your verification code is: %s`", code)
	sender := &utils.MailSender{}
	mail := utils.NewMailMessage("验证码", body, utils.HtmlType, nil)
	sender = sender.LoadConfig(utils.NewMailConfig(global.SMTP_SERVICE, global.EMAIL_PORT, global.EMAIL_USER,
		global.EMAIL_PASSWD, true))
	sender = sender.AddReceiver(request.Email).AddMail(mail).AddCc(global.EMAIL_USER)
	send, err := sender.Send()
	if err != nil {
		fmt.Printf("has error:%s", err.Error())
		return
	} else {
		fmt.Println(send)
	}
	if err != nil {
		response.FailWithMessage("fail", c)
	} else {
		response.OkWithMessage("success", c)
	}
}
