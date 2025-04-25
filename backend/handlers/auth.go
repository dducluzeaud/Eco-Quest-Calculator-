package handlers

import (
	"eco-quest-calculator/backend/models"
	"eco-quest-calculator/backend/utils"
	"encoding/json"
	"net/http"
)

const (
	headerContentType = "Content-Type"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.JSONError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		utils.JSONError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		utils.JSONError(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		utils.JSONError(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: struct {
			ID       uint   `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	}

	utils.JSONSuccess(w, "Login successful", response, http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		utils.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := input.HashPassword(); err != nil {
		utils.JSONError(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	if err := models.DB.Create(&input).Error; err != nil {
		utils.JSONError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.JSONSuccess(w, "User registered successfully", nil, http.StatusCreated)
}
