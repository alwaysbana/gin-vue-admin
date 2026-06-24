package global

import (
	"gin-vue-admin/server/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GVA_CONFIG config.Config//配置信息
var GVA_LOG *zap.Logger//zap 日志配置
var GVA_DB *gorm.DB//数据库信息
