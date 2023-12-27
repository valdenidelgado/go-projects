package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/valdenidelgado/poc-api-go/configs"
	"github.com/valdenidelgado/poc-api-go/internal/entity"
	"github.com/valdenidelgado/poc-api-go/internal/infra/database"
	"github.com/valdenidelgado/poc-api-go/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, conf.TokenAuth, conf.JwtExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(conf.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetAllProducts)
	})
	r.Post("/user", userHandler.CreateUser)
	r.Post("/user/auth", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)

	// http.HandleFunc("/products", productHandler.CreateProduct)

	// http.ListenAndServe(":8000", nil)
}
