package handlers

import (
	"grouper/src/core/interfaces/primary"

	"github.com/labstack/echo/v4"
)

type CategoryHandlers struct {
	categoryService primary.CategoryManager
}

func (handler CategoryHandlers) PostCategory(c echo.Context) error {
	// todo: implement me

	panic("implement me")
}

func (handler CategoryHandlers) GetCategory(c echo.Context) error {
	// todo: implement me

	panic("implement me")
}

func (handler CategoryHandlers) PutCategory(c echo.Context) error {
	// todo: implement me

	panic("implement me")
}

func NewCategoryHandlers(service primary.CategoryManager) *CategoryHandlers {
	return &CategoryHandlers{categoryService: service}
}
