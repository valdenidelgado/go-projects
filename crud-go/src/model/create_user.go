package model

import (
	"fmt"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Creating a new user", zap.String("journey", "createUser"))
	u.EncryptPassword()
	fmt.Println(u)
	panic("implement me")
}
