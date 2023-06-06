package handlers

import (
	"grouper/src/core/interfaces/primary"

	"github.com/labstack/echo/v4"
)

type LinkHandlers struct {
	linkService primary.LinkManager
}

func (handler LinkHandlers) PostLink(c echo.Context) error {
	// todo: implement me

	panic("implement me")
}

func (handler LinkHandlers) GetLink(c echo.Context) error {
	// todo: implement me

	panic("implement me")
}

func (handler LinkHandlers) PutLink(c echo.Context) error {
	// todo: implement me

	panic("implement me")
}

func NewLinkHandlers(service primary.LinkManager) *LinkHandlers {
	return &LinkHandlers{linkService: service}
}
