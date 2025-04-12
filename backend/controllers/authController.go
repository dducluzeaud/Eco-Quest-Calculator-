package controllers

import (
	"eco-quest-calculator/backend/models"
	"eco-quest-calculator/backend/utils"
	"fmt"
	"net/http"
	"strings"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Ensure the models.User struct has the necessary validation and hashing methods

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user input
	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not initialized"})
		return
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		user.Email = input.Email
		user.Password = input.Password

		if err := user.HashPassword(); err != nil {
			return err
		}

		if err := tx.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "idx_users_email") {
				return fmt.Errorf("account with this email already exists")
			}
			return err
		}

		return nil
	})

	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "already exists") {
			statusCode = http.StatusConflict
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// First, create a login request struct
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password format"})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate both tokens
	accessToken, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		fmt.Printf("Failed to generate access token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		fmt.Printf("Failed to generate refresh token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Store refresh token hash using SHA-256 instead of bcrypt
	h := sha256.New()
	h.Write([]byte(refreshToken))
	hashedRefreshToken := hex.EncodeToString(h.Sum(nil))

	// Update user with hashed refresh token using transaction
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		updates := map[string]interface{}{
			"refresh_token": hashedRefreshToken,
		}

		result := tx.Model(&user).Updates(updates)
		if result.Error != nil {
			fmt.Printf("Database error while updating refresh token: %v\n", result.Error)
			return result.Error
		}

		if result.RowsAffected == 0 {
			fmt.Printf("No rows affected when updating user ID: %d\n", user.ID)
			return fmt.Errorf("failed to update user record")
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Transaction failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save authentication data"})
		return
	}

	// Return successful response
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
		"expires_in":    900,
	})
}
