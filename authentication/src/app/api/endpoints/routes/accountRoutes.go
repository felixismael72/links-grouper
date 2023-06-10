package routes

import (
	"authentication/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadAccountRoutes(api *echo.Group) {
	authGroup := api.Group("/auth")

	accountHandlers := dicontainer.GetAccountHandlers()

	authGroup.POST("/sign-up", accountHandlers.SignUp)
	authGroup.POST("/sign-in", accountHandlers.SignIn)
}
