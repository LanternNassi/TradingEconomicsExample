package controllers

import (
	"net/http"
	helpers "osprey/src/helpers"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetEvents(ctx echo.Context) error {

	events, err := helpers.FetchEventsByCountry("united states")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, events)

}
