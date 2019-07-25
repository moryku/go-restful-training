package main

import (
	"net/http"

	"log"
	"strconv"

	m "github.com/iswanulumam/go-restful-training/middlewares"
	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// --------------------- model ----------------------

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func InitialMigration() {
	db.AutoMigrate(&User{})
}

// database connection
func InitDB(connectionString string) {
	var err error
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
}

// --------------------- controller ----------------------

// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	if err := db.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user User
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := db.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var user User
	id, _ := strconv.Atoi(c.Param("id"))

	db.Where("ID = ?", id).Find(&user)
	if err := db.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    &user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// binding new data
	newUser := User{}
	c.Bind(&newUser)

	user := User{}
	db.Where("ID = ?", id).Find(&user)
	if err := db.Save(&newUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func main() {
	// create a new echo instance
	e := echo.New()

	InitDB("root:root123@/go_db?charset=utf8&parseTime=True&loc=Local")
	InitialMigration()

	// implement middleware logger with call in function
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	// }))

	// implement middleware with
	m.LogMiddlewares(e)

	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Start(":8000")
}
