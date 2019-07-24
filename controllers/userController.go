package controllers

import (
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

func GetUsersLikeControlller(c echo.Context) error {
	name := c.QueryParam("name")
	users, err := models.GetUsersLike(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, err := models.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

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
