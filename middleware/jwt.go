package middleware

import (
	"github.com/gin-gonic/gin"
	"web2app/model"
	"web2app/model/common/response"
	"web2app/service"
	"web2app/utils"
)

var (
	jwtService = service.JwtService
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("xsx-token")
		if token == "" {
			response.FailTokenDetailed(utils.ERROR_TOKEN_Illegal, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		//if jwtService.IsBlacklist(token) {
		//	response.FailTokenDetailed(utils.ERROR_TOKEN_BLACK, gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		//	c.Abort()
		//	return
		//}
		claims, done := TokenVerify(c, token)
		if done {
			return
		}
		c.Set("userId", claims.ID)
		c.Set("claims", claims)
		c.Next()
	}
}

func TokenVerify(c *gin.Context, token string) (*model.CustomClaims, bool) {
	j := utils.NewJWT()
	// parseToken 解析token包含的信息
	claims, err, _ := j.ParseToken(token)
	if err != nil {

		if err == utils.TokenExpired {
			response.FailTokenDetailed(utils.TOKENERROR, gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return nil, true
		}
		response.FailTokenDetailed(utils.ERROR_TOEKN_EVENT, gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return nil, true
	}
	return claims, false
}
