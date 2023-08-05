package repository

import (
	"context"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"os"
)

var (
	MONGO_COLLECTION_NAME = "MONGO_COLLECTION_NAME"
)

func (r *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository")

	collection_name := os.Getenv(MONGO_COLLECTION_NAME)
	collection := r.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
