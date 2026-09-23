package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/handler"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

func initPersonRouter(r *gin.Engine, db *pgxpool.Pool) {
	personRouter := r.Group("/person")

	pr := repo.NewPersonRepo(db)
	ps := service.NewPersonService(pr)
	ph := handler.NewPersonHandler(ps)

	personRouter.GET("", ph.GetAllPerson)
}
