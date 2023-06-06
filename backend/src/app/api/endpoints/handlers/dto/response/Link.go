package response

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type Link struct {
	ID         uuid.UUID  `json:"id"`
	Content    string     `json:"content"`
	Categories []Category `json:"categories,omitempty"`
}

func NewLink(link domain.Link) *Link {
	var categories []Category
	for _, category := range link.Categories() {
		categories = append(categories, *NewCategory(category))
	}

	return &Link{
		ID:         link.ID(),
		Content:    link.Content(),
		Categories: categories,
	}
}
