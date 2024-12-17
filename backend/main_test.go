package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	return router
}

func TestRootRoute(t *testing.T) {
	// Setup Gin in test mode
	gin.SetMode(gin.TestMode)

	// Initialize router
	router := SetupRouter()

	// Create a test HTTP request to the root route "/"
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	// Record the response using httptest
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code) // Check status code is 200

	// Check response body
	expected := `{"message":"Hello World!"}`
	assert.JSONEq(t, expected, w.Body.String()) // Compare JSON response
}
