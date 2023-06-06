package primary

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type CategoryManager interface {
	CreateCategory(category domain.Category) (*uuid.UUID, error)
	ListCategory() ([]domain.Category, error)
	EditCategory(category domain.Category) error
}
