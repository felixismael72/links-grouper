package dicontainer

import "authentication/src/app/api/endpoints/handlers"

func GetAccountHandlers() *handlers.AccountHandlers {
	return handlers.NewAccountHandlers(GetAccountServices())
}
