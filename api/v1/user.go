package v1

import (
	"github.com/cancan927/common-gin/serializer"
	"github.com/cancan927/common-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(c *gin.Context) {
	var s service.UserRegisterService
	if err := c.ShouldBindJSON(&s); err != nil { //这里一般是参数校验失败
		c.JSON(200, serializer.Response{
			Code: 40004,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	} else {
		res := s.Register()
		c.JSON(200, res)
		return
	}
}

func UserLoginHandler(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Code: 40005,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	} else {
		res := s.Login(c)
		c.JSON(http.StatusOK, res)
		return
	}
}

//func UserInfoByIdHandler(c *gin.Context) {
//	//id := c.Query("id")
//	var s service.UserInfoService
//	id := c.Param("id")
//	u, err := strconv.Atoi(id)
//	if err != nil {
//		global.Logger.Errorf(err.Error())
//	}
//	s.Id = u
//	res := s.InfoById
//	c.JSON(http.StatusOK, res)
//	return
//}
