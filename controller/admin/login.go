package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webBlog/helper"
	"github.com/steambap/captcha"
	"webBlog/model"
	"fmt"
	"crypto/sha256"
	"strings"
)

type Login struct {
}

type LoginData struct {
	Username string `form:"name" binding:"required"`
	Password string `form:"pwd" binding:"required"`
	Cacode string `form:"code" binding:"required"`
}
//登录页面
func (l *Login) Login (c *gin.Context)  {

	//展示登录页面
	if helper.IsGet(c) {
		errorMsg := helper.GetFlash(c, "loginErrorMsg")
		c.HTML(http.StatusOK, "admin/login/login.html",gin.H{
			"errorMsg":errorMsg,
		})
	} else if helper.IsPost(c) {
		//登录逻辑
		var loginData LoginData
		if err := c.ShouldBind(&loginData); err == nil{
			//判断code
			ncode := (helper.GetFlash(c,"code")).(string)
			if strings.ToLower(loginData.Cacode) != strings.ToLower(ncode) {
				helper.SetFlash(c, "loginErrorMsg", "验证码输入有误 ！")
				c.Redirect(http.StatusFound, "/admin/login")
				return
			}
			var adminUser model.AdminUser
			config := helper.GetConfig()
			model.DB.First(&adminUser)
			h := sha256.New()
			h.Write([]byte(loginData.Password))
			secret := config.GetValue("app", "secret")
			loginData.Password = fmt.Sprintf("%x", h.Sum([]byte(secret)))
			if loginData.Username == adminUser.Name && loginData.Password == adminUser.Pwd {
				helper.SetSession(c, "userInfo", adminUser.Nickname)
				c.Redirect(http.StatusFound, "/admin/index")
			} else {
				helper.SetFlash(c, "loginErrorMsg", "用户名或密码有误 ！")
				c.Redirect(http.StatusFound, "/admin/login")
			}
		} else {
			helper.SetFlash(c, "loginErrorMsg", "请检查账户是否输入完全 ！")
			c.Redirect(http.StatusFound, "/admin/login")
		}
	}
}
//验证码
func (l *Login) Code (c *gin.Context)  {
	data, _ := captcha.New(102, 35)

	helper.SetFlash(c, "code", data.Text)
	data.WriteImage(c.Writer)
}

//退出
func (l *Login) Logout (c *gin.Context)  {
	helper.ClearSession(c)
	c.Redirect(http.StatusFound, "/admin/login")
}

