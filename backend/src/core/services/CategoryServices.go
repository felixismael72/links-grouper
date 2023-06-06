package services

import (
	"grouper/src/core/domain"
	"grouper/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ repository.CategoryLoader = (*CategoryServices)(nil)

type CategoryServices struct {
	categoryRepository repository.CategoryLoader
}

func (service CategoryServices) CreateCategory(category domain.Category) (*uuid.UUID, error) {
	categoryID, err := service.categoryRepository.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return categoryID, nil
}

func (service CategoryServices) ListCategory() ([]domain.Category, error) {
	categories, err := service.categoryRepository.ListCategory()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (service CategoryServices) EditCategory(category domain.Category) error {
	err := service.categoryRepository.EditCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func NewCategoryServices(repo repository.CategoryLoader) *CategoryServices {
	return &CategoryServices{categoryRepository: repo}
}
