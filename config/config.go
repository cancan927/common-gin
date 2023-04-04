package config

import (
	"os"

	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	//读取项目根目录，即main.go所在目录
	workDir, err := os.Getwd()
	if err != nil {
		panic("读取工作目录错误：" + err.Error())
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取错误：" + err.Error())
	}

}
