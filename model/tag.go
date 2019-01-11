package model

import (
	"strings"
)

type Tag struct {
	BaseModel
	Name string `gorm:"size:50;not null"` //名称
	ArticleNum int `gorm:"default:0;not null"` //数量
}

func HandelTagAdd(tags string)  {
	//新增
	if tags == "" {
		return
	}
	//转换tags
	tags = strings.Replace(tags,"，", "," , -1)
	tagsSlice := strings.Split(tags, ",")
	tagsMap := getTags()
	for _, tag := range tagsSlice {
		isHas := false
		for oldTag, _ := range tagsMap {
			if oldTag == tag {
				isHas = true
			}
		}
		if isHas {
			//加1
			var tagInfo Tag
			DB.Where("name = ?", tag).First(&tagInfo)
			tagInfo.ArticleNum += 1
			DB.Save(&tagInfo)
		} else {
			//新增
			DB.Create(&Tag{Name:tag,ArticleNum:1})
		}
	}
}
func HandelTagDel(tags string) {
	tagsSlice := strings.Split(tags, ",")
	for _, tag := range tagsSlice {
		//减1
		var tagInfo Tag
		DB.Where("name = ?", tag).First(&tagInfo)
		if tagInfo.ArticleNum > 1 {
			tagInfo.ArticleNum -= 1
			DB.Save(&tagInfo)
		} else {
			DB.Delete(&tagInfo)
		}
	}
}
func HandelTagEdit(oldTags string, tags string) {
	HandelTagDel(oldTags)
	HandelTagAdd(tags)
}
func getTags() map[string]int{
	var tags []*Tag
	DB.Find(&tags)
	tagsMap := make(map[string]int)
	if len(tags) > 0 {
		for _, v := range tags {
			tagsMap[v.Name] = v.ArticleNum
		}
	}
	return tagsMap
}