package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b9-gin/internal/middleware"
)

func InitMainRouter(router *gin.Engine) {
	// global middleware
	// router.Use(middleware.M2, middleware.M1, middleware.M3)
	router.Use(middleware.Cors)

	initPingRouter(router)
	initHeaderRouter(router)
}
