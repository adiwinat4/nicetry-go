package repository

import (
	"context"
	"database/sql"
	"nicetry/category/models"
	"nicetry/helper"

	_ "github.com/go-sql-driver/mysql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category models.Category) models.Category
	Update(ctx context.Context, tx *sql.Tx, category models.Category) models.Category
	// Delete(ctx context.Context, tx *sql.Tx, category models.Category)
	// FindById(ctx context.Context, tx *sql.Tx, categoryId int) (models.Category, error)
	// FindAll(ctx context.Context, tx *sql.Tx) []models.Category
}

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category models.Category) models.Category {
	query := "insert into categories (name) values (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category models.Category) models.Category {
	query := "update categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
