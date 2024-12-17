package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		c.JSON(200, gin.H{
			"message":   "Hello World!",
			"served_by": hostname,
		})
	})

	w := performRequest(router, "GET", "/")

	assert.Equal(t, 200, w.Code)

	expectedMessage := "Hello World!"
	expectedServedBy := os.Getenv("HOSTNAME")
	if expectedServedBy == "" {
		expectedServedBy = "unknown"
	}

	assert.Contains(t, w.Body.String(), expectedMessage)
}

func performRequest(router http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
