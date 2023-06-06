package dicontainer

import (
	"grouper/src/core/interfaces/primary"
	"grouper/src/core/services"
)

func GetCategoryService() primary.CategoryManager {
	return services.NewCategoryServices(GetCategoryRepository())
}

func GetLinkService() primary.LinkManager {
	return services.NewLinkServices(GetLinkRepository())
}
