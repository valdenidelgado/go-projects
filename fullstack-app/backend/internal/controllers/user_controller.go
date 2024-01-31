package controllers

import (
	"encoding/json"
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

func (uc *UserController) ShowOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	if err := uc.db.First(&user, id).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
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

	if err := uc.db.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	var user []models.User

	if err := uc.db.Find(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	if err := uc.db.First(&user, id).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userDTO models.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userWithUpdatedEmail models.User
	if err := uc.db.Where("email = ?", userDTO.Email).First(&userWithUpdatedEmail).Error; err != gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email already exists"))
		return
	}

	if userWithUpdatedEmail.Id != user.Id {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email already exists"))
		return
	}

	if user.Name != "" {
		user.Name = userDTO.Name
	}
	if user.Email != "" {
		user.Email = userDTO.Email
	}
	if user.Avatar != "" {
		user.Avatar = userDTO.Avatar
	}

	if userDTO.Password != "" && userDTO.OldPassword != "" {
		compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.OldPassword))
		if compare != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Old password does not match"))
			return
		}
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(passwordHash)
	}

	if err := uc.db.Save(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated"))
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	if err := uc.db.First(&user, id).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uc.db.Delete(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
