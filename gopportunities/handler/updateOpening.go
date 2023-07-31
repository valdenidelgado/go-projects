package handler

import (
	"encoding/json"
	"github.com/valdenidelgado/go-projects/gopportunities/schemas"
	"net/http"
)

func UpdateOpeningHandler(w http.ResponseWriter, r *http.Request) {
	request := UpdateOpeningRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		sendError(w, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, request.ID).Error; err != nil {
		sendError(w, http.StatusNotFound, err.Error())
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating request: %v", err)
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(w, opening)
}
