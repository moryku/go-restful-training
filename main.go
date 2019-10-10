package main

import (
	m "restful_training/middlewares"
	"restful_training/routes"
)

func main() {
	e := routes.New()
	// implement middleware with
	m.LogMiddlewares(e)

	e.Start(":8000")
}
