package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func FindAllUser() (interface{}, error) {
	var (
		users []User
	)
	if err := db.Find(&users).Error; err != nil {
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

func DeleteUserByID(ID int) {
	var user User
	db.Where("ID = ?", ID).Find(&user)
	db.Delete(&user)
}

func UpdateUserByID(ID int, name string, email string) {
	var user User
	db.Where("ID = ?", ID).Find(&user)
	user.Name = name
	user.Email = email
	db.Save(&user)
}
