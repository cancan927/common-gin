package router

import (
	v1 "github.com/cancan927/common-gin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouters() {
	//注册user模块路由
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {

		//登陆是公共路由
		public := rgPublic.Group("")
		{
			public.POST("/register", v1.UserRegisterHandler)
			public.GET("/login", v1.UserLoginHandler)
		}

		//需要鉴权的user分组
		//user := rgAuth.Group("user")
		//{
		//	//user.GET("/:id", v1.UserInfoByIdHandler)
		//}

	})
}
