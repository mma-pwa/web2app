package global

import (
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"time"
	"web2app/config"
)

const (
	ConfigEnv  = "GVA_CONFIG"
	ConfigFile = "config.yaml"
)

var (
	GVA_DB                  *gorm.DB
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Concurrency_Control = &singleflight.Group{}
	DEBUG                   bool //是否为调试模式
	Dc                      = cache.New(20*time.Minute, 20*time.Minute)
)

const (
	ICON_IMG = "1" //图标图片
	APP_IMG  = "2" //应用截图

	Landscape_screen = "1" //横屏
	Vertical_screen  = "0" //竖屏

	BIG_ICON   = "_512x512"
	SMALL_ICON = "_192x192"

	//EMAIL_USER   = "fenglei082@gmail.com"
	//EMAIL_PASSWD = "Doyo1028.."
	//SMTP_SERVICE = "smtp.google.com"
	//EMAIL_PORT   = 465
	//wei@guanwakeji.com
	EMAIL_USER   = "no-reply@quicka2b.com"
	EMAIL_PASSWD = "JRzXbM2wNpXC7pQVJfn1"
	SMTP_SERVICE = "smtpout.secureserver.net"
	EMAIL_PORT   = 465
)
