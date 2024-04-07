package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAppUserLoginLog = "app_user_login_log"

type AppUserLoginLog struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增id" json:"id"` // 自增id
	UserID       string         `gorm:"column:user_id;not null" json:"user_id"`
	Status       string         `gorm:"column:status;not null" json:"status"`
	PromotionURL string         `gorm:"column:promotion_url;not null" json:"promotion_url"`
	ChannelID    string         `gorm:"column:channel_id;not null" json:"channel_id"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (*AppUserLoginLog) TableName() string {
	return TableNameAppUserLoginLog
}
