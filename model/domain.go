package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameDomain = "domain"

type Domain struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增id" json:"id"` // 自增id
	Url       string         `gorm:"column:url;not null" json:"url"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;type:datetime" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;type:datetime" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"` // 删除时间
}

func (*Domain) TableName() string {
	return TableNameDomain
}
