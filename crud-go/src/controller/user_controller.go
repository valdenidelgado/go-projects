package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/model/service"
)

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	LoginUser(c *gin.Context)

	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}
