package main

import (
	"padmin/worker/core"

	"padmin/worker/global"

	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper()       // 初始化Viper
	global.GVA_LOG = core.Zap()        // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG) // 替换zap全局logger
	core.Run()
}
