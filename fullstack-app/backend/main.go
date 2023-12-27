package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/valdenidelgado/go-projects/fullstack-app/backend/internal/controllers"
	"github.com/valdenidelgado/go-projects/fullstack-app/backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	router := chi.NewRouter()

	uc := controllers.NewUserController(db)

	router.Get("/", uc.Index)
	router.Post("/users", uc.Create)
	router.Patch("/users/{id}", uc.Update)
	// router.Patch("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := chi.URLParam(r, "id")
	// 	fmt.Println(id)
	// })

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
