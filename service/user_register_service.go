package service

import (
	"github.com/cancan927/common-gin/global"
	"github.com/cancan927/common-gin/serializer"
	"github.com/cancan927/common-gin/util"
)

type UserRegisterService struct {
	NickName string `json:"nickname" binding:"required"  gorm:"column:nickname"`
	PassWord string `json:"password" binding:"required,len=8"  gorm:"column:password"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=120"`
	Sex      uint8  `json:"sex"`
	Email    string `json:"email" binding:"required,email"`
}

func (UserRegisterService) TableName() string {
	return "user"
}

func (service *UserRegisterService) Register() serializer.Response {
	//验证邮箱是否已经注册
	var cnt int
	//global.DB.Table("user").Select("count(*) ").Where("email = ?", service.Email).Scan(&cnt)
	global.DB.Raw("select count(*) from user where email = ?", service.Email).Scan(&cnt)
	if cnt > 0 {
		return serializer.Fail(serializer.REGISTED_EMAIL, "邮箱已被注册")
	} else {
		//密码加密
		password := service.PassWord
		encrypt, err := util.Encrypt(password)
		if err != nil {
			return serializer.Fail(40009, "密码加密出错")
		}
		service.PassWord = encrypt
		//创建用户
		res := global.DB.Create(service)
		if res.RowsAffected != 1 {
			return serializer.Fail(serializer.FAIL_REGISTER, "注册失败")
		}
		return serializer.Success("注册成功", service.NickName)
	}

}
