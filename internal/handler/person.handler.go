package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

// GetAllPerson
//
// @Summary			Getting All Person Data
// @Description		Get all person data from database
// @Tags			person
// @Produce			json
// @Router			/person	[get]
// @Security		BearerToken
// @Param			limit	query		int		false	"limit per page"
// @Param			page	query		int		false	"page number"
// @Success			200		{object}	dto.Response
// @Failure			500		{object}	dto.ErrorResponse
func (p *PersonHandler) GetAllPerson(c *gin.Context) {
	var qp dto.PersonQuery
	if err := c.ShouldBindWith(&qp, binding.Query); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Msg:     "terjadi kesalahan server",
			Success: false,
		})
		return
	}
	// default value
	if qp.Page == 0 {
		qp.Page = 1
	}
	if qp.Limit == 0 {
		qp.Limit = 4
	}

	persons, meta, err := p.ps.GetAllPerson(c.Request.Context(), c.Request.URL.Path, qp)
	if err != nil {
		log.Println("error: ", err.Error())
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	res := dto.Response{
		Success: true,
		Msg:     "Data Persons Berhasil Diambil",
		Data:    persons,
	}
	if len(persons) > 0 {
		res.Metadata = meta
	}

	c.JSON(http.StatusOK, res)
}
