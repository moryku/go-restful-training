package config

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/db-sql-driver/mysql"
	"restful_training/models"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
}

func InitDB() {

	//db, err := gorm.Open("mysql", "root:root123@/crud_go?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("sqlite3", "./db/crud_go.db")

	if err != nil {
		panic(err)
	}
	DB = db
	models.InitialMigration()
}
