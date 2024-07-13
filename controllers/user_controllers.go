package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sajin-shrestha/gormTemplate/models"
	"github.com/sajin-shrestha/gormTemplate/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(database *gorm.DB) {
	db = database
}

// function to create-user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	db.Create(&user)

	utils.RespondJSON(w, http.StatusCreated, user)
}

// function to get-all-users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.Find(&users)

	utils.RespondJSON(w, http.StatusOK, users)
}

// function to get-user-by-id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var user models.User
	db.First(&user, id)
	if user.ID == 0 {
		utils.RespondJSON(w, http.StatusNotFound, "User nor found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, user)
}

// function to delete-user-by-id
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var user models.User
	db.First(&user, id)
	if user.ID == 0 {
		utils.RespondJSON(w, http.StatusNotFound, "User not found")
		return
	}

	db.Delete(&user)

	utils.RespondJSON(w, http.StatusOK, "User deleted")
}

// login-function
func Login(w http.ResponseWriter, r *http.Request) {
	var reqBody models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Mock user authentication (replace with actual authentication logic)
	// For demonstration, let's assume a hardcoded user/password check
	if reqBody.Username != "exampleuser" || reqBody.Password != "examplepassword" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(reqBody.Username)
	if err != nil {
		http.Error(w, "Failed to generate JWT token", http.StatusInternalServerError)
		return
	}

	// Return the JWT token in the response
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
