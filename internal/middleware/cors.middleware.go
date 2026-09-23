package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	// simple cors
	allowedOrigins := []string{"http://localhost:5501", "http://localhost:5500"}
	if slices.Contains(allowedOrigins, c.GetHeader("Origin")) {
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	}
	c.Header("Access-Control-Allow-Headers", "Content-Type, XXX-Header")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, PATCH")
	// c.Header("Access-Control-Max-Age", "86500")
	// log.Println("\033[31m", c.Request.URL, "\033[0m")

	// preflight cors
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
