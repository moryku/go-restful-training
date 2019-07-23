package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get users",
				"users":    user,
			})
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users = remove(users, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}
	c.Bind(&user)
	users = update(users, id, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
	})
}

// ----------------------- lib ----------------------------

func remove(users []User, rmIdx int) []User {
	var idxRmv int
	for index, user := range users {
		if user.Id == rmIdx {
			idxRmv = index
		}
	}
	return append(users[:idxRmv], users[idxRmv+1:]...)
}

func update(users []User, id int, newUser User) []User {
	for idx, user := range users {
		if user.Id == id {
			users[idx] = newUser
		}
	}
	return users
}

// ---------------------------------------------------

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Start(":8000")
}
