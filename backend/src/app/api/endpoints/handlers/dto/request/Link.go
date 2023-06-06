package request

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type Link struct {
	Content     string      `json:"content"`
	CategoryIDs []uuid.UUID `json:"category_ids"`
}

func (dto Link) ToDomain(id ...uuid.UUID) *domain.Link {
	linkBuilder := domain.NewLinkBuilder()
	if len(id) > 0 {
		linkBuilder.WithID(id[0])
	}
	linkBuilder.WithContent(dto.Content)
	categoryBuilder := domain.NewCategoryBuilder()

	for _, categoryID := range dto.CategoryIDs {
		categoryBuilder.WithID(categoryID)
		linkBuilder.WithCategory(*categoryBuilder.Build())
	}
	return linkBuilder.Build()
}
