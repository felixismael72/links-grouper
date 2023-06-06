package routes

import "github.com/labstack/echo/v4"

func Load(router *echo.Echo) {
	apiGroup := router.Group("/api")

	loadCategoryRoutes(apiGroup)
	loadLinkRoutes(apiGroup)
}
