package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(255);not null" filter:"param:login;searchable;filterable;"`
	Email     string    `gorm:"uniqueIndex;not null" `
	Phone     string    `gorm:"unique;null;"`
	Password  string    `gorm:"not null"`
	Verified  bool      `gorm:"not null;default:false;"`
	Enabled   bool      `gorm:"not null;default:false;"`
	CreatedAt time.Time `gorm:"autoCreateTime;<-:create;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"`
}

type UserSignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type UserSignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
