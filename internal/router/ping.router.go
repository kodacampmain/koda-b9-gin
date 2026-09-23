package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda-b9-gin/internal/handler"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

func initPingRouter(r *gin.Engine) {
	pingRouter := r.Group("/ping")

	ps := service.NewPingService()
	// psn := service.NewPingServiceNew()
	ph := handler.NewPingHandler(ps)

	// gin.HandlerFunc
	pingRouter.GET("", ph.Pong)
	pingRouter.POST("", ph.Greet)
	// pingRouter.PATCH("")
}
