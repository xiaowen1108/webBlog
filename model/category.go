package model

type Category struct {
	BaseModel
	Name string `gorm:"size:50;not null"` //分类名
	SeoTitle string `gorm:"size:100;not null"` //SEO标题
	SeoKey string `gorm:"size:100;not null"` //SEO key
	SeoDesc string `gorm:"size:255;not null"` //SEO desc
	Sort int `gorm:"default:0;not null"` //排序
	Pid int `gorm:"default:0;not null"` //父ID
}

