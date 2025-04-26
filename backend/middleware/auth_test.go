package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateRefreshToken(t *testing.T) {
	token, err := GenerateRefreshToken(1)
	if err != nil {
		t.Errorf("Failed to generate token: %v", err)
	}
	if token == "" {
		t.Error("Generated token is empty")
	}
}

func TestValidateToken(t *testing.T) {
	// Test valid token
	userID := uint(1)
	token, err := GenerateRefreshToken(userID)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate valid token: %v", err)
	}

	if claims == nil {
		t.Fatal("Claims should not be nil")
	}

	if claims.UserID != userID {
		t.Errorf("Expected user ID %d, got %d", userID, claims.UserID)
	}

	// Test invalid token
	_, err = ValidateToken("invalid.token.string")
	if err == nil {
		t.Error("Expected error for invalid token, got nil")
	}
}

func TestAuthMiddleware(t *testing.T) {
	// Create a mock handler
	nextHandler := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		// Check if user ID is in context
		userID := r.Context().Value(UserIDKey)
		if userID == nil {
			t.Error("User ID not found in context")
		}
	})

	middleware := AuthMiddleware(nextHandler)

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "No Authorization header",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid Authorization format",
			token:          "InvalidFormat",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid token",
			token:          "Bearer invalid.token.here",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Valid token",
			token:          "", // Will be set in test
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.name == "Valid token" {
				token, _ := GenerateRefreshToken(1)
				tt.token = "Bearer " + token
			}
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}

			rr := httptest.NewRecorder()
			middleware.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}
