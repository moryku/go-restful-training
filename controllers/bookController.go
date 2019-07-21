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

func GetBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := models.GetBookById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}

func GetBookLikeController(c echo.Context) error {
	name := c.QueryParam("name")
	books, err := models.GetBookLike(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

func CreateBookController(c echo.Context) error {
	title := c.FormValue("title")
	author := c.FormValue("author")
	publisher := c.FormValue("publisher")
	isbn := c.FormValue("isbn")
	price, _ := strconv.Atoi(c.FormValue("price"))

	book := models.Book{
		Title:     title,
		Author:    author,
		Publisher: publisher,
		Isbn:      isbn,
		Price:     price,
	}

	result, err := models.CreateBook(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create book")
	}
	return c.JSON(http.StatusCreated, result)
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := models.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to delete book")
	}
	return c.JSON(http.StatusOK, book)
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi("id")

	title := c.FormValue("title")
	author := c.FormValue("author")
	publisher := c.FormValue("publisher")
	isbn := c.FormValue("isbn")
	price, _ := strconv.Atoi(c.FormValue("price"))

	book, err := models.UpdateBook(id, title, author, publisher, isbn, price)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update book")
	}
	return c.JSON(http.StatusOK, book)
}
