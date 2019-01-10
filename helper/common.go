package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)

func SetFlash(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Delete(key)
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
func GetSession(c *gin.Context, key string)  interface{}{
	session := sessions.Default(c)
	return session.Get(key)
}

func SetSession(c *gin.Context, key,value interface{}){
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}
func ClearSession(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
func IsGet(c *gin.Context) bool {
	return c.Request.Method == "GET"
}

func IsPost(c *gin.Context) bool  {
	return c.Request.Method == "POST"
}

func CheckErr(err error) {
	if err != nil{
		panic(err)
	}
}

func ReturnJson(c *gin.Context, status int, info string) {
	c.JSON(http.StatusOK, gin.H{
		"status":status,
		"info":info,
	})
}