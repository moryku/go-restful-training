package users_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/iswanulumam/go-restful-training/routes"
)

func TestGetUsers(t *testing.T) {
	// call all routes
	router := routes.New()
	// create server based on router
	server := httptest.NewServer(router)
	defer server.Close()

	// crate server app
	app := httpexpect.New(t, server.URL)
	app.GET("/api/users").
		Expect().
		Status(http.StatusOK)
}
