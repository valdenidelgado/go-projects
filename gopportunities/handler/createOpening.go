package handler

import (
	"encoding/json"
	"github.com/valdenidelgado/go-projects/gopportunities/schemas"
	"net/http"
)

func CreateOpeningHandler(w http.ResponseWriter, r *http.Request) {
	request := CreateOpeningRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating request: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, opening)
}
