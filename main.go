package main

import (
  "github.com/labstack/echo"
)

func main() {
  // create a new echo instance
  e := echo.New()
  // Route / to handler function
	e.GET("/", helloController)
  e.GET("/users", allUsers)
  e.POST("/user/:name/:email", newUser)
  e.DELETE("/user/:name", deleteUser)
  e.PUT("/user/:name/:email", updateUser)
  
  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8081"))
}
