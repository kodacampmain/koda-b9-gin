package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Response map[string]any

type Body struct {
	// key datatype struct_tag
	Name string `json:"nama" form:"nama"`
	Age  int8   `json:"umur" form:"umur"`
}

type Response struct {
	Success bool
	Data    any
	Msg     string
}

type User struct {
	Email    string
	Password string
}

var Users = []User{}

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

	router.POST("/ping", func(ctx *gin.Context) {
		// Deklarasi Body
		var data Body
		// data binding
		if e := ctx.ShouldBind(&data); e != nil {
			log.Println("error", e.Error())
			// binding error
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}
		// business logic
		if data.Name == "" && data.Age == 0 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    data,
				Msg:     "empty body",
			})
			return
		}
		// success response
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Data:    data,
			Msg:     fmt.Sprintf("Selamat Datang %s", data.Name),
		})
	})

	router.Run("localhost:9000")
}
