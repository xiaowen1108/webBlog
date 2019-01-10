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
)

func main() {
	//config file
	var configFile string
	flag.StringVar(&configFile, "app_conf_file", "config/app.ini", "web app config file")
	flag.Parse()
	//set config
	helper.InitConfigManager(configFile)
	config := helper.GetConfig()
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
	router.LoadHTMLGlob("./view/***/**/*")
	router.Run(config.GetValue("app", "runPort"))
}

func setRoute(r *gin.Engine){
	adminR := r.Group("/admin")
	adminR.Use(checkAdminLogin([]string{"/admin/login", "/admin/code"}))
	{
		adminR.GET("/login", admin.Login{}.Login)
		adminR.POST("/login", admin.Login{}.Login)
		adminR.GET("/code", admin.Login{}.Code)
		adminR.GET("/index", admin.Login{}.Index)
				//'login', 'LoginController@login');
				//Route::get('code', 'LoginController@code');
				//});
				//Route::group(['prefix' => 'admin','namespace'=>'Admin','middleware'=>'admin.auth'], function () {
				//Route::get('index', 'IndexController@index');
				//Route::get('/', 'IndexController@index');
				//Route::get('info', 'IndexController@info');
				//Route::get('layout', 'IndexController@layout');
				//Route::any('pass', 'IndexController@pass');
				//Route::post('category/changeorder', 'CategoryController@changeorder');
				//Route::resource('category', 'CategoryController');
				//Route::post('article/changeorder', 'ArticleController@changeorder');
				//Route::post('article/set_recom/{id}', 'ArticleController@set_recom');
				//Route::get('article/recom', 'ArticleController@recom');
				//Route::resource('article', 'ArticleController');
				//Route::get('recom', 'IndexController@recom');
				//Route::any('upload', 'BaseController@upload');
				//Route::post('links/changeorder', 'LinksController@changeorder');
				//Route::resource('links', 'LinksController');
				//Route::post('navs/changeorder', 'NavsController@changeorder');
				//Route::resource('navs', 'NavsController');
				//Route::get('config/putfile', 'ConfigController@putFile');
				//Route::post('config/changecontent', 'ConfigController@changeContent');
				//Route::post('config/changeorder', 'ConfigController@changeOrder');
				//Route::resource('config', 'ConfigController');
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

