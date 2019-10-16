package main

import (
	"github.com/moryku/go-restful-training/config"
	"github.com/moryku/go-restful-training/db"
	m "github.com/moryku/go-restful-training/middlewares"
	"github.com/moryku/go-restful-training/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	config.InitDB()
	db.InitialMigration()

	e := routes.New()
	m.LogMiddlewares(e)
	if port == "" {
		port = "9000"
		//log.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
