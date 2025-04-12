package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Username     string         `gorm:"uniqueIndex" json:"username" validate:"required,min=3,max=32"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Password     string         `json:"password" validate:"required,min=12,max=50"`
	RefreshToken string         `json:"-" gorm:"type:text"`
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// HashPassword hashes the user's password.
func (u *User) HashPassword() error {
	if u.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword compares the provided password with the hashed password.
func (u *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}
