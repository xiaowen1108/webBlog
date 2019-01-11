package main

import (
	"flag"
	"webBlog/model"
	"github.com/gin-gonic/gin"
	"webBlog/controller/admin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"fmt"
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
	//自定义admin auth中间件
	//adminR.Use(middleware.CheckAdminLogin([]string{"/admin/login", "/admin/code"}))
	{
		//登录
		loginController := &admin.Login{}
		adminR.GET("/login", loginController.Login)
		adminR.POST("/login", loginController.Login)
		adminR.GET("/layout", loginController.Logout)
		adminR.GET("/code", loginController.Code)
		//首页
		indexController := &admin.Index{}
		adminR.GET("/index", indexController.Index)
		adminR.GET("/", indexController.Index)
		adminR.GET("/info", indexController.Info)
		adminR.GET("/pass", indexController.Pass)
		adminR.POST("/pass", indexController.Pass)
		adminR.POST("/recom", indexController.Recom)
		//分类
		categoryController := &admin.Category{}
		adminR.GET("/category/index", categoryController.Index)
		adminR.GET("/category/add", categoryController.Add)
		adminR.POST("/category/add", categoryController.Add)
		adminR.GET("/category/edit/:id", categoryController.Edit)
		adminR.POST("/category/edit/:id", categoryController.Edit)
		adminR.POST("/category/del/:id", categoryController.Del)
		adminR.POST("/category/changeorder", categoryController.ChangeOrder)
		//文章
		articleController := &admin.Article{}
		adminR.GET("/article/index", articleController.Index)
		adminR.GET("/article/add", articleController.Add)
		adminR.POST("/article/add", articleController.Add)
		adminR.POST("/article/edit/:id", articleController.Edit)
		adminR.GET("/article/del/:id", articleController.Del)
		adminR.POST("/article/changeorder/:id", articleController.ChangeOrder)
		adminR.POST("/article/set_recom/:id", articleController.SetRecom)
		adminR.GET("/article/recom", articleController.Recom)
		//base
		baseController := &admin.Base{}
		adminR.POST("/upload", baseController.Upload)
		//友情链接
		linkController := &admin.Link{}
		adminR.GET("/links/index", linkController.Index)
		adminR.GET("/links/add", linkController.Add)
		adminR.POST("/links/edit/:id", linkController.Edit)
		adminR.GET("/links/del/:id", linkController.Del)
		adminR.POST("/links/changeorder", linkController.ChangeOrder)
		//导航
		navController := &admin.Link{}
		adminR.GET("/navs/index", navController.Index)
		adminR.GET("/navs/add", navController.Add)
		adminR.POST("/navs/edit/:id", navController.Edit)
		adminR.GET("/navs/del/:id", navController.Del)
		adminR.POST("/navs/changeorder", navController.ChangeOrder)
		//设置
		confController := &admin.Config{}
		adminR.GET("/config/index", confController.Index)
		adminR.GET("/config/add", confController.Add)
		adminR.POST("/config/edit/:id", confController.Edit)
		adminR.GET("/config/del/:id", confController.Del)
		adminR.POST("/config/changeorder", confController.ChangeOrder)
		adminR.GET("/config/putfile", confController.PutFile)
		adminR.POST("/config/changecontent", confController.ChangeContent)
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

