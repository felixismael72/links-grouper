package routes

import "github.com/labstack/echo/v4"

func loadCategoryRoutes(apiGroup *echo.Group) {
	_ = apiGroup.Group("/category")
}
