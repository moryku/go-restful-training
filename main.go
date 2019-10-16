package main

import (
	"log"
	"os"
	m "restful_training/middlewares"
	"restful_training/routes"
)

func main() {
	port := os.Getenv("PORT")

	e := routes.New()
	// implement middleware with
	m.LogMiddlewares(e)
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}
