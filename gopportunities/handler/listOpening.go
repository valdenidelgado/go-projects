package handler

import (
	"github.com/valdenidelgado/go-projects/gopportunities/schemas"
	"net/http"
)

func ListOpeningHandler(w http.ResponseWriter, r *http.Request) {
	var openings []schemas.Opening
	if err := db.Find(&openings).Error; err != nil {
		sendError(w, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(w, openings)
}
