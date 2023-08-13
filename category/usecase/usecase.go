package usecase

import (
	"context"
	"database/sql"
	"log"
	"nicetry/category/models"
	"nicetry/category/repository"
	"nicetry/helper"

	"github.com/go-playground/validator/v10"
)

type CategoryUsecase interface {
	Save(ctx context.Context, request *models.CreateNew) models.Category
	Update(ctx context.Context, request *models.CategoryRequest) models.Category
}

type CategoryUsecaseImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryUsecase(repo repository.CategoryRepository, db *sql.DB, val *validator.Validate) CategoryUsecase {
	return &CategoryUsecaseImpl{
		CategoryRepository: repo,
		DB:                 db,
	}
}

func (category *CategoryUsecaseImpl) Save(ctx context.Context, request *models.CreateNew) models.Category {
	log.Println(request)
	err := category.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := category.DB.Begin()
	helper.PanicIfError(err)

	newData := models.Category{Name: request.Name}
	result := category.CategoryRepository.Save(ctx, tx, newData)
	helper.PanicIfError(err)

	return result
}

func (category *CategoryUsecaseImpl) Update(ctx context.Context, request *models.CategoryRequest) models.Category {
	err := category.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := category.DB.Begin()
	helper.PanicIfError(err)

	newData := models.Category{Id: request.Id, Name: request.Name}
	result := category.CategoryRepository.Update(ctx, tx, newData)
	helper.PanicIfError(err)

	return result
}
