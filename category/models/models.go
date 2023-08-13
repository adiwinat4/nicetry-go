package models

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CreateNew struct {
	Name string `json:"name" validate:"required"`
}

type CategoryRequest struct {
	Id   int    `json:"id" validate:"required, integer"`
	Name string `json:"name" validate:"required, string"`
}
