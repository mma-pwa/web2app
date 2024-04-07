package utils

const (
	TOKENERROR                     = 10001
	ERROR_TOKEN_Illegal            = 10002
	ERROR_TOKEN_BLACK              = 10003
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 10004
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10005
	ERROR_AUTH_TOKEN               = 10006
	ERROR_TOEKN_EVENT              = 10007
	LOGINERROR                     = 10008
	CAPTCHAERROR                   = 10009
	ERROR_TOKEN_OVERTIME           = 10010
	ERROR_PASSWORD                 = 10011
	TOKEN_MAL_FORMED               = 10012
	TOKEN_EXPIRED                  = 10013
	TOKEN_NotVALID_Yet             = 10014
	TOKEN_INVALID                  = 10015
	ERROR_PASSWORD_INCONSISTENCY   = 10016
	ERROR_USEREXIST                = 10017
	ERROR_CAPTCHA_NOTEXIST         = 10018
	ERROR_PASSWD_EXPIRE            = 11000
	ERROR_NOT_ENOUGH_ARGS          = 11001
	ERROR_ARGS                     = 11002
	ERROR_AUTH_ROLE_FAIL           = 11008
	ERROR_FIRST_LOGIN              = 12000
	ERROR_DATABASE_UPDATE          = 30000
	ERROR_DATABASE_CREATE          = 30001
	ERROR_DATABASE_FINDONE         = 30002
	ERROR_DATABASE_FINDMANY        = 30013
	ERROR_DATABASE_DELETEONE       = 30004
	ERROR_MFA_BIND_FAIL            = 30005
	ERROR_MFA_NO_PASS              = 30006
	PassowrdOrUserNameError        = 30002
	Unauthoriz                     = 402 //无权限
	ERROR                          = 7
	SUCCESS                        = 0
	INVALID_PARAMS                 = 400
	CREATE_FAIL                    = 401
	UnauthorizedIp                 = 30001
	ERROR_COUNT_DOCUMENTS          = 40000
	PassowordOrUserNameErrorMsg    = 31000
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "Internal server error",
	INVALID_PARAMS:                 "Bad Request",
	Unauthoriz:                     "无权限",
	ERROR_TOKEN_BLACK:              "token黑名单",
	ERROR_TOKEN_Illegal:            "token不合法",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	CREATE_FAIL:                    "",
	UnauthorizedIp:                 "未授权的ip",
	ERROR_FIRST_LOGIN:              "第一次登录",
	ERROR_PASSWD_EXPIRE:            "密码过期",
}
