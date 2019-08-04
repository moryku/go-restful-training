package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
}

func InitDB() {
	db, err := gorm.Open("mysql", "root:root123@/crud_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = db
}
