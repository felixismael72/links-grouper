package handlers

import (
	"net/http"

	"grouper/src/app/api/endpoints/handlers/dto/request"
	"grouper/src/app/api/endpoints/handlers/dto/response"
	"grouper/src/core/interfaces/primary"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandlers struct {
	categoryService primary.CategoryManager
}

func (handler CategoryHandlers) PostCategory(c echo.Context) error {
	var categoryDTO request.Category
	if err := c.Bind(&categoryDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{ErrMsg: err.Error()})
	}

	category := categoryDTO.ToDomain()

	categoryID, err := handler.categoryService.CreateCategory(*category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{ErrMsg: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Created{ID: *categoryID})
}

func (handler CategoryHandlers) GetCategory(c echo.Context) error {
	categories, err := handler.categoryService.ListCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{ErrMsg: err.Error()})
	}

	var categoryDTOList []response.Category
	for _, category := range categories {
		categoryDTOList = append(categoryDTOList, *response.NewCategory(category))
	}

	return c.JSON(http.StatusOK, categoryDTOList)
}

func (handler CategoryHandlers) PutCategory(c echo.Context) error {
	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{ErrMsg: err.Error()})
	}

	var categoryDTO request.Category
	if err = c.Bind(&categoryDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{ErrMsg: err.Error()})
	}

	category := categoryDTO.ToDomain(categoryID)

	err = handler.categoryService.EditCategory(*category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{ErrMsg: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func NewCategoryHandlers(service primary.CategoryManager) *CategoryHandlers {
	return &CategoryHandlers{categoryService: service}
}
