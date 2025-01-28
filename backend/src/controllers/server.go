package controllers

import (
	"log"
	"net/http"
	"os"

	"osprey/src/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	echo      *echo.Echo
	DB        *mongo.Database
	scheduler *cron.Cron
}

func (s Server) Ready(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Server ok")
}

func (s Server) StartServer() error {

	err_env := godotenv.Load(".env")

	if err_env != nil {
		log.Fatal(err_env)
	}

	if err := s.echo.Start(os.Getenv("APPHOST")); err != nil {
		log.Fatal(err)
		return err
	}

	go s.scheduler.Start()

	return nil

}

func NewServerClient(db_client *mongo.Database, echo_client *echo.Echo, scheduler *cron.Cron) config.IServerClient {

	client := &Server{
		echo:      echo_client,
		DB:        db_client,
		scheduler: scheduler,
	}

	client.registerRoutes()

	return client
}

func (s *Server) registerRoutes() {

	// s.echo.Use(middleware.CORS())
	s.echo.Use(middleware.CORS())

	s.echo.GET("/readiness", s.Ready)

	ug := s.echo.Group("/Users")
	ug.GET("/:id", s.GetUser)
	ug.GET("/", s.GetUsers)
	ug.POST("/", s.AddUser)
	ug.POST("/login", s.Login)

	eg := s.echo.Group("Events")
	eg.GET("/", s.GetEvents)
	eg.POST("/Notifications", s.SetEventNotification)
	eg.GET("/Notifications/:email", s.GetEventNotifications)
	eg.DELETE("/Notifications/:CalendarId", s.RemoveEventNotification)

}
