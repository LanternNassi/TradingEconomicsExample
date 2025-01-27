package main

import (
	"context"
	"log"
	"os"
	"osprey/src/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	defer client.Disconnect(context.TODO())

	server := controllers.NewServerClient(client.Database("TradingEconomics"), echo.New())

	//start the server
	server_err := server.StartServer()

	if server_err != nil {
		log.Fatal(server_err)
		return
	}

}
