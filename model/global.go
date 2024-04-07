package model

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
