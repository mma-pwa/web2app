package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"web2app/global"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initRouters()

	address := fmt.Sprintf("%s", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	
	默认自动化文档地址:http://%s/swagger/index.html
	默认前端文件运行地址:http://%s
`, address, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
