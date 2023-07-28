package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func initializeRoutes(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/opening", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]string{"msg": "hello"})
		})
		r.Post("/opening", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]string{"msg": "hello"})
		})
		r.Delete("/opening", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]string{"msg": "hello"})
		})
		r.Put("/opening", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]string{"msg": "hello"})
		})
		r.Get("/openings", func(w http.ResponseWriter, r *http.Request) {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, map[string]string{"msg": "hello"})
		})
	})
}
