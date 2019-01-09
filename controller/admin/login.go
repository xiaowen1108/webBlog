package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webBlog/helper"
	"github.com/steambap/captcha"
)

type Login struct {
}
//登录页面
func (l Login) Login (c *gin.Context)  {
	//展示登录页面
	if helper.IsGet(c) {
		errorMsg := helper.GetFlash(c, "loginErrorMsg")
		c.HTML(http.StatusOK, "admin/login/login.html",gin.H{
			"errorMsg":errorMsg,
		})
	}
	//登录逻辑
	if helper.IsPost(c) {
		ncode := (helper.GetFlash(c,"code")).(string)
		c.String(http.StatusOK, ncode)
	}
}
//验证码
func (l Login) Code (c *gin.Context)  {
	data, _ := captcha.New(102, 35)
	helper.SetFlash(c, "code", data.Text)
	data.WriteImage(c.Writer)
}