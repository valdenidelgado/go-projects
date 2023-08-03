package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/database/mongodb"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller/routes"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/service"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
