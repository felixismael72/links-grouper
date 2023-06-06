package routes

import (
	"grouper/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadLinkRoutes(apiGroup *echo.Group) {
	linkHandlers := dicontainer.GetLinkHandlers()

	linkGroup := apiGroup.Group("/link")

	linkGroup.POST("/new", linkHandlers.PostLink)
	linkGroup.GET("", linkHandlers.GetLink)
	linkGroup.PUT("/:linkID/edit", linkHandlers.PutLink)
}
