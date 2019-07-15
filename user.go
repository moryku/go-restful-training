package main

import (
	"net/http"
	"fmt"
	
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *gorm.DBD
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err := gorm.Open("mysql", "root:root123@/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database.")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

// handler - Simple handler to make sure environment is setup
func helloController(c echo.Context) error {
  // return the string "Hello World" as the response body
  // with an http.StatusOK (200) status
  return c.String(http.StatusOK, "Hello World")
}

func allUsers(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:root123@/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users": users,
	})
}

func newUser(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:root123@/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	name := c.Param("name")
	email := c.Param("email")
	db.Create(&User{Name: name, Email: email})

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success create new user",
	})
}

func deleteUser(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:root123@/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	name := c.Param("name")	

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success delete user",
	})
}

func updateUser(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:root123@/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	
	name := c.Param("name")
	email := c.Param("email")

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email

	db.Save(&user)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "success delete user",
	})	
}