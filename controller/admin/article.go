package admin

import "github.com/gin-gonic/gin"

type Article struct {
	*Base
}

func (t Article) Index (c *gin.Context){

}

func (t Article) Add (c *gin.Context){

}

func (t Article) Edit (c *gin.Context){

}

func (t Article) Del (c *gin.Context){

}

func (t Article) ChangeOrder (c *gin.Context){

}

func (t Article) SetRecom (c *gin.Context){

}

func (t Article) Recom (c *gin.Context){

}