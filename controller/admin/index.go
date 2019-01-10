package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
	"webBlog/helper"
	"net"
	"webBlog/model"
	"crypto/sha256"
	"fmt"
)

type Index struct {
	*Base
}

//首页
func (i Index) Index (c *gin.Context)  {
	userInfo := helper.GetSession(c, "userInfo")
	c.HTML(http.StatusOK, "admin/index/index.html",  gin.H{
		"nickName":userInfo,
	})
}

//info
func (i Index) Info (c *gin.Context)  {
	addrs, err := net.InterfaceAddrs()
	helper.CheckErr(err)
	var ip string
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	vars := gin.H{
		"os":runtime.GOOS,
		"goarch":runtime.GOARCH,
		"time":time.Now(),
		"ip":ip,
	}
	c.HTML(http.StatusOK, "admin/index/info.html",  vars)
}
type PassData struct {
	OldPwd string `form:"old_pwd" binding:"required"`
	NewPwd string `form:"new_pwd" binding:"required"`
	CekPwd string `form:"new_pwd_confirmation" binding:"required"`
}
//修改密码
func (i Index) Pass (c *gin.Context)  {
	//展示登录页面
	if helper.IsGet(c) {
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/index/pass.html",gin.H{
			"errorMsg":errorMsg,
		})
	} else if helper.IsPost(c) {
		var passData PassData
		if err := c.ShouldBind(&passData); err == nil{

			if passData.CekPwd != passData.NewPwd {
				helper.SetFlash(c, "errorMsg", "确认密码不一致 ！")
				c.Redirect(http.StatusFound, "/admin/pass")
				return
			}
			config := helper.GetConfig()
			var adminUser model.AdminUser
			model.DB.First(&adminUser)
			h := sha256.New()
			h.Write([]byte(passData.OldPwd))
			secret := config.GetValue("app", "secret")
			passData.OldPwd = fmt.Sprintf("%x", h.Sum([]byte(secret)))
			if adminUser.Pwd != passData.OldPwd {
				helper.SetFlash(c, "errorMsg", "原密码不正确 ！")
				c.Redirect(http.StatusFound, "/admin/pass")
				return
			}
			h = sha256.New()
			h.Write([]byte(passData.NewPwd))
			adminUser.Pwd = fmt.Sprintf("%x", h.Sum([]byte(secret)))
			err := model.DB.Save(&adminUser).Error
			if err != nil {
				helper.SetFlash(c, "errorMsg", "修改失败 ！")
				c.Redirect(http.StatusFound, "/admin/pass")
			} else {
				helper.SetFlash(c, "errorMsg", "修改成功 ！")
				c.Redirect(http.StatusFound, "/admin/pass")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "请检查账户是否输入完全 ！")
			c.Redirect(http.StatusFound, "/admin/pass")
		}
	}
}

func (i Index) Recom (c *gin.Context){

}