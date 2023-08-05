package repository

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{databaseConnection: database}
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
