package router

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/pkg"
)

func initAuthRouter(r *gin.Engine) {
	authRouter := r.Group("/auth")

	authRouter.POST("/pwd", func(ctx *gin.Context) {
		type body struct {
			Password string `json:"pwd"`
		}
		var reqBody body
		if err := ctx.ShouldBindWith(&reqBody, binding.JSON); err != nil {
			log.Println("error", err.Error())
			// binding error
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		hc := pkg.NewRecommendedHashConfig()
		hash := hc.GenHash(reqBody.Password)

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"pwd":  reqBody.Password,
				"hash": hash,
			},
		})
	})

	authRouter.POST("/compare", func(ctx *gin.Context) {
		type body struct {
			Password string `json:"pwd"`
			Hash     string `json:"hash"`
		}
		var reqBody body
		if err := ctx.ShouldBindWith(&reqBody, binding.JSON); err != nil {
			log.Println("error", err.Error())
			// binding error
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		err := pkg.Compare(reqBody.Password, reqBody.Hash)
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, pkg.ErrMismatchHash) {
				ctx.JSON(http.StatusUnauthorized, dto.Response{
					Success: false,
					Msg:     "password salah",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Msg:     "password betul",
		})
	})
}
