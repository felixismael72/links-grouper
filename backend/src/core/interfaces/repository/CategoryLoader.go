package repository

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type CategoryLoader interface {
	CreateCategory(category domain.Category) (*uuid.UUID, error)
	ListCategory() ([]domain.Category, error)
	EditCategory(category domain.Category) error
}
