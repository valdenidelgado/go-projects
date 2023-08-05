package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/database/mongodb"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s\n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
