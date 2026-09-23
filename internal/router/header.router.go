package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b9-gin/internal/handler"
	"github.com/kodacampmain/koda-b9-gin/internal/middleware"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

func initHeaderRouter(r *gin.Engine) {
	headerRouter := r.Group("/header")

	hs := service.NewHeaderService()
	hh := handler.NewHeaderHandler(hs)

	headerRouter.GET("", hh.ProcessHeader)

	headerRouter.GET("data/:id/:slug", hh.ProcessParam)

	headerRouter.GET("query", middleware.M2, middleware.M3, hh.ProcessQuery)
}
