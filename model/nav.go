package model

type Nav struct {
	BaseModel
	Name string `gorm:"size:50;not null"` //名称
	Alias string `gorm:"size:50;not null"` //别名
	Pid int `gorm:"default:0;not null"` //父ID
	Url string `gorm:"size:255;not null;"` //链接
	Sort int `gorm:"default:0;not null"` //排序
}

func GetNavTree() []*Nav {
	var navs, _navs []*Nav
	DB.Order("sort desc").Find(&navs)
	for _, v := range navs {
		if v.Pid == 0 {
			_navs = append(_navs, v)
			for _, v1 := range navs{
				if v.ID == (uint)(v1.Pid) {
					_navs = append(_navs, v1)
				}
			}
		}
	}
	return _navs
}
