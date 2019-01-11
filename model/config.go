package model

type Config struct {
	BaseModel
	Title string `gorm:"size:50;not null;default:''"` //标题
	Name string `gorm:"size:50;not null;default:''"` //变量名
	Content string `gorm:"size:50;not null;default:''"` //变量值
	Sort int `gorm:"default:0;not null"` //排序
	Tips string `gorm:"size:255;not null;default:''"` //描述
	FieldType string `gorm:"size:50;not null;default:''"` //字段类型
	FieldValue string `gorm:"size:255;not null;default:''"` //类型值
}

func GetAllConfig() map[string]string {
	var configs []*Config
	DB.Select("name,content").Order("sort desc").Find(&configs)
	m := make(map[string]string)
	for _, config := range configs {
		m[config.Name] = config.Content
	}
	return m
}
