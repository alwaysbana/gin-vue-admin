package core

import (
	"fmt"
	"gin-vue-admin/server/initialize/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() {
	//拼接 dsn

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.GVA_CONFIG.Mysql.User,     // → root
		global.GVA_CONFIG.Mysql.Password, // → 11111111
		global.GVA_CONFIG.Mysql.Host,     // → 127.0.0.1
		global.GVA_CONFIG.Mysql.Port,     // → 3306
		global.GVA_CONFIG.Mysql.Db,       // → gin-vue-admin
	)

	//打开数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.GVA_LOG.Error("数据库连接失败", zap.Error(err))
		return // 或 panic
	}

	global.GVA_DB = db
	global.GVA_LOG.Info("数据库连接成功")
}
