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
	// todo: implement me

	panic("implement me")
}

func (service CategoryServices) ListCategory() ([]domain.Category, error) {
	// todo: implement me

	panic("implement me")
}

func (service CategoryServices) EditCategory(category domain.Category) error {
	// todo: implement me

	panic("implement me")
}

func NewCategoryServices(repo repository.CategoryLoader) *CategoryServices {
	return &CategoryServices{categoryRepository: repo}
}
