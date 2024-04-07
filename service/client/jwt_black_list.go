package client

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"web2app/global"
	"web2app/model"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	jwtList.CreateTime = time.Now()
	return global.GVA_DB.Create(jwtList).Error
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {

	err := global.GVA_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}
