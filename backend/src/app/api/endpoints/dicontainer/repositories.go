package dicontainer

import (
	"grouper/src/core/interfaces/repository"
	"grouper/src/infra/postgres"
)

func GetCategoryRepository() repository.CategoryLoader {
	return postgres.NewCategoryRepository(GetPGConnectionManager())
}

func GetLinkRepository() repository.LinkLoader {
	return postgres.NewLinkRepository(GetPGConnectionManager())
}

func GetPGConnectionManager() postgres.ConnectorManager {
	return &postgres.DatabaseConnectionManager{}
}
