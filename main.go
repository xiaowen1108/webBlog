package main

import (
	"flag"
	"github.com/Unknwon/goconfig"
	"webBlog/model"
	"github.com/gin-gonic/gin"
	"webBlog/controller/admin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	"crypto/sha256"
)
var appConfigFile, dbConfigFile string
var appConf,dbConf *goconfig.ConfigFile
func main() {
	//config file
	flag.StringVar(&appConfigFile, "app_conf_file", "config/app.ini", "web app config file")
	flag.StringVar(&dbConfigFile,"db_conf_file", "config/db.ini", "web db config file")
	flag.Parse()
	//read app.ini
	var err error
	appConf, err = goconfig.LoadConfigFile(appConfigFile)
	checkErr(err)
	//read db.ini
	dbConf, err = goconfig.LoadConfigFile(dbConfigFile)
	checkErr(err)
	//read db select
	dbName, err := appConf.GetValue("app", "db")
	checkErr(err)
	//read db config
	dbConfig, err := dbConf.GetSection(dbName)
	checkErr(err)
	//init DB
	DB, err := model.InitDB(dbName, dbConfig)
	checkErr(err)
	defer DB.Close()
	//init user
	createAdminUser(DB)
	//gin start
	ginMode, err := appConf.GetValue("app", "runMode")
	checkErr(err)
	gin.SetMode(ginMode)
	router := gin.Default()
	//session
	secret, err := appConf.GetValue("app", "secret")
	checkErr(err)
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}) //Also set Secure: true if using SSL, you should though
	sessionName, err := appConf.GetValue("app", "sessionName")
	checkErr(err)
	router.Use(sessions.Sessions(sessionName, store))
	//static
	router.Static("/static", "./static")
	//setRoute
	setRoute(router)
	runPort, err := appConf.GetValue("app", "runPort")
	checkErr(err)
	//setView
	router.LoadHTMLGlob("./view/***/**/*")
	router.Run(runPort)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func setRoute(r *gin.Engine){
	adminR := r.Group("/admin")
	adminR.Use(checkAdminLogin())
	{
		adminR.GET("/login", admin.Login{}.Login)
		adminR.POST("/login", admin.Login{}.Login)
		adminR.GET("/code", admin.Login{}.Code)
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

func checkAdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL)
		url := c.Request.URL.Path
		if url == "/admin/login" || url == "/admin/code" {
			c.Next()
		} else {
			session := sessions.Default(c)
			userInfo := session.Get("userInfo")
			if userInfo == nil {
				c.Redirect(http.StatusMovedPermanently, "/admin/login")
			} else {
				c.Next()
			}
		}
	}
}

func createAdminUser(DB *gorm.DB) {
	var adminUser model.AdminUser

	if DB.First(&adminUser).RecordNotFound() {
		//创建用户
		username, err := appConf.GetValue("account", "username")
		checkErr(err)
		password, err := appConf.GetValue("account", "password")
		checkErr(err)
		adminUser.Name = username
		adminUser.Nickname = username
		h := sha256.New()
		h.Write([]byte(password))
		secret, err := appConf.GetValue("app", "secret")
		checkErr(err)
		adminUser.Pwd = fmt.Sprintf("%x", h.Sum([]byte(secret)))
		DB.Create(&adminUser)
	}

}
