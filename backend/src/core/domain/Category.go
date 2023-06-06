package domain

import "github.com/google/uuid"

type Category struct {
	id   uuid.UUID
	name string
}

func (entity Category) ID() uuid.UUID {
	return entity.id
}

func (entity Category) Name() string {
	return entity.name
}
