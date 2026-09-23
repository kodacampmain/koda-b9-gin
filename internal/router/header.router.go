package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
)

func initHeaderRouter(r *gin.Engine) {
	headerRouter := r.Group("/header")

	headerRouter.GET("", func(ctx *gin.Context) {
		// get header from request
		auth := ctx.GetHeader("Authorization")
		ctype := ctx.GetHeader("Content-Type")
		cookies := ctx.GetHeader("Cookie")
		cookie, _ := ctx.Cookie("ccc")
		uagent := ctx.GetHeader("User-Agent")
		custom := ctx.GetHeader("XXX-Header")

		// set header in response
		ctx.Header("my-header", "koda")
		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"auth":    auth,
				"ctype":   ctype,
				"cookies": cookies,
				"cookie":  cookie,
				"uagent":  uagent,
				"custom":  custom,
			},
		})
	})

	headerRouter.GET("data/:id/:slug", func(ctx *gin.Context) {
		id := ctx.Param("id")
		slug := ctx.Param("slug")

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"id":   id,
				"slug": slug,
			},
		})
	})

	headerRouter.GET("query", func(ctx *gin.Context) {
		title := ctx.Query("title")
		genres := ctx.QueryArray("genre")

		var qp dto.HeaderQuery
		if err := ctx.ShouldBindWith(&qp, binding.Query); err != nil {
			log.Println("error", err.Error())
			// binding error
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		// call service for validation

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"title":  title,
				"genres": genres,
				"qp":     qp,
			},
		})
	})
}
