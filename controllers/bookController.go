package controllers

import (
	"net/http"
	"strconv"

	"github.com/iswanulumam/go-restful-training/models"
	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	books, err := models.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := models.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

func GetBooksLikeController(c echo.Context) error {
	name := c.QueryParam("name")
	books, err := models.GetBooksLike(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	result, err := models.CreateBook(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := models.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi("id")

	newBook := models.Book{}
	c.Bind(&newBook)

	book, err := models.UpdateBook(id, &newBook)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}
