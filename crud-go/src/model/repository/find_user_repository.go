package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/repository/entity"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

func (r *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository", zap.String("journey", "FindUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := r.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User with email %s not found", email)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository successfully",
		zap.String("journey", "FindUserByEmail"),
		zap.String("email", email),
		zap.String("id", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(userEntity), nil
}

func (r *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID repository", zap.String("journey", "FindUserByID"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := r.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User with id %s not found", id)
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository successfully",
		zap.String("journey", "FindUserByID"),
		zap.String("id", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(userEntity), nil
}

func (r *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword repository", zap.String("journey", "FindUserByEmailAndPassword"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := r.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmailAndPassword"))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage, err, zap.String("journey", "FindUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository successfully",
		zap.String("journey", "FindUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("id", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(userEntity), nil
}
