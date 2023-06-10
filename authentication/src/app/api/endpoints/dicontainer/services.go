package dicontainer

import (
	"authentication/src/core/interfaces/primary"
	"authentication/src/core/services"
)

func GetAccountServices() primary.AccountManager {
	return services.NewAccountServices(GetAccountRepository())
}
