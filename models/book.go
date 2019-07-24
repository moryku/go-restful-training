package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Isbn      string `json:"isbn" form:"isbn"`
	Price     int    `json:"price" form:"price"`
}

func GetBooks() (interface{}, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookById(id int) (interface{}, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBookLike(match string) (interface{}, error) {
	var books []Book
	src := "%" + match + "%"
	if err := db.Where("title LIKE ?", src).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBook(book *Book) (*Book, error) {
	if err := db.Create(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (*Book, error) {
	var book Book
	db.Where("ID = ?", id).Find(&book)
	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBook(id int, title string, author string, publisher string, isbn string, price int) (*Book, error) {
	var book Book
	db.Where("ID = ?", id).Find(&book)

	book.Title = title
	book.Author = author
	book.Publisher = publisher
	book.Isbn = isbn
	book.Price = price

	if err := db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
