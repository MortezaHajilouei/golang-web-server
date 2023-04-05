package models

import (
	"time"
)

type User struct {
	Id           uint      `gorm:"autoIncrement;primary_key"`
	FirstName    string    `gorm:"type:varchar(255);"` //filter:"param:login;searchable;filterable;"
	LastName     string    `gorm:"type:varchar(255);"`
	Email        string    `gorm:"type:varchar(255);unique;"`
	Mobile       string    `gorm:"type:varchar(255);unique;"`
	NationalCode string    `gorm:"type:varchar(50);unique;"`
	Telephone    string    `gorm:"type:varchar(255);"`
	Address      string    `gorm:"type:string;"`
	Password     string    `gorm:"type:varchar(255);"`
	SecretCode   string    `gorm:"type:varchar(255);"`
	BirthDay     time.Time `gorm:"type:Date;"`

	ProfilePic  string
	IBan        string
	BankId      string
	WalletAlias string

	IsSuperUser bool `gorm:"default:false"`
	UserName    string

	IsLegal         bool `gorm:"default:false"`
	ZipCode         string
	RegistrationNum string
	NationalNum     string
	EconomyNum      string
	GuildTitle      string

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
	FirstName       string    `json:"firstname" binding:"required" default:"John"`
	LastName        string    `json:"lastname" binding:"required" default:"Lenon"`
	Email           string    `json:"email" binding:"required,email" default:"Email@example.com"`
	Mobile          string    `json:"phone" binding:"required,min=11,max=11" default:"09129999999"`
	NationalCode    string    `json:"national_code" binding:"required,min=10,max=10" default:"0000000000"`
	Password        string    `json:"password" binding:"required,min=5" default:"password"`
	PasswordConfirm string    `json:"passwordConfirm" binding:"required,min=5" default:"password"`
	Telephone       string    `json:"telephone" binding:"required" default:"02122222222"`
	Address         string    `json:"address" binding:"required" default:"تهران"`
	BirthDay        time.Time `json:"date" binding:"required" default:"1990-01-01T00:00:00.0Z"`

	IBan   string `json:"iBan" binding:"required" default:"1"`
	BankId string `json:"bankId" binding:"required" default:"2"`
}

type UserSignUpInputLegal struct {
	FirstName       string    `json:"firstname" binding:"required" default:"John"`
	LastName        string    `json:"lastname" binding:"required" default:"Lenon"`
	Email           string    `json:"email" binding:"required,email" default:"Email@example.com"`
	Mobile          string    `json:"phone" binding:"required,min=11,max=11" default:"09129999999"`
	NationalCode    string    `json:"national_code" binding:"required,min=10,max=10" default:"0000000000"`
	Password        string    `json:"password" binding:"required,min=5" default:"password"`
	PasswordConfirm string    `json:"passwordConfirm" binding:"required,min=5" default:"password"`
	Telephone       string    `json:"telephone" binding:"required" default:"02122222222"`
	Address         string    `json:"address" binding:"required" default:"تهران"`
	BirthDay        time.Time `json:"date" binding:"required" default:"1990-01-01T00:00:00.0Z"`

	IBan   string `json:"iBan" binding:"required" default:"1"`
	BankId string `json:"bankId" binding:"required" default:"1"`

	ZipCode         string `json:"zipCode" binding:"required" default:"1111111111"`
	RegistrationNum string `json:"registrationNum" binding:"required" default:"1111111111"`
	NationalNum     string `json:"nationalNum" binding:"required" default:"1111111111" `
	EconomyNum      string `json:"economyNum" binding:"required" default:"1111111111" `
	GuildTitle      string `json:"guildTitle" binding:"required" default:"GuildTitle" `
}

type UserSignUpInputAdmin struct {
	UserName string `json:"username" binding:"required" default:"john"`

	FirstName       string    `json:"firstname" binding:"required" default:"John"`
	LastName        string    `json:"lastname" binding:"required" default:"Lenon"`
	Email           string    `json:"email" binding:"required,email" default:"Email@example.com"`
	Mobile          string    `json:"phone" binding:"required,min=11,max=11" default:"09129999999"`
	NationalCode    string    `json:"national_code" binding:"required,min=10,max=10" default:"0000000000"`
	Password        string    `json:"password" binding:"required,min=5" default:"password"`
	PasswordConfirm string    `json:"passwordConfirm" binding:"required,min=5" default:"password"`
	Telephone       string    `json:"telephone" binding:"required" default:"02122222222"`
	Address         string    `json:"address" binding:"required" default:"تهران"`
	BirthDay        time.Time `json:"date" binding:"required" default:"1990-01-01T00:00:00.0Z"`
}

type UserSignInInput struct {
	Email    string `json:"email"  binding:"required" default:"email@example.com"`
	Password string `json:"password"  binding:"required" default:"password"`
}

type UserResponse struct {
	Id        uint      `json:"id,omitempty" `
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
