package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

type User struct {
	ID            string         `gorm:"column:id;primaryKey;comment:Uuid md5 16位" json:"id"` // Uuid md5 16位
	Email         string         `gorm:"column:email;not null;comment:邮箱，账户" json:"email"`    // 邮箱，账户
	Password      string         `gorm:"column:password;not null;comment:密码" json:"password"` // 密码
	Money         float64        `gorm:"column:money;not null;comment:余额" json:"money"`       // 余额
	Rate          float64        `gorm:"column:rate;not null;default:0.080" json:"rate"`
	Utype         int            `json:"utype" gorm:"column:utype;not null;default:0" `
	CreatedAt     time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Code          string         `json:"code" gorm:"-"`
	ConfirmPasswd string         `json:"confirmPasswd" gorm:"-"`
}
type SysCaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}
type SysUserRes struct {
	ID    string  `gorm:"column:id;primaryKey;comment:Uuid md5 16位" json:"id"` // Uuid md5 16位
	Email string  `gorm:"column:email;not null;comment:邮箱，账户" json:"email"`    // 邮箱，账户
	Money float64 `gorm:"column:money;not null;comment:余额" json:"money"`       // 余额
}

type SendEmailRequest struct {
	Email string `json:"email"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
