package router

import (
	"context"
	"fmt"
	"github.com/cancan927/common-gin/global"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// IFnRegisterRoute rgGroup：公共路由组  rgAuth：需要鉴权的路由组
type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	//各个基础平台注册路由的函数
	gfnRoutes []IFnRegisterRoute
)

// RegisterRoute 注册路由方法
func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

// InitBaseRouters 初始化基础模块的路由
func InitBaseRouters() {
	//初始化user路由
	//InitUserRouters()
}

// InitRouter 初始化系统路由
func InitRouter() {

	// 创建监听信号的context
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// 初始化gin，并注册相关路由
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	//初始化基础路由
	InitBaseRouters()

	//遍历路由方法，注册其中的路由
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	//读取配置文件，获取启动端口
	srvName := viper.GetString("server.name")
	srvPort := viper.GetString("server.port")
	if srvPort == "" {
		srvPort = "8080"
	}

	serv := &http.Server{
		Addr:    fmt.Sprintf(":%s", srvPort),
		Handler: r,
	}

	go func() {
		global.Logger.Infof("Start Listen:%s", srvPort)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start %s Server Error:%s", srvName, err.Error()))
			return
		}
		//fmt.Println(fmt.Sprintf("Start %s Server Listenning :%s", srvName, srvPort))
	}()

	<-ctx.Done()
	cancel()

	timeoutCtx, shutdown := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutdown()

	if err := serv.Shutdown(timeoutCtx); err != nil {
		// 记录日志
		global.Logger.Errorf("Stop %s Server Error:%s", srvName, err.Error())
		return
	}

	fmt.Printf("Stop %s Server Success", srvName)

}
