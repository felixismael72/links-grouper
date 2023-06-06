package models

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type Link struct {
	ID      uuid.UUID `db:"link_id"`
	Content string    `db:"link_content"`
	Category
}

func (dto Link) ToDomain() domain.Link {
	linkBuilder := domain.NewLinkBuilder()
	linkBuilder.WithID(dto.ID).WithContent(dto.Content)
	return *linkBuilder.Build()
}
