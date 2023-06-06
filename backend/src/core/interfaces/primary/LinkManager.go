package primary

import (
	"grouper/src/core/domain"

	"github.com/google/uuid"
)

type LinkManager interface {
	RegisterLink(link domain.Link) (*uuid.UUID, error)
	ListLink() ([]domain.Link, error)
	EditLink(link domain.Link) error
}
