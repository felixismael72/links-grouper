package routes

import "github.com/labstack/echo/v4"

func loadLinkRoutes(apiGroup *echo.Group) {
	_ = apiGroup.Group("/link")
}
