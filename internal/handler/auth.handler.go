package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	apperror "github.com/kodacampmain/koda-b9-gin/internal/error"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

type AuthHandler struct {
	as *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var body dto.Account
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan server",
		})
		return
	}
	// gunakan service
	if err := a.as.CreateNew(c.Request.Context(), body); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan server",
		})
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Msg:     "user registered",
	})
}

func (a *AuthHandler) Login(c *gin.Context) {
	var body dto.Account
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan server",
		})
		return
	}
	// gunakan service
	token, err := a.as.Login(c.Request.Context(), body)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, apperror.ErrEmptyUsernamePassword) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Msg:     err.Error(),
			})
			return
		}
		if errors.Is(err, apperror.ErrInvalidUsernamePassword) {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Msg:     err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"token": token,
		},
	})
}
