package handler

import (
	"fmt"
	"github.com/valdenidelgado/go-projects/gopportunities/schemas"
	"net/http"
)

func DeleteOpeningHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		sendError(w, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(w, http.StatusNotFound, fmt.Sprintf("opening whit id: %s not found", id))
		return
	}

	if err := db.Delete(&opening, id).Error; err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(w, opening)
}
