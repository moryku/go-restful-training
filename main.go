package main

import (
	"github.com/iswanulumam/go-restful-training/routes"

	m "github.com/iswanulumam/go-restful-training/middlewares"
)

func main() {
	e := routes.New()
	// implement middleware with
	m.LogMiddlewares(e)

	e.Start(":8000")
}
