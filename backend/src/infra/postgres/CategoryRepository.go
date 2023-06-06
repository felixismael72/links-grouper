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

var _ repository.CategoryLoader = (*CategoryRepository)(nil)

type CategoryRepository struct {
	ConnectorManager
}

func (repo CategoryRepository) CreateCategory(category domain.Category) (*uuid.UUID, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	query := `insert into category(name) values ($1) returning id;`

	var categoryID uuid.UUID
	if err = conn.Get(&categoryID, query, category.Name()); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.InsertionErr)
	}

	return &categoryID, nil
}

func (repo CategoryRepository) ListCategory() ([]domain.Category, error) {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	query := `select id as category_id, name as category_name from category;`

	var categoryDTOList []models.Category
	if err = conn.Select(&categoryDTOList, query); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, errors.New(messages.FetchErr)
	}

	var categories []domain.Category
	for _, categoryDTO := range categoryDTOList {
		categories = append(categories, categoryDTO.ToDomain())
	}

	return categories, nil
}

func (repo CategoryRepository) EditCategory(category domain.Category) error {
	conn, err := repo.getConnection()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.New(messages.UnavailableDatabaseErr)
	}
	defer repo.closeConnection(conn)

	query := `update category set name = $1 where id = $2;`

	_, err = conn.Exec(query, category.Name(), category.ID())
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.New(messages.UpdateErr)
	}

	return nil
}

func NewCategoryRepository(connector ConnectorManager) *CategoryRepository {
	return &CategoryRepository{connector}
}
