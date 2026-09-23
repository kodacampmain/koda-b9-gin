package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b9-gin/internal/handler"
	"github.com/kodacampmain/koda-b9-gin/internal/middleware"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

func initPingRouter(r *gin.Engine) {
	pingRouter := r.Group("/ping")

	ps := service.NewPingService()
	// psn := service.NewPingServiceNew()
	ph := handler.NewPingHandler(ps)

	// route group middleware
	pingRouter.Use(middleware.M1, middleware.M2)
	// gin.HandlerFunc
	// route middleware
	// router.GET("", mid1, mid2, handler)
	pingRouter.GET("", ph.Pong)
	pingRouter.PATCH("", ph.Greet)
	// pingRouter.PATCH("")
}
