package domain

import "github.com/google/uuid"

type Link struct {
	id         uuid.UUID
	content    string
	categories []Category
}

func (entity Link) ID() uuid.UUID {
	return entity.id
}

func (entity Link) Content() string {
	return entity.content
}

func (entity Link) Categories() []Category {
	return entity.categories
}
