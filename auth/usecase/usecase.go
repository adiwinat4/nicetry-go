package usecase

import (
	"context"
	"errors"
	"nicetry/auth/models"
	"nicetry/auth/repository"
)

type AuthUsecase interface {
	Login(ctx context.Context, username, password string) (*models.User, error)
}

type usecaseAuth struct {
	repo repository.AuthRepo
}

func NewAuthUsecase(repo repository.AuthRepo) AuthUsecase {
	return &usecaseAuth{
		repo: repo,
	}
}

func (u *usecaseAuth) Login(ctx context.Context, username, password string) (*models.User, error) {
	query := `select id, username, password, role_id FROM users WHERE username = ? AND password = ?`
	user, err := u.repo.Login(ctx, query, username, password)
	if err != nil {
		return nil, errors.New("error aja")
	}
	return user, nil
}
