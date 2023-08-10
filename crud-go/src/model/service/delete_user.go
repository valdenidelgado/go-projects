package service

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Deleting a user model", zap.String("journey", "deleteUser"))

	err := u.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call DeleteUser repository", err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("User model deleted successfully",
		zap.String("id", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
