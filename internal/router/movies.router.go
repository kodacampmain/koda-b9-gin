package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/handler"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
	"github.com/kodacampmain/koda-b9-gin/internal/service"
)

func initMoviesRouter(r *gin.Engine, db *pgxpool.Pool) {
	moviesRouter := r.Group("movies")

	mr := repo.NewMoviesRepo()
	ms := service.NewMoviesService(mr, db)
	mh := handler.NewMoviesHandler(ms)

	moviesRouter.POST("new", mh.AddMovie)
}
