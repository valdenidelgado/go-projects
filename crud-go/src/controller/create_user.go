package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/validation"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller/model/request"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Creating a new user",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while trying to bind user request", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		logger.Error("Error while trying to create a new user", err,
			zap.String("journey", "createUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)
	c.JSON(200, gin.H{"message": "User created successfully"})
}
