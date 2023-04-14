package model

import "gorm.io/gorm"

type User struct {
	Id       int
	NickName string `json:"nickname"  gorm:"column:nickname"`
	PassWord string `json:"password"   gorm:"column:password"`
	Age      uint32 `json:"age"`
	Sex      uint32 `json:"sex"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(orm *gorm.DB) {

}
