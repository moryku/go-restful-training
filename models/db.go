package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// migration schema
func InitialMigration() {
	db.AutoMigrate(Book{})
	db.AutoMigrate(User{})
}
