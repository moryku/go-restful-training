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

	e.Logger.Fatal(e.Start(":" + port))
}
