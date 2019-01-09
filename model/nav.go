package model

type Nav struct {
	BaseModel
	Name string `gorm:"size:50;not null"` //名称
	Alias string `gorm:"size:50;not null"` //别名
	Pid int `gorm:"default:0;not null"` //父ID
	Url string `gorm:"size:255;not null;"` //链接
	Sort int `gorm:"default:0;not null"` //排序
}
