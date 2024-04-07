package request

type Login struct {
	Email     string `json:"email"`     // 用户名
	PassWord  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
	UserId    int64  `json:"userId"`
	MfaCode   string `json:"mfaCode"`
}
