package handlers

import (
	"net/http"

	"authentication/src/app/api/endpoints/handlers/dto/request"
	"authentication/src/app/api/endpoints/handlers/dto/response"
	"authentication/src/core/interfaces/primary"
	"authentication/src/core/messages"

	"github.com/labstack/echo/v4"
)

type AccountHandlers struct {
	accountService primary.AccountManager
}

func (handler AccountHandlers) SignUp(c echo.Context) error {
	var dto request.SignUp
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(messages.BadRequestErr))
	}

	account, err := dto.ToDomain()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(messages.BadRequestErr))
	}

	accessToken, err := handler.accountService.SignUp(*account)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.NewError(err.Error()))
	}

	return c.JSON(http.StatusCreated, response.NewSession(accessToken))
}

func (handler AccountHandlers) SignIn(c echo.Context) error {
	var dto request.SignIn
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(messages.BadRequestErr))
	}

	account, err := dto.ToDomain()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(messages.BadRequestErr))
	}

	accessToken, err := handler.accountService.SignIn(*account)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.NewError(err.Error()))
	}

	return c.JSON(http.StatusCreated, response.NewSession(accessToken))
}

func NewAccountHandlers(accountService primary.AccountManager) *AccountHandlers {
	return &AccountHandlers{accountService: accountService}
}
