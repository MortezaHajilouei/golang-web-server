package models

import (
	"time"
)

type User struct {
	Id           uint      `gorm:"type:uint;autoIncrement;primary_key"`
	FirstName    string    `gorm:"type:varchar(255);not null" ` //filter:"param:login;searchable;filterable;"
	LastName     string    `gorm:"type:varchar(255);not null" `
	Email        string    `gorm:"type:varchar(255);unique;not null" `
	Mobile       string    `gorm:"type:varchar(255);unique;not null"`
	NationalCode string    `gorm:"type:varchar(50);unique;default:null"`
	Telephone    string    `gorm:"type:varchar(255);not null"`
	Address      string    `gorm:"type:string;not null"`
	Password     string    `gorm:"not null"`
	SecretCode   string    `gorm:"type:varchar(255);not null"`
	BirthDay     time.Time `gorm:"type:Date;not null"`

	ProfilePic  string
	IBan        string
	BankId      string
	WalletAlias string

	IsLegal         bool   `gorm:"type:boolean;default:false"`
	ZipCode         string `gorm:"type:string;not null"`
	RegistrationNum string `gorm:"type:string;not null"`
	NationalNum     string `gorm:"type:string;not null"`
	EconomyNum      string `gorm:"type:string;not null"`
	GuildTitle      string `gorm:"type:string;not null"`

	GetNewsLetter  bool `gorm:"default:false;"`
	GetSms         bool `gorm:"default:false;"`
	GetEmail       bool `gorm:"default:false;"`
	MobileVerified bool `gorm:"default:false;"`
	EmailVerified  bool `gorm:"default:false;"`
	HasWallet      bool `gorm:"default:false;"`
	IsActive       bool `gorm:"default:false;"`
	Deposit        bool `gorm:"default:false;"`

	CreatedAt          time.Time `gorm:"autoCreateTime;<-:create;"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime;"`
	PasswordModifiedAt time.Time `gorm:"autoCreateTime;"`

	Files []File `gorm:"foreignKey:Id;"`
	Role  Role   `gorm:"foreignKey:Id"`
}

type UserSignUpInput struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Mobile          string `json:"phone" binding:"required"`
	NationalCode    string `json:"national_code" binding:"required"`
	Password        string `json:"password" binding:"required,min=5"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type UserSignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	Id        uint      `json:"id,omitempty" `
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
