package postgres

import (
	"grouper/src/core/domain"
	"grouper/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ repository.LinkLoader = (*LinkRepository)(nil)

type LinkRepository struct {
	ConnectorManager
}

func (repo LinkRepository) RegisterLink(link domain.Link) (*uuid.UUID, error) {
	// todo: implement me

	panic("implement me")
}

func (repo LinkRepository) ListLink() ([]domain.Link, error) {
	// todo: implement me

	panic("implement me")
}

func (repo LinkRepository) EditLink(link domain.Link) error {
	// todo: implement me

	panic("implement me")
}

func NewLinkRepository(connector ConnectorManager) *LinkRepository {
	return &LinkRepository{connector}
}
