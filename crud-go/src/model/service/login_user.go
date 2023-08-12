package service

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init loginUser service", zap.String("journey", "loginUser"))
	userDomain.EncryptPassword()

	user, err := u.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, err
	}

	logger.Info("Login user service executed successfully",
		zap.String("id", user.GetID()),
		zap.String("journey", "loginUser"))

	return user, nil
}
