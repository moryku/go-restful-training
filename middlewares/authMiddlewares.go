package middlewares

import (
	"github.com/iswanulumam/go-restful-training/config"
	m "github.com/iswanulumam/go-restful-training/models"
	"github.com/labstack/echo"
)

// implementing basic auth with hardcoded username & password
func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}

// implementing basic auth with check database
func BasicAuth2(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user m.User
	if err := db.Where("email LIKE ? AND password LIKE ?", username, password).Find(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
