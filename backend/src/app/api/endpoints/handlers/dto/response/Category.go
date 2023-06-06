package response

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name,omitempty"`
}

func NewCategory(category domain.Category) *Category {
	return &Category{
		ID:   category.ID(),
		Name: category.Name(),
	}
}
