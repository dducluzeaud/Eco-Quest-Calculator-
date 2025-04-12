package models

import (
	"testing"
	"time"
)

func TestUserValidate(t *testing.T) {
	baseUser := User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123456",
	}

	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name:    "valid user",
			user:    baseUser,
			wantErr: false,
		},
		{
			name: "invalid email",
			user: User{
				Username: baseUser.Username,
				Email:    "invalid-email",
				Password: baseUser.Password,
			},
			wantErr: true,
		},
		{
			name: "password too short",
			user: User{
				Username: baseUser.Username,
				Email:    baseUser.Email,
				Password: "short",
			},
			wantErr: true,
		},
		{
			name: "username too short",
			user: User{
				Username: "ab",
				Email:    baseUser.Email,
				Password: baseUser.Password,
			},
			wantErr: true,
		},
		{
			name: "email missing @",
			user: User{
				Username: baseUser.Username,
				Email:    "testexample.com",
				Password: baseUser.Password,
			},
			wantErr: true,
		},
		{
			name: "password too long",
			user: User{
				Username: baseUser.Username,
				Email:    baseUser.Email,
				Password: "thispasswordiswaytoolongandshouldfailthevalidationcheck123456",
			},
			wantErr: true,
		},
		{
			name: "username too long",
			user: User{
				Username: "thisusernameiswaytoolongandshouldfailvalidation",
				Email:    baseUser.Email,
				Password: baseUser.Password,
			},
			wantErr: true,
		},
		{
			name:    "empty fields",
			user:    User{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	// Setup test database


	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "validpassword123",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: tt.password,
			}
			err := u.HashPassword()
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && u.Password == tt.password {
				t.Error("HashPassword() did not change the password")
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	password := "testpassword123"
	u := &User{Password: password}

	// First hash the password
	if err := u.HashPassword(); err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	tests := []struct {
		name          string
		inputPassword string
		wantErr       bool
	}{
		{
			name:          "correct password",
			inputPassword: password,
			wantErr:       false,
		},
		{
			name:          "incorrect password",
			inputPassword: "wrongpassword123",
			wantErr:       true,
		},
		{
			name:          "empty password",
			inputPassword: "",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := u.CheckPassword(tt.inputPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserFields(t *testing.T) {
	now := time.Now()
	u := &User{
		ID:        1,
		CreatedAt: now,
		UpdatedAt: now,
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "testpassword123",
	}

	if u.ID != 1 {
		t.Errorf("Expected ID = 1, got %v", u.ID)
	}

	if !u.CreatedAt.Equal(now) {
		t.Errorf("Expected CreatedAt = %v, got %v", now, u.CreatedAt)
	}

	if !u.UpdatedAt.Equal(now) {
		t.Errorf("Expected UpdatedAt = %v, got %v", now, u.UpdatedAt)
	}

	if u.DeletedAt.Valid {
		t.Error("Expected DeletedAt to be invalid for new user")
	}
}

func TestPasswordOperations(t *testing.T) {
	user := User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "testpassword123",
	}

	// Test password hashing
	originalPassword := user.Password
	if err := user.HashPassword(); err != nil {
		t.Errorf("Failed to hash password: %v", err)
	}
	if user.Password == originalPassword {
		t.Error("Password was not hashed")
	}

	// Test password verification
	if err := user.CheckPassword(originalPassword); err != nil {
		t.Errorf("Failed to verify correct password: %v", err)
	}
	if err := user.CheckPassword("wrongpassword"); err == nil {
		t.Error("Should fail with wrong password")
	}
}


