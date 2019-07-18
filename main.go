package main

import (
	"fmt"

	"github.com/iswanulumam/go-restful-training/routes"
)

func main() {
	e := routes.New()

	fmt.Println("App listening on port :8000")
	e.Start(":8080")
}
