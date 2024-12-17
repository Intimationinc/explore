package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
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

	router.Run(":8081")
}
