package controllers

import (
	"net/http"
	"strconv"

	"github.com/iswanulumam/go-restful-training/models"
	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, err := models.GetUsers()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// get single user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.GetUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, err := models.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

// remove user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := models.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	newUser := models.User{}
	c.Bind(&newUser)
	result, err := models.UpdateUser(id, newUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
