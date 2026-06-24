package main

import (
	"fmt"
	"gin-vue-admin/server/core"
	"gin-vue-admin/server/initialize/global"
)

func main() {

	core.Viper() //获取配置
	global.GVA_LOG.Info(fmt.Sprintf("配置加载成功: %+v", global.GVA_CONFIG))
	core.Zap()  //日志配置
	core.Gorm() //数据库链接
	defer global.GVA_LOG.Sync()
	global.GVA_LOG.Info("测试日志")
}
