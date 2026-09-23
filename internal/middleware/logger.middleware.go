package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	fmt.Printf("%ds", duration)
}
