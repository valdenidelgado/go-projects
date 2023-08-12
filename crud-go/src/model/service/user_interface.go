package service

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	LoginUserServices(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
