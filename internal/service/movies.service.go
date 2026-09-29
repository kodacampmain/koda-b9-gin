package service

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
)

type MoviesService struct {
	mr *repo.MoviesRepo
	db *pgxpool.Pool
}

func NewMoviesService(mr *repo.MoviesRepo, db *pgxpool.Pool) *MoviesService {
	return &MoviesService{
		mr: mr,
		db: db,
	}
}

func (m *MoviesService) AddMovie(ctx context.Context, body dto.AddMovies) error {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Println(err.Error())
		}
	}()

	movieId, err := m.mr.InsertMovie(ctx, tx, body.Title)
	if err != nil {
		return err
	}

	if _, err := m.mr.InsertMovieGenre(ctx, tx, movieId, body.Genres); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
