package core

import (
	"fmt"
	"gin-vue-admin/server/initialize/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper { //读取配置信息
	v := viper.New()
	//读取配置文件
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	//处理读取err
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败! err : %w", err))
	}
	//处理解析 err
	err = v.Unmarshal(&global.GVA_CONFIG)
	if err != nil {
		panic(fmt.Errorf("配置解析失败! err : %w", err))
	}
	//配置热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e.Name, "配置文件已经修改")
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println("重新解析配置文件失效! err : ", err)
		}
	})
	
	return v
}
