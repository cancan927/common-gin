package service

import (
	"github.com/cancan927/common-gin/global"
	"github.com/cancan927/common-gin/model"
	"github.com/cancan927/common-gin/serializer"
	"github.com/cancan927/common-gin/util"
	"github.com/gin-gonic/gin"
)

type UserLoginService struct {
	Email      string `json:"email" binding:"required,email"`
	PassWord   string `json:"password" binding:"required,len=8"  gorm:"column:password"`
	RePassWord string `json:"repassword" binding:"required,len=8"  gorm:"column:password"`
}

func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	//1 检查repassword是否和password相同
	if service.PassWord != service.RePassWord {
		return serializer.Fail(serializer.DIFF_REPASSWORD, "两次密码不一致")
	}
	//2 查询是否有这个用户
	user := &model.User{}
	result := global.DB.Where("email = ?", service.Email).Find(&user)
	if result.RowsAffected == 0 {
		return serializer.Fail(serializer.USER_NOT_EXSIST, "用户不存在")
	}
	//3 验证密码是否正确
	if !util.ValidatePassword(user.PassWord, service.PassWord) {
		return serializer.Fail(serializer.INCORRECT_PASSWD, "密码不正确")
	}
	//密码正确响应时返回token
	token, err := util.GenerateToken(user.Id, user.NickName)
	if err != nil {
		global.Logger.Errorf("生成token失败：%s", err)
	}
	c.Header("Authorization", token)
	return serializer.Success("登录成功", user.NickName)

}
