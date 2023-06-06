package postgres

import (
	"errors"

	"grouper/src/core/domain"
	"grouper/src/core/interfaces/repository"
	"grouper/src/core/messages"
	"grouper/src/infra/postgres/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var _ repository.LinkLoader = (*LinkRepository)(nil)

type LinkRepository struct {
	ConnectorManager
}

func (repo LinkRepository) RegisterLink(link domain.Link) (*uuid.UUID, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	transaction, err := conn.Beginx()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.TransactionErr)
	}
	defer func() { _ = transaction.Rollback() }()

	insertLink := `insert into link(content) values($1) returning id;`

	var linkID uuid.UUID
	if err = transaction.Get(&linkID, insertLink, link.Content()); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.InsertionErr)
	}

	relateLinkCategory := `insert into link_category(link_id, category_id) values($1, $2);`
	for _, category := range link.Categories() {
		_, err = transaction.Exec(relateLinkCategory, linkID, category.ID())
		if err != nil {
			log.Err(err).Msg(err.Error())
			return nil, errors.New(messages.InsertionErr)
		}
	}

	if err = transaction.Commit(); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.TransactionErr)
	}

	return &linkID, nil
}

func (repo LinkRepository) ListLink() ([]domain.Link, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	fetchLinks := `select link.id as link_id, link.content as link_content, category.id as category_id, category.name as category_name
	from link
	left join link_category on link.id = link_category.link_id
	left join category on link_category.category_id = category.id;`

	var linkDTOList []models.Link
	if err = conn.Select(&linkDTOList, fetchLinks); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.FetchErr)
	}

	linkBuilders := make(map[uuid.UUID]*domain.LinkBuilder)

	for _, linkDTO := range linkDTOList {
		linkBuilder, exists := linkBuilders[linkDTO.ID]
		if !exists {
			linkBuilder = domain.NewLinkBuilder().WithID(linkDTO.ID).WithContent(linkDTO.Content)
			linkBuilders[linkDTO.ID] = linkBuilder
		}

		category := models.Category{
			ID:   linkDTO.Category.ID,
			Name: linkDTO.Category.Name,
		}
		linkBuilder.WithCategory(category.ToDomain())
	}

	var links []domain.Link
	for _, linkBuilder := range linkBuilders {
		links = append(links, *linkBuilder.Build())
	}

	return links, nil
}

func (repo LinkRepository) EditLink(link domain.Link) error {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	transaction, err := conn.Beginx()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.New(messages.TransactionErr)
	}
	defer func() { _ = transaction.Rollback() }()

	updateLink := `update link set content = $1 where id = $2;`

	_, err = transaction.Exec(updateLink, link.Content(), link.ID())
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.New(messages.UpdateErr)
	}

	relateLinkCategory := `
		insert into link_category(link_id, category_id) 
		select $1, $2
		where not exists (
			select true from link_category where link_id = $1 and category_id = $2
		)
		on conflict do nothing;
	`
	for _, category := range link.Categories() {
		_, err = transaction.Exec(relateLinkCategory, link.ID(), category.ID())
		if err != nil {
			log.Err(err).Msg(err.Error())
			return errors.New(messages.InsertionErr)
		}
	}

	if err = transaction.Commit(); err != nil {
		log.Err(err).Msg(err.Error())
		return errors.New(messages.TransactionErr)
	}

	return nil
}

func NewLinkRepository(connector ConnectorManager) *LinkRepository {
	return &LinkRepository{connector}
}
