package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/logger"
	"github.com/valdenidelgado/go-projects/crud-go/src/configuration/validation"
	"github.com/valdenidelgado/go-projects/crud-go/src/controller/model/request"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
	"github.com/valdenidelgado/go-projects/crud-go/src/view"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while trying to bind user request", err,
			zap.String("journey", "loginUser"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call LoginUser service", err,
			zap.String("journey", "loginUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User logged successfully",
		zap.String("id", domainResult.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
