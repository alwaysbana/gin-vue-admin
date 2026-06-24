package core

import (
	"gin-vue-admin/server/initialize/global"

	"go.uber.org/zap"
)

func Zap() {
	//判断模式
	if global.GVA_CONFIG.Server.Mode == "debug" {
		global.GVA_LOG, _ = zap.NewDevelopment()
		// global.GVA_LOG.Info("服务启动成功!",
		// zap.Int("port", 8080),
		// zap.String("mode", "debug")
		// )
		// global.GVA_LOG.Warn("磁盘空间不足",
		// zap.String("path", "/data"),
		// zap.Int64("available_mb", 128),)

	} else {
		global.GVA_LOG, _ = zap.NewProduction()
	}
}
