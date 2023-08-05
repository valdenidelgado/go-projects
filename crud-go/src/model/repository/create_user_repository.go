package repository

import (
	"context"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (r *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository", zap.String("journey", "CreateUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := r.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to call CreateUser", err,
			zap.String("journey", "CreateUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("CreateUser repository successfully",
		zap.String("id", value.ID.Hex()),
		zap.String("journey", "CreateUser"))

	return converter.ConvertEntityToDomain(value), nil
}
