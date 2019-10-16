package db

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/moryku/go-restful-training/config"
	"github.com/moryku/go-restful-training/models"
)

// migration schema
func InitialMigration() {
	config.DB.AutoMigrate(models.Book{})
	config.DB.AutoMigrate(models.User{})
}
