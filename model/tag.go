package model

type Tag struct {
	BaseModel
	Name string `gorm:"size:50;not null"` //名称
	ArticleNum int `gorm:"default:0;not null"` //数量
}
