package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

type PersonHandler struct {
	ps *service.PersonService
}

func NewPersonHandler(ps *service.PersonService) *PersonHandler {
	return &PersonHandler{
		ps: ps,
	}
}

func (p *PersonHandler) GetAllPerson(c *gin.Context) {
	persons, err := p.ps.GetAllPerson(c.Request.Context())
	if err != nil {
		log.Println("error: ", err.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Msg:     "Data Persons Berhasil Diambil",
		Data:    persons,
	})
}
