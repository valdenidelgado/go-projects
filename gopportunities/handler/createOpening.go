package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

func CreateOpeningHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"msg": "hello"})
}
