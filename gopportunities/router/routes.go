package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/valdenidelgado/go-projects/gopportunities/handler"
)

func initializeRoutes(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/opening", handler.ShowOpeningHandler)
		r.Post("/opening", handler.CreateOpeningHandler)
		r.Delete("/opening", handler.DeleteOpeningHandler)
		r.Put("/opening", handler.UpdateOpeningHandler)
		r.Get("/openings", handler.ListOpeningHandler)
	})
}
