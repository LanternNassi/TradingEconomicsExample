package config

import (
	"github.com/labstack/echo/v4"
)

type IServerClient interface {
	Ready(ctx echo.Context) error
	StartServer() error

	AddUser(ctx echo.Context) error
	GetUser(ctx echo.Context) error
	GetUsers(ctx echo.Context) error
	Login(ctx echo.Context) error

	GetEvents(ctx echo.Context) error
	SetEventNotification(ctx echo.Context) error
	GetEventNotifications(ctx echo.Context) error
	RemoveEventNotification(ctx echo.Context) error
}
