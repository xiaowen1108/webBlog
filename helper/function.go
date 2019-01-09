package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func SetFlash(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.AddFlash(value, key)
	session.Save()
}

func GetFlash(c *gin.Context, key string)  interface{}{
	session := sessions.Default(c)
	flash := session.Flashes(key)
	session.Save()
	if len(flash) > 0 {
		return flash[0]
	}
	return ""
}

func IsGet(c *gin.Context) bool {
	return c.Request.Method == "GET"
}

func IsPost(c *gin.Context) bool  {
	return c.Request.Method == "POST"
}
