package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/iswanulumam/go-restful-training/routes"
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

func TestGetUsers(t *testing.T) {
	// call all routes
	router := routes.New()
	// // create server based on router
	server := httptest.NewServer(router)
	defer server.Close()

	// crate server app
	app := httpexpect.New(t, server.URL)
	app.GET("/api/users").
		Expect().
		Status(http.StatusOK)
}
