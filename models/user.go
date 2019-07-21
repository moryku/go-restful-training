package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func GetAllUser() (interface{}, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserLike(match string) (interface{}, error) {
	var users []User
	src := "%" + match + "%"
	if err := db.Where("name LIKE ?", src).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(m *User) (*User, error) {
	if err := db.Create(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func DeleteUserById(id int) (*User, error) {
	var user User
	db.Where("ID = ?", id).Find(&user)
	if err := db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserById(id int, name string, email string) (*User, error) {
	var user User
	db.Where("ID = ?", id).Find(&user)
	user.Name = name
	user.Email = email
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
