package cmd

import (
	"github.com/cancan927/common-gin/config"
	"github.com/cancan927/common-gin/global"
)

// Start 开启一个web服务
func Start() {
	//====================================================================
	// = 初始化配置
	config.InitConfig()

	//====================================================================
	// = 初始化日志
	//TODO

	//====================================================================
	// = 初始化数据库
	db, err := config.InitDB()
	if err != nil {
		panic(err) //初始化数据库失败就直接panic
	}
	global.DB = db

	//====================================================================
	// = 初始化redis
	rdClient, rdErr := config.InitRedis()
	global.RedisClient = rdClient
	if rdErr != nil {
		global.Logger.Error(rdErr)
	}

	//====================================================================
	// = 初始化路由

}
