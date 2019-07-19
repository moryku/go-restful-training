package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/iswanulumam/go-restful-training/models"
	"github.com/labstack/echo"
)

func FindAllUserController(c echo.Context) error {
	users, err := models.FindAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUserController(c echo.Context) error {
	user := models.User{}

	name := c.FormValue("name")
	email := c.FormValue("email")

	defer c.Request().Body.Close()

	user.Name = name
	user.Email = email

	result, err := models.CreateUser(&user)
	if err != nil {
		log.Printf("FAILED TO CREATE : %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create new user")
	}
	return c.JSON(http.StatusCreated, result)
}

func DeleteUserController(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("id"))
	models.DeleteUserByID(ID)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "success delete user",
	})
}
