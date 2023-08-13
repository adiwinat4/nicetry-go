package repository

import (
	"errors"
	"fmt"
	"nicetry/auth/models"

	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type AuthRepo interface {
	Login(ctx context.Context, query, username, password string) (*models.User, error)
}

type repoAuth struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepo {
	return &repoAuth{
		db: db,
	}
}

func (r *repoAuth) Login(ctx context.Context, query, username, password string) (*models.User, error) {
	var result models.User
	if err := r.db.QueryRowContext(ctx, query, username, password).Scan(&result.Id, &result.Username, &result.Password, &result.RoleId); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("username password no match")
		}
		return nil, err
	}
	fmt.Println(result)

	return &result, nil
}
