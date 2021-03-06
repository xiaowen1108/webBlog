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
		c.Redirect(http.StatusFound, "/admin/article/edit/"+id)
	}
}

func (t *Article) Del (c *gin.Context){
	id := c.Param("id")
	var article model.Article
	err := model.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该文章，删除失败")
		return
	}
	model.HandelTagDel(article.Tags)
	model.DB.Delete(&article)
	helper.ReturnJson(c, 1, "删除成功")
	return
}

func (t *Article) ChangeOrder (c *gin.Context){
	id := c.PostForm("id")
	sort := c.PostForm("sort")
	if id =="" || sort == "" {
		helper.ReturnJson(c, 0, "参数错误")
		return
	}
	var article model.Article
	err := model.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该文章，设置失败")
		return
	}
	sortInt, err := strconv.Atoi(sort)
	if err != nil {
		helper.ReturnJson(c, 0, "非法参数")
		return
	}
	model.DB.Model(&article).Update("sort", sortInt)
	helper.ReturnJson(c, 1, "设置成功")
}

func (t *Article) SetRecom (c *gin.Context){
	id := c.PostForm("id")
	recom := c.PostForm("recom")
	if id =="" || recom == "" {
		helper.ReturnJson(c, 0, "参数错误")
		return
	}
	var article model.Article
	err := model.DB.Where("id = ?", id).First(&article).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该文章，设置失败")
		return
	}
	isRecom, err := strconv.ParseBool(recom)
	if err != nil {
		helper.ReturnJson(c, 0, "非法参数")
		return
	}
	if isRecom {
		model.DB.Model(&article).Update("is_recom", false)
		helper.ReturnJson(c, 1, "取消推荐成功")
	} else {
		model.DB.Model(&article).Update("is_recom", true)
		helper.ReturnJson(c, 1, "推荐成功")
	}
}

func (t *Article) Recom (c *gin.Context){
	p := c.DefaultQuery("page", "1")
	pnum := 10
	pi, err := strconv.Atoi(p)
	if err != nil {
		pi = 1
	}
	pi = (pi - 1) * pnum
	var articles []*model.Article
	model.DB.Select("id,title,cover,is_recom,sort,created_at,updated_at").Where("is_recom = ?", 1).Offset(pi).Limit(pnum).Order("sort desc,id desc").Find(&articles)
	var count int
	model.DB.Model(&model.Article{}).Count(&count)
	pagination := helper.NewPagination(c.Request, count, pnum)
	c.HTML(http.StatusOK, "admin/article/recom.html",gin.H{
		"articles":articles,
		"pages":template.HTML(pagination.Pages()),
	})
}