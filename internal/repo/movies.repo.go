package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type MoviesRepo struct{}

func NewMoviesRepo() *MoviesRepo {
	return &MoviesRepo{}
}

func (m *MoviesRepo) InsertMovie(ctx context.Context, db DBTX, title string) (int, error) {
	sql := "INSERT INTO movies (title) VALUES ($1) RETURNING id"
	args := []any{title}
	var id int
	if err := db.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MoviesRepo) InsertMovieGenre(ctx context.Context, db DBTX, movieId int, genreIds []int) (pgconn.CommandTag, error) {
	// var sqlx strings.Builder
	// sqlx.WriteString()
	sql := "INSERT INTO movies_genres (movie_id, genre_id) VALUES "
	args := []any{}
	for idx, genreId := range genreIds {
		num := (idx * 2) + 1
		sql += fmt.Sprintf("($%d, $%d)", num, num+1)
		args = append(args, movieId, genreId)
		if idx < len(genreIds)-1 {
			sql += ","
		}
	}
	return db.Exec(ctx, sql, args...)
}
