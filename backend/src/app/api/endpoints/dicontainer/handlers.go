package dicontainer

import "grouper/src/app/api/endpoints/handlers"

func GetCategoryHandlers() *handlers.CategoryHandlers {
	return handlers.NewCategoryHandlers(GetCategoryService())
}

func GetLinkHandlers() *handlers.LinkHandlers {
	return handlers.NewLinkHandlers(GetLinkService())
}
