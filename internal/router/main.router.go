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
	_ "github.com/kodacampmain/koda-b9-gin/docs"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/middleware"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitMainRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	// global middleware
	// router.Use(middleware.M2, middleware.M1, middleware.M3)
	router.Use(middleware.Cors)

	router.Static("img", path.Join("public", "img"))
	// router.Static("docs", path.Join("public", "docs"))

	router.GET("documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	initPingRouter(router)
	initHeaderRouter(router)
	initPersonRouter(router, db, rdb)
	initAuthRouter(router, db)
	initMoviesRouter(router, db)

	router.PATCH("edit", editUser)
}

// Edit User
//
// @Summary			Update user info
// @Description		Update user info with name and image
// @Tags			user
// @Accept			mpfd
// @Produce			json
// @Router			/edit	[patch]
// @Param			name	formData	string	true	"name to update user"
// @Param			image	formData	file	true	"image to update user"
// @Success			200		{object}	dto.Response
// @Failure			400		{object}	dto.ErrorResponse
// @Failure			500		{object}	dto.ErrorResponse
func editUser(ctx *gin.Context) {
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
}
