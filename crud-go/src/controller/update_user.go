package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/validation"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller/model/request"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Updating a new user",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest
	userId := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(userId) == "" {
		logger.Error("Error while trying to bind user request", err,
			zap.String("journey", "updateUser"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service", err,
			zap.String("journey", "updateUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully",
		zap.String("id", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
