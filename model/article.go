package model

type Article struct {
	BaseModel
	Cid int  `gorm:"not null"` //分类
	Title string `gorm:"size:100;not null"` //标题
	Description string `gorm:"size:255;not null"` //描述
	Tags string `gorm:"size:100;not null"` //标签
	Cover string `gorm:"size:255;not null"` //封面图
	IsRecom bool `gorm:"default:0;not null"` //是否推荐
	LikeNum int `gorm:"default:0;not null"` //喜欢数量
	Sort int `gorm:"default:0;not null"` //排序
	Hits int `gorm:"default:0;not null"` //点击量
	Content string `gorm:"not null;type:text"`//文章内容
}
