package controllers

import (
	"fmt"
	"net/http"
	helpers "osprey/src/helpers"
	models "osprey/src/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *Server) GetEvents(ctx echo.Context) error {

	events, err := helpers.FetchEventsByCountry("united states")

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, events)

}

func (s *Server) SetEventNotification(ctx echo.Context) error {

	user_event := new(models.UserEvent)

	if err := ctx.Bind(user_event); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}

	fmt.Println(user_event)

	document := bson.M{
		"email":      user_event.Email,
		"CalendarId": user_event.CalendarId,
		"Importance": user_event.Importance,
		"Event":      user_event.Event,
		"Date":       user_event.Date,
	}

	result, err := s.DB.Collection("EventsSchedule").InsertOne(ctx.Request().Context(), document)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	fmt.Println(result.InsertedID)

	return ctx.JSON(http.StatusOK, map[string]string{"Inserted": "Notification inserted"})

}

func (s *Server) RemoveEventNotification(ctx echo.Context) error {
	CalendarId := ctx.Param("CalendarId")

	filter := bson.M{"CalendarId": CalendarId}

	result, err := s.DB.Collection("EventsSchedule").DeleteOne(ctx.Request().Context(), filter)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "No matching event found"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"Deleted": "Notification deleted"})
}

func (s *Server) GetEventNotifications(ctx echo.Context) error {
	email := ctx.Param("email")

	filter := bson.M{"email": email}

	cursor, err := s.DB.Collection("EventsSchedule").Find(ctx.Request().Context(), filter)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error1": err.Error()})
	}

	type Response struct {
		CalendarId string `bson:"CalendarId"`
		Importance int    `bson:"importance"`
		Event      string `bson:"event"`
		Date       string `bson:"date"`
	}

	var notifications []Response

	if err := cursor.All(ctx.Request().Context(), &notifications); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error2": err.Error()})
	}

	return ctx.JSON(http.StatusOK, notifications)

}
