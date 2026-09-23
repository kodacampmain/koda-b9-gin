package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
)

type IPingService interface {
	EmptyValidation(data dto.User) error
}

type PingHandler struct {
	ps IPingService
}

func NewPingHandler(ps IPingService) *PingHandler {
	return &PingHandler{
		ps: ps,
	}
}

// func (p *PingHandler) Pong(ps *service.PingService) gin.HandlerFunc {
func (p *PingHandler) Pong(c *gin.Context) {
	// Send Response
	// return func(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "http://localhost:5501")
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
	// }
}

func (p *PingHandler) Greet(ctx *gin.Context) {
	// Deklarasi Body
	var data dto.User
	// data binding
	if e := ctx.ShouldBindWith(&data, binding.JSON); e != nil {
		log.Println("error", e.Error())
		// binding error
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Data:    nil,
			Msg:     "terjadi kesalahan server",
		})
		return
	}
	// business logic
	// service.NewPingService()
	if err := p.ps.EmptyValidation(data); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Data:    data,
			Msg:     err.Error(),
		})
		return
	}
	// success response
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    data,
		Msg:     fmt.Sprintf("Selamat Datang %s", data.Name),
	})
}

// func (p *PingHandler) Check(c *gin.Context) {}
