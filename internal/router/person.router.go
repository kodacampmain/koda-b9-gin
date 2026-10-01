package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/handler"
	"github.com/kodacampmain/koda-b9-gin/internal/middleware"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
	"github.com/redis/go-redis/v9"
)

func initPersonRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	personRouter := r.Group("/person")

	pr := repo.NewPersonRepo(db)
	ps := service.NewPersonService(pr, rdb)
	ph := handler.NewPersonHandler(ps)

	personRouter.GET("", middleware.CheckToken, middleware.AdminOnly, ph.GetAllPerson)
}
