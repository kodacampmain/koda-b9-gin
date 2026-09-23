package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b9-gin/internal/router"
)

// type Response map[string]any

func main() {
	// Generate gin Engine
	r := gin.Default()
	// Deklarasi Router (endpoint & method HTTP)

	router.InitMainRouter(r)

	r.Run("localhost:9000")
}
