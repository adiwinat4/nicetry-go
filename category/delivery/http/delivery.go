package http

import (
	"log"
	"net/http"
	"nicetry/category/models"
	"nicetry/category/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryDelivery struct {
	CategoryUsecase usecase.CategoryUsecase
}

func NewCategoryDelivery(uc usecase.CategoryUsecase) *CategoryDelivery {
	return &CategoryDelivery{
		CategoryUsecase: uc,
	}
}

func (category *CategoryDelivery) Save(c *fiber.Ctx) error {
	data := new(models.CreateNew)
	if err := c.BodyParser(data); err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	newData := models.CreateNew{Name: data.Name}

	res := category.CategoryUsecase.Save(c.Context(), &newData)
	log.Println(res)
	return nil
}

func (category *CategoryDelivery) Update(c *fiber.Ctx) error {
	data := models.CategoryRequest{}
	if err := c.BodyParser(data); err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	res := category.CategoryUsecase.Update(c.Context(), &data)
	log.Println(res)
	return nil
}
