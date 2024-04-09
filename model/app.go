package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameApp = "app"

// App 用户表
type App struct {
	ID                string         `gorm:"column:id;primaryKey;comment:Uuid md5 16位" json:"id"` // Uuid md5 16位
	UserID            string         `gorm:"column:user_id;not null" json:"user_id"`
	Name              string         `gorm:"column:name;not null;comment:appname" json:"name"`                               // appname
	ShortName         string         `gorm:"column:short_name;not null;comment:短名称" json:"short_name"`                       // 短名称
	Description       string         `gorm:"column:description;not null;comment:描述" json:"description"`                      // 描述
	Icons             string         `gorm:"column:icons;not null;comment:Icon" json:"icons"`                                // Icon
	AppURL            string         `gorm:"column:app_url;not null;comment:App实际地址" json:"app_url"`                         // App实际地址
	AppDevName        string         `gorm:"column:app_dev_name;not null;comment:开发者名称" json:"app_dev_name"`                 // 开发者名称
	AppRateNum        float64        `gorm:"column:app_rate_num;not null;comment:评分" json:"app_rate_num"`                    // 评分
	AppRateCount      float64        `gorm:"column:app_rate_count;not null;comment:评分人数" json:"app_rate_count"`              // 评分人数
	AppInstallCount   int64          `gorm:"column:app_install_count;not null;comment:安装人数" json:"app_install_count"`        // 安装人数
	AppSafeAge        int64          `gorm:"column:app_safe_age;not null;comment:适配年龄" json:"app_safe_age"`                  // 适配年龄
	AppScreen         string         `gorm:"column:app_screen;not null;comment:五图" json:"app_screen"`                        // 五图
	AppDescription    string         `gorm:"column:app_description;not null;comment:应用介绍" json:"app_description"`            // 应用介绍
	Status            *int           `gorm:"column:status;not null;comment:状态，0:下架，1:在架" json:"status"`                      // 状态，0:下架，1:在架
	ScreenOrientation *int           `gorm:"column:screen_orientation;not null;comment:0:竖屏，1:横屏" json:"screen_orientation"` // 0:竖屏，1:横屏
	PromotionURL      string         `gorm:"column:promotion_url;not null;comment:推广链接" json:"promotion_url"`                // 推广链接
	InstallTemplate   *int           `gorm:"column:install_template;not null;comment:安装模板" json:"install_template"`          //安装模板
	Iframe            *int           `gorm:"column:iframe;not null;comment:安装模板" json:"iframe""`                             //在iframe打开
	IosJump           *int           `gorm:"column:ios_jump;not null;comment:是否配置iOS跳转地址" json:"ios_jump""`                  //ios是否调整
	AnySiteInstall    *int           `gorm:"column:any_site_install;not null;comment:点击任意位置安装应用" json:"any_site_install"`    //点击任意位置安装应用:0否1是
	SiteModify        *int           `gorm:"column:site_modify;not null;comment:网站地址是否允许修改" json:"site_modify"`              //网站地址是否允许修改:0否1是
	DataBuriedPoint   string         `gorm:"column:data_buried_point;not null;comment:数据埋点" json:"data_buried_point"`        //网站地址是否允许修改:0否1是
	BuriedPointId     string         `gorm:"column:buried_point_id;not null;comment:埋点ID" json:"buried_point_id"`            //数据埋点
	Remarks           string         `gorm:"column:remarks;comment:备注" json:"remarks""`                                      //埋点Id
	ConsumeUstdCount  int            `json:"consume_ustd_count" gorm:"-"`
	IconImg           []ImgInfo      `json:"icon_img" gorm:"-"`
	ScreenImg         []ImgInfo      `json:"screen_img" gorm:"-"`
	CreatedAt         time.Time      `gorm:"column:created_at;not null;type:datetime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;not null;type:datetime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"` // 删除时间
}
type ImgInfo struct {
	Url     string `json:"url"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	ImgType string `json:"img_type"`
}

type AppRes struct {
	Id   string `gorm:"column:id;not null" json:"id"`
	Name string `gorm:"column:name;not null;comment:appname" json:"name"`
}

// TableName App's table name
func (*App) TableName() string {
	return TableNameApp
}
