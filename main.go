package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Response map[string]any

func main() {
	// Generate gin Engine
	router := gin.Default()

	// Deklarasi Router (endpoint & method HTTP)
	router.GET("/ping", func(c *gin.Context) {
		// Send Response
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	router.Run("localhost:9000")
}
