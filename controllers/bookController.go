package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/moryku/go-restful-training/models"
)

// get all books
func GetBooksController(c echo.Context) error {
	books, err := models.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

// get single book by id
func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := models.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// get book by match input
func GetBooksLikeController(c echo.Context) error {
	name := c.QueryParam("name")
	books, err := models.GetBooksLike(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	result, err := models.CreateBook(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

// remove book by id
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := models.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

// update book by id
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
