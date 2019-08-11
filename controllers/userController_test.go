package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/iswanulumam/go-restful-training/routes"
)

var (
	newRoutes = routes.New()
	server    = httptest.NewServer(newRoutes)
)

func TestGetUsers(t *testing.T) {
	app := httpexpect.New(t, server.URL)

	// Send GET request to '/api/users'
	app.GET("/api/users").
		Expect().
		Status(http.StatusOK)
}

func TestGetUserController(t *testing.T) {
	app := httpexpect.New(t, server.URL)

	// Get user with invalid id
	app.GET("/api/users/0").
		Expect().
		Status(http.StatusBadRequest)
}
