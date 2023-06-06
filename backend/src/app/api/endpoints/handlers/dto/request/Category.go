package request

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type Category struct {
	Name string `json:"name"`
}

func (dto Category) ToDomain(id ...uuid.UUID) *domain.Category {
	categoryBuilder := domain.NewCategoryBuilder()
	if len(id) > 0 {
		categoryBuilder.WithID(id[0])
	}
	categoryBuilder.WithName(dto.Name)
	return categoryBuilder.Build()
}
