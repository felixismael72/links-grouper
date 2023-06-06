package repository

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type LinkLoader interface {
	RegisterLink(link domain.Link) (*uuid.UUID, error)
	ListLink() ([]domain.Link, error)
	EditLink(link domain.Link) error
}
