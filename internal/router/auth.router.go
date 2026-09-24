package router

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/handler"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
	"github.com/kodacampmain/koda-b9-gin/pkg"
)

func initAuthRouter(r *gin.Engine, db *pgxpool.Pool) {
	authRouter := r.Group("/auth")

	ar := repo.NewAuthRepo(db)
	as := service.NewAuthService(ar)
	ah := handler.NewAuthHandler(as)

	authRouter.POST("pwd", func(ctx *gin.Context) {
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

	authRouter.POST("compare", func(ctx *gin.Context) {
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

	authRouter.POST("new", ah.Register)
	authRouter.POST("", ah.Login)
}
