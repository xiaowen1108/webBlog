package admin

import (
	"github.com/gin-gonic/gin"
	"webBlog/helper"
	"net/http"
)

type Article struct {
}

func (t *Article) Index (c *gin.Context){

}

func (t *Article) Add (c *gin.Context){
	if helper.IsGet(c) {
		//获取所有的顶级分类
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/article/add.html",gin.H{
			"errorMsg":errorMsg,
		})
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