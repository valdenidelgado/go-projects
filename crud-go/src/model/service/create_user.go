package service

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Creating a new user model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	userDomainRepository, err := u.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call CreateUser repository", err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("User model created successfully",
		zap.String("id", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
