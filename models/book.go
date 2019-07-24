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

func GetBook(id int) (interface{}, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBooksLike(match string) (interface{}, error) {
	var books []Book
	src := "%" + match + "%"
	if err := db.Where("title LIKE ?", src).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBook(book *Book) (interface{}, error) {
	if err := db.Create(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book Book
	db.Where("ID = ?", id).Find(&book)
	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBook(id int, newBook *Book) (interface{}, error) {
	var book Book
	db.Where("ID = ?", id).Find(&book)
	if err := db.Save(&newBook).Error; err != nil {
		return nil, err
	}
	return newBook, nil
}
