package main

import (
	"go.uber.org/zap"
	"web2app/global"
	"web2app/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	global.GVA_VP = initialize.Viper()
	global.GVA_LOG = initialize.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm()
	initialize.RunWindowsServer()
}
