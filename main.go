package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func CreateUser(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	// routing with query parameter
	e.POST("/users", CreateUser)

	e.Start(":8000")
}
