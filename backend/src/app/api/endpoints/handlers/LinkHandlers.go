package handlers

import (
	"net/http"

	"grouper/src/app/api/endpoints/handlers/dto/request"
	"grouper/src/app/api/endpoints/handlers/dto/response"
	"grouper/src/core/interfaces/primary"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LinkHandlers struct {
	linkService primary.LinkManager
}

func (handler LinkHandlers) PostLink(c echo.Context) error {
	var linkDTO request.Link
	if err := c.Bind(&linkDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{ErrMsg: err.Error()})
	}

	link := linkDTO.ToDomain()

	linkID, err := handler.linkService.RegisterLink(*link)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{ErrMsg: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Created{ID: *linkID})
}

func (handler LinkHandlers) GetLink(c echo.Context) error {
	links, err := handler.linkService.ListLink()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{ErrMsg: err.Error()})
	}

	var linkDTOList []response.Link
	for _, link := range links {
		linkDTOList = append(linkDTOList, *response.NewLink(link))
	}

	return c.JSON(http.StatusOK, linkDTOList)
}

func (handler LinkHandlers) PutLink(c echo.Context) error {
	linkID, err := uuid.Parse(c.Param("linkID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{ErrMsg: err.Error()})
	}

	var linkDTO request.Link
	if err = c.Bind(&linkDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{ErrMsg: err.Error()})
	}

	link := linkDTO.ToDomain(linkID)

	err = handler.linkService.EditLink(*link)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{ErrMsg: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func NewLinkHandlers(service primary.LinkManager) *LinkHandlers {
	return &LinkHandlers{linkService: service}
}
