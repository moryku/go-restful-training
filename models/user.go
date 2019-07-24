package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func GetUsers() (interface{}, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsersLike(match string) (interface{}, error) {
	var users []User
	src := "%" + match + "%"
	if err := db.Where("name LIKE ?", src).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *User) (interface{}, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user User
	db.Where("ID = ?", id).Find(&user)
	if err := db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id int, newUser User) (interface{}, error) {
	var user User
	db.Where("ID = ?", id).Find(&user)
	if err := db.Save(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}
