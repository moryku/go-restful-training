package main

import (
	"github.com/iswanulumam/go-restful-training/models"
	"github.com/iswanulumam/go-restful-training/routes"

	m "github.com/iswanulumam/go-restful-training/middlewares"
)

func main() {
	e := routes.New()

	// implement middleware with
	m.LogMiddlewares(e)

	// init db connection
	models.InitDB("root:root123@/crud_go?charset=utf8&parseTime=True&loc=Local")
	models.InitialMigration()

	e.Start(":8000")
}
