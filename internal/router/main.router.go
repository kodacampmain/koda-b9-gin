package router

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/middleware"
)

func InitMainRouter(router *gin.Engine, db *pgxpool.Pool) {
	// global middleware
	// router.Use(middleware.M2, middleware.M1, middleware.M3)
	router.Use(middleware.Cors)

	router.Static("img", path.Join("public", "img"))

	initPingRouter(router)
	initHeaderRouter(router)
	initPersonRouter(router, db)
	initAuthRouter(router, db)
	initMoviesRouter(router, db)

	router.PATCH("edit", func(ctx *gin.Context) {
		var body dto.EditUser
		if err := ctx.ShouldBindWith(&body, binding.FormMultipart); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Msg:     "terjadi kesalahan sistem",
			})
			return
		}

		log.Println("size", body.Image.Size)

		// validasi ekstensi (jpg, jpeg, png)
		// validasi ukuran (ex. max 2MB)

		filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), body.Name, path.Ext(body.Image.Filename))
		filepath := path.Join("public", "img", filename)

		if err := ctx.SaveUploadedFile(body.Image, filepath); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Msg:     "terjadi kesalahan sistem",
			})
			return
		}
		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"filepath": fmt.Sprintf("img/%s", filename),
			},
		})
	})
}
