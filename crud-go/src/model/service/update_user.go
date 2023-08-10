package service

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Updating a new user model", zap.String("journey", "updateUser"))

	err := u.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser repository", err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("User model updated successfully",
		zap.String("id", userId),
		zap.String("journey", "updateUser"))

	return nil
}
