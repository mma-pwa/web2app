package initialize

import (
	"fmt"
	"io"

	//"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"os"
	"web2app/global"
	"web2app/middleware"
	"web2app/router"
)

// InitRouters 初始化总路由
func initRouters() *gin.Engine {
	logfile, err := os.Create("/var/log/http_gin.log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
	Router := gin.Default()
	//if global.GVA_CONFIG.Runmod.RunEnv == gin.DebugMode {
	//	ginpprof.Wrap(Router)
	//}
	//if "debug" == gin.DebugMode {
	//	ginpprof.Wrap(Router)
	//}
	//Router.Use(middleware.Recovery())
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware logger")
	global.GVA_LOG.Info("use middleware cors")
	Router.Static("/img", global.GVA_CONFIG.Local.UploadPath)
	client := router.RouterGroupApp.Client
	Router.Static("/static", "./dist/static/")
	Router.LoadHTMLGlob("./dist/*.html")
	Router.GET("/", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(200, "index.html", nil)
	})
	PublicGroup := Router.Group("api")
	{
		// 健康监测
		PublicGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		client.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		client.InitAppRouter(PrivateGroup)
		client.InitUsersRouter(PrivateGroup)
		client.InitUpgradeRouter(PrivateGroup)
		client.InitAppCustomURLRouter(PrivateGroup)
		client.InitAppUserLoginLogRouter(PrivateGroup)
		client.InitAppUserReportDayRouter(PrivateGroup)
		client.InitDomainApiRoute(PrivateGroup)

	}
	return Router
}
