package model

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增id" json:"id"` // 自增id
	Url       string         `gorm:"column:url;not null" json:"url"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"` // 删除时间
}
