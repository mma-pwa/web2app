package client

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
	"web2app/global"
	"web2app/model"
	"web2app/utils"
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) CreateUser(user *model.User) error {

	if !errors.Is(global.GVA_DB.Where("email = ?", user.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册")
	}
	if !utils.ValidateEmail(user.Email) {
		return errors.New("valid email address error")
	}
	if user.Password != user.ConfirmPasswd {
		return errors.New("the Two Passwords Entered Are Inconsistent")
	}
	user.ID = utils.MD5V([]byte(uuid.New().String()))
	user.Password, _ = utils.PasswordHash(user.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Money = 1.0
	err := global.GVA_DB.Create(user).Error
	if err != nil {
		return err
	}
	return err
}

func (s *UsersService) DeleteUser(user *model.User) error {
	err := global.GVA_DB.Delete(&user).Error
	return err
}
func (s *UsersService) UpdateUser(user *model.User) (err error, userInfo *model.User) {
	err = global.GVA_DB.Updates(&user).Error
	return err, user
}

func (s *UsersService) GetUser(id string) (err error, users *model.User) {
	var u model.User
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func (s *UsersService) Login(u *model.User) (err error, userInfo *model.User) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}
	var user model.User
	err = global.GVA_DB.Where("email = ?", u.Email).First(&user).Error
	if !utils.PasswordVerify(u.Password, user.Password) {
		return errors.New("passwd error"), nil
	}
	return err, &user
}
func (s *UsersService) ChangePassword(u *model.User, newPassword string) (err error, userInter *model.User) {
	var user model.User
	u.Password, _ = utils.PasswordHash(user.Password)
	newPassword, _ = utils.PasswordHash(newPassword)
	err = global.GVA_DB.Where("email = ? AND password = ?", u.Email, u.Password).First(&user).Update("password", newPassword).Error
	return err, u
}

func (s *UsersService) TokenNext(user *model.User) (string, error) {
	expiresTime := global.GVA_CONFIG.JWT.ExpiresTime

	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(model.BaseClaims{
		Email: user.Email,
		ID:    user.ID,
	}, expiresTime)
	token, err := j.CreateToken(claims)

	token = utils.TokenCrypt(token, time.Now().Format(time.DateOnly))
	//fmt.Println("************加密后", token)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		//response.FailWithMessage("获取token失败", c)
		return "获取Token失败", err
	}
	return token, nil
}
