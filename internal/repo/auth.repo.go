package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	apperror "github.com/kodacampmain/koda-b9-gin/internal/error"
	"github.com/kodacampmain/koda-b9-gin/internal/model"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (a *AuthRepo) FindAccount(ctx context.Context, username string) (model.Account, error) {
	sql := "SELECT id, role, password FROM accounts WHERE username=$1"
	args := []any{username}
	var data model.Account
	if err := a.db.QueryRow(ctx, sql, args...).Scan(&data.Id, &data.Role, &data.Password); err != nil {
		return model.Account{}, err
	}
	return data, nil
}

func (a *AuthRepo) NewAccount(ctx context.Context, body model.Account) error {
	sql := "INSERT INTO accounts (username, password) VALUES ($1, $2)"
	args := []any{body.Username, body.Password}
	cmd, err := a.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return apperror.ErrNoRowAffected
	}
	return nil
}
