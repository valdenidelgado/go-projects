package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Initialize() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json"))
	initializeRoutes(router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
