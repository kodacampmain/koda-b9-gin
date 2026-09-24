package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	apperror "github.com/kodacampmain/koda-b9-gin/internal/error"
	"github.com/kodacampmain/koda-b9-gin/internal/model"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
	"github.com/kodacampmain/koda-b9-gin/pkg"
)

type AuthService struct {
	ar *repo.AuthRepo
}

func NewAuthService(ar *repo.AuthRepo) *AuthService {
	return &AuthService{
		ar: ar,
	}
}

func (a *AuthService) CreateNew(ctx context.Context, body dto.Account) error {
	// validasi
	if len(body.Username) == 0 || len(body.Password) == 0 {
		return errors.New("username atau password tidak boleh kosong")
	}
	// cek exists di repo
	_, err := a.ar.FindAccount(ctx, body.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if err == nil {
		return apperror.ErrAlreadyExists
	}
	// hash pwd
	hc := pkg.NewRecommendedHashConfig()
	hashedPwd := hc.GenHash(body.Password)
	// menambah data ke db
	if err := a.ar.NewAccount(ctx, model.Account{
		Username: body.Username,
		Password: hashedPwd,
	}); err != nil {
		return err
	}
	return nil
}

func (a *AuthService) Login(ctx context.Context, body dto.Account) (string, error) {
	// validasi
	if len(body.Username) == 0 || len(body.Password) == 0 {
		return "", apperror.ErrEmptyUsernamePassword
	}
	// cek account apakah sudah ada di db (melalui repo)
	acc, err := a.ar.FindAccount(ctx, body.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", apperror.ErrInvalidUsernamePassword
		}
		return "", err
	}
	if err := pkg.Compare(body.Password, acc.Password); err != nil {
		return "", apperror.ErrInvalidUsernamePassword
	}
	// password dan email cocok
	// generate JWT
	claims := pkg.NewJWTClaims(acc.Id, acc.Role)
	return claims.GenToken()
}
