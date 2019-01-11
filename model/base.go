package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"errors"
	"fmt"
	"time"
	"strconv"
)

type BaseModel struct {
	ID uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

var DB *gorm.DB

func InitDB(dbName string, dbConfig map[string]string) (*gorm.DB, error) {

	if dbName == "mysql" {
		//set prefix table name
		gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
			return dbConfig["prefix"] + defaultTableName;
		}
		var err error
		args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=%s&parseTime=True",
			dbConfig["username"],
			dbConfig["password"],
			dbConfig["host"],
			dbConfig["port"],
			dbConfig["database"],
			dbConfig["charset"],
			dbConfig["collation"],
			)
		//user:password@tcp(localhost:5555)/dbname?charset=xxx&collation=utf8mb4_unicode_ci&parseTime=True
		DB, err = gorm.Open(dbName, args)
		if err == nil {
			DB.LogMode(true)
			//AutoMigrate
			DB.AutoMigrate(&AdminUser{}, &Article{}, &Category{}, &Config{}, &Link{}, &Nav{}, &Tag{})
			//连接池
			maxIdleConns, err := strconv.Atoi(dbConfig["maxIdleConns"])
			if err != nil {
				maxIdleConns = 10
			}
			maxOpenConns, err := strconv.Atoi(dbConfig["maxOpenConns"])
			if err != nil {
				maxOpenConns = 100
			}
			DB.DB().SetMaxIdleConns(maxIdleConns)
			DB.DB().SetMaxOpenConns(maxOpenConns)
		}
		return DB, err
	}
	return nil, errors.New("InitDB Error")
}
