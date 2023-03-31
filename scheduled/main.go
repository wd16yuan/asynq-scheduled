package main

import (
	"padmin/scheduled/core"
	"padmin/scheduled/global"
	"padmin/scheduled/initialize"

	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper()       // 初始化Viper
	global.GVA_LOG = core.Zap()        // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG) // 替换zap全局logger

	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.Run()
}
