package admin

import (
	"github.com/gin-gonic/gin"
	"webBlog/model"
	"net/http"
	"webBlog/helper"
	"strconv"
)

type Link struct {
}

func (t *Link) Index (c *gin.Context){
	var links  []*model.Link
	model.DB.Order("sort desc").Find(&links)
	c.HTML(http.StatusOK, "admin/link/index.html",gin.H{
		"links":links,
	})
}
type LinkData struct {
	Name string `form:"name" binding:"required"`
	Title string `form:"title" binding:"required"`
	Url string `form:"url" binding:"required"`
	Sort int `form:"sort"`
}
func (t *Link) Add (c *gin.Context){
	if helper.IsGet(c) {
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/link/add.html",gin.H{
			"errorMsg":errorMsg,
		})
	} else if helper.IsPost(c) {
		var linkData LinkData
		if err := c.ShouldBind(&linkData); err == nil{
			link := model.Link{
				Name:linkData.Name,
				Title:linkData.Title,
				Url:linkData.Url,
				Sort:linkData.Sort,
			}
			err := model.DB.Create(&link).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "添加成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "添加失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/links/add")
	}
}

func (t *Link) Edit (c *gin.Context){
	//获取当前
	id := c.Param("id")
	var link model.Link
	err := model.DB.Where("id = ?", id).First(&link).Error
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/link/index")
		return
	}
	if helper.IsGet(c) {
		//获取所有的顶级分类
		errorMsg := helper.GetFlash(c, "errorMsg")

		c.HTML(http.StatusOK, "admin/link/edit.html",gin.H{
			"errorMsg":errorMsg,
			"link":link,
		})
	} else if helper.IsPost(c) {
		var linkData LinkData
		if err := c.ShouldBind(&linkData); err == nil{
			link.Name = linkData.Name
			link.Title = linkData.Title
			link.Url = linkData.Url
			link.Sort = linkData.Sort
			err := model.DB.Save(&link).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "修改成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "修改失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/links/edit/"+id)
	}
}

func (t *Link) Del (c *gin.Context){
	id := c.Param("id")
	var link model.Link
	err := model.DB.Where("id = ?", id).First(&link).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该友链，删除失败")
		return
	}
	model.DB.Delete(&link)
	model.DB.Where("pid = ?", id).Delete(&model.Link{})
	helper.ReturnJson(c, 1, "删除成功")
	return
}

func (t *Link) ChangeOrder (c *gin.Context){
	id := c.PostForm("id")
	sort := c.PostForm("sort")
	if id =="" || sort == "" {
		helper.ReturnJson(c, 0, "参数错误")
		return
	}
	var link model.Link
	err := model.DB.Where("id = ?", id).First(&link).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该友链，设置失败")
		return
	}
	link.Sort, err = strconv.Atoi(sort)
	if err != nil {
		helper.ReturnJson(c, 0, "非法排序值")
		return
	}
	model.DB.Save(&link)
	helper.ReturnJson(c, 1, "修改成功")
}