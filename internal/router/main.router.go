package router

import (
	"github.com/gin-gonic/gin"
)

func InitMainRouter(router *gin.Engine) {

	initPingRouter(router)
	initHeaderRouter(router)
}
