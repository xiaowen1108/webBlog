package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/sessions"

	"fmt"
)

type Login struct {
}

func (l Login) Login (c *gin.Context)  {
	session := sessions.Default(c)
	errorMsg := session.Get("errorMsg")
	fmt.Println(errorMsg)
	session.Set("errMsg", "hahahahah")
	session.Save()
	c.HTML(http.StatusOK, "admin/login/login.html",gin.H{
		"errorMsg":errorMsg,
	})
}
