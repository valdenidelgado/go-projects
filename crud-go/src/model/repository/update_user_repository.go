package repository

import (
	"context"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

func (r *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser repository", zap.String("journey", "UpdateUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := r.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to call UpdateUser", err,
			zap.String("journey", "UpdateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateUser repository successfully",
		zap.String("id", userId),
		zap.String("journey", "UpdateUser"))

	return nil
}
