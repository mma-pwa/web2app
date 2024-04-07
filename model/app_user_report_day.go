package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAppUserReportDay = "app_user_report_day"

// AppUserReportDay mapped from table <app_user_report_day>
type AppUserReportDay struct {
	ID             string         `gorm:"column:id;primaryKey;comment:自增id" json:"id"` // 自增id
	ReportDate     time.Time      `gorm:"column:report_date;not null" json:"report_date"`
	AppID          string         `gorm:"column:app_id;not null" json:"app_id"`
	AppName        string         `gorm:"column:app_name;not null" json:"app_name"`
	AppCustomURL   string         `gorm:"column:app_custom_url;not null" json:"app_custom_url"`
	CostMoney      float64        `gorm:"column:cost_money;not null" json:"cost_money"`
	InstallUsers   int64          `gorm:"column:install_users;not null;comment:安装用户数" json:"install_users"`          // 安装用户数
	OpenIndexUsers int64          `gorm:"column:open_index_users;not null;comment:打开推广页用户数" json:"open_index_users"` // 打开推广页用户数
	OpenAppUsers   int64          `gorm:"column:open_app_users;not null;comment:打开应用用户数" json:"open_app_users"`      // 打开应用用户数
	CreatedAt      time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (*AppUserReportDay) TableName() string {
	return TableNameAppUserReportDay
}
