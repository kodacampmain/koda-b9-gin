package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

type MoviesHandler struct {
	ms *service.MoviesService
}

func NewMoviesHandler(ms *service.MoviesService) *MoviesHandler {
	return &MoviesHandler{
		ms: ms,
	}
}

func (m *MoviesHandler) AddMovie(ctx *gin.Context) {
	var body dto.AddMovies
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan server",
		})
		return
	}

	if err := m.ms.AddMovie(ctx.Request.Context(), body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan server",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Msg:     "Movies Created",
	})
}
