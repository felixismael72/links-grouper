package dicontainer

import (
	"authentication/src/core/interfaces/repository"
	"authentication/src/infra/postgres"
)

func GetAccountRepository() repository.AccountLoader {
	return postgres.NewAccountRepository(GetPGConnectionManager())
}

func GetPGConnectionManager() postgres.ConnectorManager {
	return postgres.DatabaseConnectionManager{}
}
