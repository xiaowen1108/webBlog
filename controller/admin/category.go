package admin

import (
	"github.com/gin-gonic/gin"
	"webBlog/helper"
	"net/http"
	"webBlog/model"
	"strconv"
)

type Category struct {
}

func (t *Category) Index (c *gin.Context){
	categorys := model.GetTree()
	c.HTML(http.StatusOK, "admin/category/index.html",gin.H{
		"categorys":categorys,
	})
}
type CategoryData struct {
	Name string `form:"name" binding:"required"`
	Pid int `form:"pid"`
	SeoTitle string `form:"seo_title"`
	SeoKey string `form:"seo_key"`
	SeoDesc string `form:"seo_desc"`
	Sort int `form:"sort"`
}
func (t *Category) Add (c *gin.Context){
	if helper.IsGet(c) {
		//获取所有的顶级分类
		var categorys []*model.Category
		model.DB.Where("pid = ?", "0").Order("sort desc").Find(&categorys)
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/category/add.html",gin.H{
			"errorMsg":errorMsg,
			"categorys":categorys,
		})
	} else if helper.IsPost(c) {
		var categoryData CategoryData
		if err := c.ShouldBind(&categoryData); err == nil{
			category := model.Category{
				Name:categoryData.Name,
				SeoTitle:categoryData.SeoTitle,
				SeoKey:categoryData.SeoKey,
				SeoDesc:categoryData.SeoDesc,
				Sort:categoryData.Sort,
				Pid:categoryData.Pid,
			}
			err := model.DB.Create(&category).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "添加成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "添加失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/category/add")
	}
}

func (t *Category) Edit (c *gin.Context){
	//获取当前
	id := c.Param("id")
	var category model.Category
	err := model.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/category/index")
		return
	}
	if helper.IsGet(c) {
		//获取所有的顶级分类
		var categorys []*model.Category
		model.DB.Where("pid = ?", "0").Order("sort desc").Find(&categorys)
		errorMsg := helper.GetFlash(c, "errorMsg")

		c.HTML(http.StatusOK, "admin/category/edit.html",gin.H{
			"errorMsg":errorMsg,
			"categorys":categorys,
			"category":category,
		})
	} else if helper.IsPost(c) {
		var categoryData CategoryData
		if err := c.ShouldBind(&categoryData); err == nil{
			category.Name = categoryData.Name
			category.SeoTitle = categoryData.SeoTitle
			category.SeoKey = categoryData.SeoKey
			category.SeoDesc = categoryData.SeoDesc
			category.Sort = categoryData.Sort
			category.Pid = categoryData.Pid
			err := model.DB.Save(&category).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "修改成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "修改失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/category/edit/"+id)
	}
}

func (t *Category) Del (c *gin.Context){
	id := c.Param("id")
	var category model.Category
	err := model.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该分类，删除失败")
		return
	}
	model.DB.Delete(&category)
	model.DB.Where("pid = ?", id).Delete(&model.Category{})
	helper.ReturnJson(c, 1, "删除成功")
	return
}

func (t *Category) ChangeOrder (c *gin.Context){
	id := c.PostForm("id")
	sort := c.PostForm("sort")
	if id =="" || sort == "" {
		helper.ReturnJson(c, 0, "参数错误")
		return
	}
	var category model.Category
	err := model.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该分类，设置失败")
		return
	}
	category.Sort, err = strconv.Atoi(sort)
	if err != nil {
		helper.ReturnJson(c, 0, "非法排序值")
		return
	}
	model.DB.Save(&category)
	helper.ReturnJson(c, 1, "修改成功")
}