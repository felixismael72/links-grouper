package models

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `db:"category_id"`
	Name string    `db:"category_name"`
}

func (dto Category) ToDomain() domain.Category {
	categoryBuilder := domain.NewCategoryBuilder()
	categoryBuilder.WithID(dto.ID).WithName(dto.Name)
	return *categoryBuilder.Build()
}
