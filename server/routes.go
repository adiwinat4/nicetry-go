package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	categoryDelivery "nicetry/category/delivery/http"
	categoryRepository "nicetry/category/repository"
	categoryUsecase "nicetry/category/usecase"

	authDelivery "nicetry/auth/delivery/http"
	authRepository "nicetry/auth/repository"
	authUsecase "nicetry/auth/usecase"
)

type Config struct {
	JWT_PRIV string
	APP_PORT string
}

func SetupRoutes(app *fiber.App, c Config) {
	db := NewDBConnection()
	val := validator.New()

	categoryRepo := categoryRepository.NewCategoryRepository()
	categoryUC := categoryUsecase.NewCategoryUsecase(categoryRepo, db, val)
	categoryHandler := categoryDelivery.NewCategoryDelivery(categoryUC)

	authRepo := authRepository.NewAuthRepo(db)
	authUC := authUsecase.NewAuthUsecase(authRepo)
	authHandler := authDelivery.NewAuthDelivery(c.JWT_PRIV, authUC)

	api := app.Group("/api", logger.New())

	api.Post("/category", categoryHandler.Save)
	api.Put("/category", categoryHandler.Update)

	api.Post("/login", authHandler.Login)
	api.Get("/me", authHandler.Me)
}
