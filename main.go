package main

import (
	m "middlewares"
	"routes"
)

func main() {
	e := routes.New()
	// implement middleware with
	m.LogMiddlewares(e)

	e.Start(":8000")
}
