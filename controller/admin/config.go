package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webBlog/helper"
	"webBlog/model"
	"strings"
	"bytes"
	"fmt"
	"strconv"
)

type Config struct {

}

func (t *Config) Index (c *gin.Context){
	var configs []*model.Config
	model.DB.Order("sort desc").Find(&configs)
	for k, v := range configs {
		switch v.FieldType {
			case "input":
				configs[k].FieldValue = "<input type='text' class='lg' name='content[]' value='" + v.Content +"'>"
				break;
			case "textarea":
				configs[k].FieldValue = "<textarea type='text' class='lg' name='content[]'>"+v.Content+"</textarea>"
			case "radio":
				//1|开启,0|关闭
				arr :=  strings.Split(v.FieldValue, ",")
				var str bytes.Buffer
				for _, n := range arr {
					r := strings.Split(n, "|")
					var ck string
					if v.Content == r[0] {
						ck = "checked"
					}
					str.WriteString(r[1])
					str.WriteString("<input type='radio' name='content[]' value='")
					str.WriteString(r[0])
					str.WriteString("'")
					str.WriteString(ck)
					str.WriteString(">")
					str.WriteString(" ")
				}
				configs[k].FieldValue = str.String()
		}
	}
	c.HTML(http.StatusOK, "admin/config/index.html",gin.H{
		"configs":configs,
	})
}
type ConfigData struct {
	Title string `form:"title" binding:"required"`
	Name string `form:"name" binding:"required"`
	FieldType string `form:"field_type"`
	FieldValue string `form:"field_value"`
	Sort int `form:"sort"`
	Tips string `form:"tips"`
}
func (t *Config) Add (c *gin.Context){
	if helper.IsGet(c) {
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/config/add.html",gin.H{
			"errorMsg":errorMsg,
		})
	} else if helper.IsPost(c) {
		var configData ConfigData
		if err := c.ShouldBind(&configData); err == nil{
			config := model.Config{
				Name:configData.Name,
				Title:configData.Title,
				FieldType:configData.FieldType,
				FieldValue:configData.FieldValue,
				Tips:configData.Tips,
				Sort:configData.Sort,
			}
			err := model.DB.Create(&config).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "添加成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "添加失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/config/add")
	}
}

func (t *Config) Edit (c *gin.Context){
	id := c.Param("id")
	var config model.Config
	err := model.DB.Where("id = ?", id).First(&config).Error
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/config/index")
		return
	}
	if helper.IsGet(c) {
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/config/edit.html",gin.H{
			"errorMsg":errorMsg,
			"config":config,
		})
	} else if helper.IsPost(c) {
		var configData ConfigData
		if err := c.ShouldBind(&configData); err == nil{
			config.Name = configData.Name
			config.Title = configData.Title
			config.FieldType = configData.FieldType
			config.FieldValue = configData.FieldValue
			config.Tips = configData.Tips
			config.Sort = configData.Sort
			err := model.DB.Save(&config).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "编辑成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "编辑失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/config/edit/"+id)
	}
}

func (t *Config) Del (c *gin.Context){
	id := c.Param("id")
	var config model.Config
	err := model.DB.Where("id = ?", id).First(&config).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该配置，删除失败")
		return
	}
	model.DB.Delete(&config)
	helper.ReturnJson(c, 1, "删除成功")
	return
}

func (t *Config) ChangeOrder (c *gin.Context){
	id := c.PostForm("id")
	sort := c.PostForm("sort")
	if id =="" || sort == "" {
		helper.ReturnJson(c, 0, "参数错误")
		return
	}
	var config model.Config
	err := model.DB.Where("id = ?", id).First(&config).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该配置项，设置失败")
		return
	}
	config.Sort, err = strconv.Atoi(sort)
	if err != nil {
		helper.ReturnJson(c, 0, "非法排序值")
		return
	}
	model.DB.Save(&config)
	helper.ReturnJson(c, 1, "修改成功")
}

type ContentData struct {
	Ids []string   `form:"id[]" binding:"required"`
	Content []string `form:"content[]" binding:"required"`
}
func (t *Config) ChangeContent (c *gin.Context){
	var contentData ContentData
	if err := c.ShouldBind(&contentData); err == nil{
		fmt.Println(contentData)
		for k,id := range contentData.Ids {
			var config model.Config
			err := model.DB.Where("id = ?", id).First(&config).Error
			if err == nil {
				content := contentData.Content[k]
				if content != "" {
					model.DB.Model(&config).Update("content", content)
				}
			}
		}
		helper.SetFlash(c, "errorMsg", "保存成功 ！")
	} else {
		helper.SetFlash(c, "errorMsg", "缺少参数 ！")
	}
	c.Redirect(http.StatusFound, "/admin/config/index")
}