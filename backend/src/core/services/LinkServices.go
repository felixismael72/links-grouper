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
	linkID, err := service.linkRepository.RegisterLink(link)
	if err != nil {
		return nil, err
	}

	return linkID, nil
}

func (service LinkServices) ListLink() ([]domain.Link, error) {
	links, err := service.linkRepository.ListLink()
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (service LinkServices) EditLink(link domain.Link) error {
	err := service.linkRepository.EditLink(link)
	if err != nil {
		return err
	}

	return nil
}

func NewLinkServices(repo repository.LinkLoader) *LinkServices {
	return &LinkServices{linkRepository: repo}
}
