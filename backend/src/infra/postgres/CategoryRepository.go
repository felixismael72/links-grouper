package postgres

import (
	"grouper/src/core/domain"
	"grouper/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ repository.CategoryLoader = (*CategoryRepository)(nil)

type CategoryRepository struct {
	ConnectorManager
}

func (repo CategoryRepository) CreateCategory(category domain.Category) (*uuid.UUID, error) {
	// todo: implement me

	panic("implement me")
}

func (repo CategoryRepository) ListCategory() ([]domain.Category, error) {
	// todo: implement me

	panic("implement me")
}

func (repo CategoryRepository) EditCategory(category domain.Category) error {
	// todo: implement me

	panic("implement me")
}

func NewCategoryRepository(connector ConnectorManager) *CategoryRepository {
	return &CategoryRepository{connector}
}
