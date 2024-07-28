package main_test

import (
	"loan-management-system/config"
	"loan-management-system/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Assuming LoadConfig modifies some global vars or similar
	config.LoadConfig()
	// Verify expected config values, this is just a placeholder
	assert.NotEmpty(t, config.AppConfig.DBName, "Config value should not be empty")
}

func TestRouteSetup(t *testing.T) {

	gin.SetMode(gin.TestMode)

	// Create a new router
	r := gin.Default()
	routes.AuthRoutes(r)

	// Create a new HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	// Perform the request
	r.ServeHTTP(w, req)

	// Check if the response status code is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Check if the response body is as expected
	expectedBody := `{"status":200,"message":"UP"}`
	assert.JSONEq(t, expectedBody, w.Body.String(), "Response body should match")
}
