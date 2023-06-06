package routes

import (
	"grouper/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadCategoryRoutes(apiGroup *echo.Group) {
	categoryHandlers := dicontainer.GetCategoryHandlers()

	categoryGroup := apiGroup.Group("/category")

	categoryGroup.POST("/new", categoryHandlers.PostCategory)
	categoryGroup.GET("", categoryHandlers.GetCategory)
	categoryGroup.PUT("/:categoryID/edit", categoryHandlers.PutCategory)
}
