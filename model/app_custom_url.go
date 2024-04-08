package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameAppCustomURL = "app_custom_url"

type AppCustomURL struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增id" json:"id"` // 自增id
	AppID        string         `gorm:"column:app_id;not null" json:"app_id"`
	DomainPrefix string         `json:"domain_prefix" gorm:"column:domain_prefix;not null"`
	CustomURL    string         `gorm:"column:custom_url;not null" json:"custom_url"`
	Status       int32          `gorm:"column:status;not null;comment:Url当前状态，0为正常，1为关闭" json:"status"` // Url当前状态，0为正常，1为关闭
	CreatedAt    time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type AppCustomURLRes struct {
	CreatedAt    time.Time      `gorm:"column:created_at;not null;type:datetime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;type:datetime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                            // 删除时间
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增id" json:"id"` // 自增id
	AppID        string         `gorm:"column:app_id;not null" json:"app_id"`
	CustomURL    string         `gorm:"column:custom_url;not null" json:"custom_url"`
	Status       int32          `gorm:"column:status;not null;comment:Url当前状态，0为正常，1为关闭" json:"status"` // Url当前状态，0为正常，1为关闭
	Name         string         `gorm:"name" json:"name"`
	DomainPrefix string         `json:"domain_prefix" gorm:"column:domain_prefix;not null"`
}

func (*AppCustomURL) TableName() string {
	return TableNameAppCustomURL
}
