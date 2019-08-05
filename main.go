package main

import (
	"github.com/iswanulumam/go-restful-training/config"
	"github.com/iswanulumam/go-restful-training/models"
	"github.com/iswanulumam/go-restful-training/routes"

	m "github.com/iswanulumam/go-restful-training/middlewares"
)

func main() {
	e := routes.New()

	// implement middleware
	m.LogMiddlewares(e)
	// migration database
	InitialMigration()

	e.Start(":8000")
}

func InitialMigration() {
	var db = config.DB
	if !db.HasTable(&models.User{}) { // check database exist or not
		db.AutoMigrate(&models.User{})
	}
}
