package main

import (
	"flag"
	"webBlog/model"
	"github.com/gin-gonic/gin"
	"webBlog/controller/admin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"crypto/sha256"
	"webBlog/helper"
	"html/template"
	"time"
)

func main() {
	//config file
	var configFile string
	flag.StringVar(&configFile, "app_conf_file", "config/app.ini", "web app config file")
	flag.Parse()
	//set config
	helper.InitConfigManager(configFile)
	config := helper.GetConfig()
	//time
	timelocal,_ := time.LoadLocation("Asia/Shanghai")
	time.Local = timelocal
	//init DB
	dbName := config.GetValue("app", "db")
	DB, err := model.InitDB(dbName, config.GetSection(dbName))
	helper.CheckErr(err)
	defer DB.Close()
	//init user
	createAdminUser(DB)
	//gin start
	gin.SetMode(config.GetValue("app", "runMode"))
	router := gin.Default()
	//session init
	store := cookie.NewStore([]byte(config.GetValue("app", "secret")))
	//store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}) //Also set Secure: true if using SSL, you should though
	router.Use(sessions.Sessions(config.GetValue("app", "sessionName"), store))
	//static
	router.Static("/static", "./static")
	//setRoute
	setRoute(router)
	//setView
	funcMap := template.FuncMap{
		"url": helper.Url,
	}
	router.SetFuncMap(funcMap)
	router.LoadHTMLGlob("./view/***/**/*")
	router.Run(config.GetValue("app", "runPort"))
}

func setRoute(r *gin.Engine){
	adminR := r.Group("/admin")
	//adminR.Use(checkAdminLogin([]string{"/admin/login", "/admin/code"}))
	{
		//登录
		adminR.GET("/login", admin.Login{}.Login)
		adminR.POST("/login", admin.Login{}.Login)
		adminR.GET("/layout", admin.Login{}.Logout)
		adminR.GET("/code", admin.Login{}.Code)
		//首页
		adminR.GET("/index", admin.Index{}.Index)
		adminR.GET("/", admin.Index{}.Index)
		adminR.GET("/info", admin.Index{}.Info)
		adminR.GET("/pass", admin.Index{}.Pass)
		adminR.POST("/pass", admin.Index{}.Pass)
		adminR.POST("/recom", admin.Index{}.Recom)
		//分类
		adminR.GET("/category/index", admin.Category{}.Index)
		adminR.GET("/category/edit", admin.Category{}.Add)
		adminR.POST("/category/edit/:id", admin.Category{}.Edit)
		adminR.GET("/category/del/:id", admin.Category{}.Del)
		adminR.POST("/category/changeorder/:id", admin.Category{}.ChangeOrder)
		//文章
		adminR.GET("/article/index", admin.Article{}.Index)
		adminR.GET("/article/edit", admin.Article{}.Add)
		adminR.POST("/article/edit/:id", admin.Article{}.Edit)
		adminR.GET("/article/del/:id", admin.Article{}.Del)
		adminR.POST("/article/changeorder/:id", admin.Article{}.ChangeOrder)
		adminR.POST("/article/set_recom/:id", admin.Article{}.SetRecom)
		adminR.GET("/article/recom", admin.Article{}.Recom)
		//base
		adminR.POST("/upload", admin.Base{}.Upload)
		//友情链接
		adminR.GET("/links/index", admin.Link{}.Index)
		adminR.GET("/links/edit", admin.Link{}.Add)
		adminR.POST("/links/edit/:id", admin.Link{}.Edit)
		adminR.GET("/links/del/:id", admin.Link{}.Del)
		adminR.POST("/links/changeorder", admin.Link{}.ChangeOrder)
		//导航
		adminR.GET("/navs/index", admin.Nav{}.Index)
		adminR.GET("/navs/edit", admin.Nav{}.Add)
		adminR.POST("/navs/edit/:id", admin.Nav{}.Edit)
		adminR.GET("/navs/del/:id", admin.Nav{}.Del)
		adminR.POST("/navs/changeorder", admin.Nav{}.ChangeOrder)
		//设置
		adminR.GET("/config/index", admin.Config{}.Index)
		adminR.GET("/config/edit", admin.Config{}.Add)
		adminR.POST("/config/edit/:id", admin.Config{}.Edit)
		adminR.GET("/config/del/:id", admin.Config{}.Del)
		adminR.POST("/config/changeorder", admin.Config{}.ChangeOrder)
		adminR.GET("/config/putfile", admin.Config{}.PutFile)
		adminR.POST("/config/changecontent", admin.Config{}.ChangeContent)
	}
}

func checkAdminLogin(exceptPath []string) gin.HandlerFunc {
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

func createAdminUser(DB *gorm.DB) {
	var adminUser model.AdminUser

	if DB.First(&adminUser).RecordNotFound() {
		config := helper.GetConfig()
		//创建用户
		username := config.GetValue("account", "username")
		password := config.GetValue("account", "password")
		adminUser.Name = username
		adminUser.Nickname = username
		h := sha256.New()
		h.Write([]byte(password))
		secret := config.GetValue("account", "secret")
		adminUser.Pwd = fmt.Sprintf("%x", h.Sum([]byte(secret)))
		DB.Create(&adminUser)
	}

}

