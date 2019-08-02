package controllers

import (
	"net/http"
	"net/http/httptest"
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
