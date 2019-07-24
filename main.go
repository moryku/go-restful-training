package main

import (
	"github.com/iswanulumam/go-restful-training/models"
	"github.com/iswanulumam/go-restful-training/routes"
)

func main() {
	e := routes.New()

	// init db connection
	models.InitDB("root:root123@/go_db?charset=utf8&parseTime=True&loc=Local")
	models.InitialMigration()

	e.Start(":8000")
}
