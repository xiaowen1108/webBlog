package main

import (
	"flag"
	"github.com/Unknwon/goconfig"
	"webBlog/model"
)

func main() {
	//config file
	var appConfigFile, dbConfigFile string
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

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}