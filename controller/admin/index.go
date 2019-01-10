package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
	"webBlog/helper"
	"net"
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
