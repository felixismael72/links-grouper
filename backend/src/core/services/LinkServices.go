package services

import (
	"grouper/src/core/domain"
	"grouper/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ repository.LinkLoader = (*LinkServices)(nil)

type LinkServices struct {
	linkRepository repository.LinkLoader
}

func (service LinkServices) RegisterLink(link domain.Link) (*uuid.UUID, error) {
	// todo: implement me

	panic("implement me")
}

func (service LinkServices) ListLink() ([]domain.Link, error) {
	// todo: implement me

	panic("implement me")
}

func (service LinkServices) EditLink(link domain.Link) error {
	// todo: implement me

	panic("implement me")
}

func NewLinkServices(repo repository.LinkLoader) *LinkServices {
	return &LinkServices{linkRepository: repo}
}
