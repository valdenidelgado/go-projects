package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Deleting a new user",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error while trying to convert userId to Hex", err,
			zap.String("journey", "deleteUser"),
		)
		errRest := rest_err.NewBadRequestError("Invalid id, must be a valid hex")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call DeleteUser service", err,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully",
		zap.String("id", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)
}
