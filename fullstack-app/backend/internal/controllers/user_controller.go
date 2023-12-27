package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/valdenidelgado/go-projects/fullstack-app/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(database *gorm.DB) *UserController {
	return &UserController{
		db: database,
	}
}

func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) {
	// welcome endpoint
	w.Write([]byte("Welcome to the User Controller"))
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(passwordHash)

	// check if user already exists
	var existingUser models.User
	if err := uc.db.Where("email = ?", user.Email).First(&existingUser); err.Error != gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already exists"))
		return
	}

	err = uc.db.Create(&user).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	// show user
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	//get id from url
	id := chi.URLParam(r, "id")

	fmt.Println(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//find user
	var userModel models.User

	if err := uc.db.First(&userModel, id).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(userModel)

	// update user
	var user models.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Name != "" {
		userModel.Name = user.Name
	}

	if user.Email != "" {
		userModel.Email = user.Email
	}
	if user.Avatar != "" {
		userModel.Avatar = user.Avatar
	}
	if user.Password != "" && user.OldPassword != "" {
		compare := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.OldPassword))
		if compare != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Old password does not match"))
			return
		}
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		userModel.Password = string(passwordHash)
	}

	err = uc.db.Save(&userModel).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
