package main

import (
	"flag"
	"github.com/Unknwon/goconfig"
	"webBlog/model"
	"github.com/gin-gonic/gin"
	"webBlog/controller/admin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)
var appConfigFile, dbConfigFile string

func main() {
	//config file
	flag.StringVar(&appConfigFile, "app_conf_file", "config/app.ini", "web app config file")
	flag.StringVar(&dbConfigFile,"db_conf_file", "config/db.ini", "web db config file")
	flag.Parse()
	//read app.ini
	appConf, err := goconfig.LoadConfigFile(appConfigFile)
	checkErr(err)
	//read db.ini
	dbConf, err := goconfig.LoadConfigFile(dbConfigFile)
	checkErr(err)
	//read db select
	dbName, err := appConf.GetValue("", "db")
	checkErr(err)
	//read db config
	dbConfig, err := dbConf.GetSection(dbName)
	checkErr(err)
	//init DB
	DB, err := model.InitDB(dbName, dbConfig)
	checkErr(err)
	defer DB.Close()
	//gin start
	ginMode, err := appConf.GetValue("", "runMode")
	checkErr(err)
	gin.SetMode(ginMode)
	router := gin.Default()
	//session
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}) //Also set Secure: true if using SSL, you should though
	router.Use(sessions.Sessions("webBlog", store))
	//static
	router.Static("/static", "./static")
	//setRoute
	setRoute(router)
	runPort, err := appConf.GetValue("", "runPort")
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
	}
}
