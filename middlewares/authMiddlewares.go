package middlewares

import (
	"github.com/iswanulumam/go-restful-training/models"
	"github.com/labstack/echo"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	_, err := models.GetUserByUsernamePassword(username, password)
	if err == nil {
		return false, nil
	}
	return true, nil
}
