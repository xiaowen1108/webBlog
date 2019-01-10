package middleware

import (
	"github.com/gin-gonic/gin"
	"webBlog/helper"
	"net/http"
)

func CheckAdminLogin(exceptPath []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		var flag bool
		for _, path := range exceptPath {
			if path == url {
				flag = true
				break
			}
		}
		if flag {
			c.Next()
		} else {
			userInfo := helper.GetSession(c, "userInfo")
			if userInfo == nil {
				c.Redirect(http.StatusFound, "/admin/login")
				return
			} else {
				c.Next()
			}
		}
	}
}
