package model

type Link struct {
	BaseModel
	Name string `gorm:"size:255;not null;"` //名称
	Title string `gorm:"size:255;not null;"` //标题
	Url string `gorm:"size:255;not null;"` //链接
	Sort int `gorm:"default:0;not null"` //排序
}
