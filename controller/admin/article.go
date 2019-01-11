package admin

import (
	"github.com/gin-gonic/gin"
	"webBlog/helper"
	"net/http"
	"webBlog/model"
	"fmt"
)

type Article struct {
}

func (t *Article) Index (c *gin.Context){

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
				Tags:articleData.Tags,
				Cover:articleData.Cover,
				Content:articleData.Content,
			}
			err := model.DB.Create(&article).Error
			if err == nil {
				fmt.Println(article.ID)
				model.HandelTagAdd(articleData.Tags)
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

}

func (t *Article) Del (c *gin.Context){

}

func (t *Article) ChangeOrder (c *gin.Context){

}

func (t *Article) SetRecom (c *gin.Context){

}

func (t *Article) Recom (c *gin.Context){

}