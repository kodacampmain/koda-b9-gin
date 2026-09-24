package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/middleware"
)

func InitMainRouter(router *gin.Engine, db *pgxpool.Pool) {
	// global middleware
	// router.Use(middleware.M2, middleware.M1, middleware.M3)
	router.Use(middleware.Cors)

	initPingRouter(router)
	initHeaderRouter(router)
	initPersonRouter(router, db)
	initAuthRouter(router)
}
