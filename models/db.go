package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func InitialMigration() {
	db.AutoMigrate(Book{})
	db.AutoMigrate(User{})
}

func InitDB(connectionString string) {
	var err error
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
}
