package models

import (
	"github.com/iswanulumam/go-restful-training/config"
	"github.com/jinzhu/gorm"
)

var db = config.DB

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// get all user
func GetUsers() (interface{}, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// get single user by id
func GetUser(id int) (interface{}, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// get many user by match input
func GetUsersLike(match string) (interface{}, error) {
	var users []User
	src := "%" + match + "%"
	if err := db.Where("name LIKE ?", src).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// create new user
func CreateUser(user *User) (interface{}, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// remove user by id
func DeleteUser(id int) (interface{}, error) {
	var user User
	if err := db.Where("ID = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	if err := db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// update user by id
func UpdateUser(id int, newUser User) (interface{}, error) {
	var user User
	if err := db.Where("ID = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	if err := db.Save(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

// ---------------------------------------------------------

func GetUserByUsernamePassword(username, password string) (interface{}, error) {
	var user User
	if err := db.Where("email LIKE ? AND password LIKE ?", username, password).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
