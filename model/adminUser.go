package model

type AdminUser struct {
	BaseModel
	Name string `gorm:"size:50;not null"` //登录名
	Nickname string `gorm:"size:50;not null"` //用户名
	Pwd string `gorm:"size:350;not null"` //密码
}
