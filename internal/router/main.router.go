package router

import (
	"github.com/gin-gonic/gin"
)

func InitMainRouter(mainRouter *gin.Engine) {

	initPingRouter(mainRouter)
}
