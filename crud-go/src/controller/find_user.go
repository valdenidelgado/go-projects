package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"github.com/valdenidelgado/go-projects/crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/mail"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserById controller.", zap.String("journey", "findUserById"))
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to parse userId", err, zap.String("journey", "findUserById"))
		errorMessage := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call FindUserByIDServices", err, zap.String("journey", "findUserById"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User model found successfully")

	c.JSON(200, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller.", zap.String("journey", "findUserByEmail"))
	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to parse userEmail", err, zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("Invalid user email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmailServices", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User model found successfully")
	c.JSON(200, view.ConvertDomainToResponse(
		userDomain,
	))
}
