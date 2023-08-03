package view

import (
	"github.com/valdenidelgado/go-projects/crud-go/src/controller/model/response"
	"github.com/valdenidelgado/go-projects/crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
