package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func M1(c *gin.Context) {
	// run before handler/controller (alur request)
	log.Println("middleware 1")
	c.Next() // move to next middleware/handler
	// run after handler/controller (alur response)
	log.Println("middleware 1")
}

func M2(c *gin.Context) {
	log.Println("middleware 2")
	c.Next()
	log.Println("middleware 2")
}

func M3(c *gin.Context) {
	log.Println("middleware 3")
	c.Next()
	log.Println("middleware 3")
}
