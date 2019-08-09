package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	// c.SetPath("/api/users")
	h := GetUsersController(c)

	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()

	data := `{"name":"iswanul", "email":"admin@admin.com", "password":"admin"}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	// c.SetPath("/api/users")
	h := CreateUserController(c)

	if assert.NoError(t, h) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var jsonBody interface{}

		json.NewDecoder(rec.Body).Decode(&jsonBody)
		jsonMap := jsonBody.(map[string]interface{})
		assert.Equal(t, "iswanul", jsonMap["name"].(string))
		assert.Equal(t, "admin@admin.com", jsonMap["email"].(string))
		assert.Equal(t, "admin", jsonMap["password"].(string))
	}
}
