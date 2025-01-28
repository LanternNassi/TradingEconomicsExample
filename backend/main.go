package main

import (
	"context"
	"log"
	"os"
	"osprey/src/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongo_uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(options.Client().
		ApplyURI(mongo_uri))

	if err != nil {
		panic(err)
	}

	scheduler := cron.New()

	defer client.Disconnect(context.TODO())

	defer scheduler.Stop()

	server := controllers.NewServerClient(client.Database("TradingEconomics"), echo.New(), scheduler)

	//start the server
	server_err := server.StartServer()

	if server_err != nil {
		log.Fatal(server_err)
		return
	}

}
