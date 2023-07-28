package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

func CreateOpeningHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"msg": "hello"})
}

func ShowOpeningHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"msg": "hello"})

}

func DeleteOpeningHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"msg": "hello"})

}

func UpdateOpeningHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"msg": "hello"})

}

func ListOpeningHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"msg": "hello"})

}
