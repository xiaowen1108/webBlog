package admin

import (
	"github.com/gin-gonic/gin"
	"webBlog/helper"
	"net/http"
	"webBlog/model"
	"strconv"
	"html/template"
	"strings"
)

type Article struct {
}

func (t *Article) Index (c *gin.Context){
	//分页
	p := c.DefaultQuery("page", "1")
	pnum := 10
	pi, err := strconv.Atoi(p)
	if err != nil {
		pi = 1
	}
	pi = (pi - 1) * pnum
	var articles []*model.Article
	model.DB.Select("id, title,cover,is_recom,created_at,updated_at").Offset(pi).Limit(pnum).Order("id desc").Find(&articles)
	var count int
	model.DB.Model(&model.Article{}).Count(&count)
	pagination := helper.NewPagination(c.Request, count, pnum)
	c.HTML(http.StatusOK, "admin/article/index.html",gin.H{
		"articles":articles,
		"pages":template.HTML(pagination.Pages()),
	})
}
type ArticleData struct {
	Title string `form:"name" binding:"required"`
	Cid int `form:"cid" binding:"required"`
	Description string `form:"description"`
	Tags string `form:"tags"`
	Cover string `form:"cover" binding:"required"`
	Content string`form:"content"`
}
func (t *Article) Add (c *gin.Context){
	if helper.IsGet(c) {
		//获取所有的分类
		categorys := model.GetTree()
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/article/add.html",gin.H{
			"errorMsg":errorMsg,
			"categorys":categorys,
		})
	} else if helper.IsPost(c) {
		var articleData ArticleData
		if err := c.ShouldBind(&articleData); err == nil {
			article := model.Article{
				Title:articleData.Title,
				Cid:articleData.Cid,
				Description:articleData.Description,
				Tags:strings.Replace(articleData.Tags,"，", "," , -1),
				Cover:articleData.Cover,
				Content:articleData.Content,
			}
			err := model.DB.Create(&article).Error
			if err == nil {
				model.HandelTagAdd(article.Tags)
				helper.SetFlash(c, "errorMsg", "添加成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "添加失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/article/add")
	}
}

func (t *Article) Edit (c *gin.Context){
	id := c.Param("id")
	var article model.Article
	err := model.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/article/index")
		return
	}
	oldTags := article.Tags
	if helper.IsGet(c) {
		//获取所有的分类
		categorys := model.GetTree()
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/article/edit.html",gin.H{
			"errorMsg":errorMsg,
			"categorys":categorys,
			"article":article,
		})
	} else if helper.IsPost(c) {
		var articleData ArticleData
		if err := c.ShouldBind(&articleData); err == nil {
			article.Title = articleData.Title
			article.Cid = articleData.Cid
			article.Description = articleData.Description
			article.Tags = articleData.Tags
			article.Cover = articleData.Cover
			article.Content = articleData.Content
			err := model.DB.Save(&article).Error
			if err == nil {
				model.HandelTagEdit(oldTags, article.Tags)
				helper.SetFlash(c, "errorMsg", "编辑成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "编辑失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/article/eidt/"+id)
	}
}

func (t *Article) Del (c *gin.Context){

}

func (t *Article) ChangeOrder (c *gin.Context){

}

func (t *Article) SetRecom (c *gin.Context){

}

func (t *Article) Recom (c *gin.Context){

}