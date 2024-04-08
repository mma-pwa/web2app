package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
	"web2app/global"
	"web2app/model"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims model.BaseClaims, ExpiresTime int64) model.CustomClaims {
	claims := model.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + ExpiresTime,
			Issuer:    global.GVA_CONFIG.JWT.Issuer,
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) CreateTokenByOldToken(oldToken string, claims model.CustomClaims) (string, error) {
	v, err, _ := global.GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error, int) {
	//fmt.Println("**************解密前", tokenString)
	tokenString = Tokendecrypt(tokenString, time.Now().Format(time.DateOnly))
	//fmt.Println("--------------解密后", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed, TOKEN_MAL_FORMED
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired, TOKEN_EXPIRED
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet, TOKEN_NotVALID_Yet
			} else {
				fmt.Println("----------------", err)
				return nil, TokenInvalid, TOKEN_INVALID
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil, SUCCESS
		}
		return nil, TokenInvalid, TOKEN_INVALID

	} else {
		return nil, TokenInvalid, TOKEN_INVALID
	}
}

func TokenCrypt(token, now string) string {
	tokenList := []byte(token)
	hashstr := []byte(Md5Crypt(now))[3:11]
	fmt.Println(hashstr)
	newToken := string(tokenList[0:18]) + string(hashstr) + string(tokenList[18:])
	return newToken
}
func Tokendecrypt(token, now string) string {
	hashstr := []byte(Md5Crypt(now))[3:11]
	newToken := strings.ReplaceAll(token, string(hashstr), "")
	return newToken
}

func GetUserToken(tokenurl, userName, passWord string) string {
	type UserInfo struct {
		UserName string `json:"userName"` // 用户名
		PassWord string `json:"passWord"` // 密码
	}
	t := &http.Transport{
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: 60 * time.Second,
		DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true, //添加这一行跳过验证
		},
	}
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	url := fmt.Sprintf("%s/api/system/token", tokenurl)
	user := UserInfo{UserName: userName, PassWord: passWord}
	jsonBytes, _ := json.Marshal(user)
	rst, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	rst.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout:   60 * time.Second,
		Transport: t,
	}
	res, err := client.Do(rst)
	if err != nil {
		return ""
	}
	body, _ := ioutil.ReadAll(res.Body)
	var mapData map[string]interface{}
	if err := json.Unmarshal(body, &mapData); err != nil {
		global.GVA_LOG.Error("解析token数据失败", zap.Error(err))
	}
	token := ""
	if v, ok := mapData["data"]; ok {
		if ok {
			if t, ok := v.(map[string]interface{})["token"]; ok {
				token = t.(string)
			}
		}
	}

	if res != nil {
		defer func() {
			if err := res.Body.Close(); err != nil {
				global.GVA_LOG.Error("关闭body失败", zap.Error(err))
			}
		}()
	}
	return token
}
