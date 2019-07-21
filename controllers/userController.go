package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/iswanulumam/go-restful-training/models"
	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := models.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.GetUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func GetUserLikeControlller(c echo.Context) error {
	name := c.QueryParam("name")
	users, err := models.GetUserLike(name)
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
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := models.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete user")
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	name := c.FormValue("name")
	email := c.FormValue("email")

	user, err := models.UpdateUser(id, name, email)

	if err != nil {
		log.Printf("FAILED TO UPDATE: %s\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update user")
	}
	return c.JSON(http.StatusOK, user)
}
