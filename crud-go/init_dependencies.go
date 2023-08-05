package main

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/controller"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/repository"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)
	return userController
}
