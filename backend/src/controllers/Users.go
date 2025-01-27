package controllers

import (
	"net/http"
	"osprey/src/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *Server) AddUser(ctx echo.Context) error {

	user := new(models.User)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	coll, err := s.DB.Collection("Users").InsertOne(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, coll)

}

func (s *Server) GetUsers(ctx echo.Context) error {

	user := new(models.User)

	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	cursor, err := s.DB.Collection("Users").Find(ctx.Request().Context(), bson.M{})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	defer cursor.Close(ctx.Request().Context())

	var users []models.User

	if err := cursor.All(ctx.Request().Context(), &users); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to get users"})
	}

	return ctx.JSON(http.StatusOK, users)

}

func (s *Server) GetUser(ctx echo.Context) error {

	var user models.User

	err := s.DB.Collection("Users").FindOne(ctx.Request().Context(), bson.M{"id": ctx.Param("id")}).Decode(user)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)

}

func (s *Server) Login(ctx echo.Context) error {

	type Login struct {
		username string
		password string
	}

	var LoginCredentials Login

	if err := ctx.Bind(&LoginCredentials); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}

	var user models.User

	filter := bson.M{}

	filter["username"] = LoginCredentials.username
	filter["password"] = LoginCredentials.password

	err := s.DB.Collection("Users").FindOne(ctx.Request().Context(), filter).Decode(user)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}

	return ctx.JSON(http.StatusAccepted, map[string]string{"Status": "Authorized"})
}
