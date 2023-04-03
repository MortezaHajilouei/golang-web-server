package models

type Permission struct {
	Id     uint   `gorm:"autoIncrement;primary_key"`
	Title  string `gorm:"unique;not null;"`
	Method string
	Path   string
}

type Role struct {
	Id          uint         `gorm:"autoIncrement;primary_key"`
	Title       string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"foreignKey:Id;"`
}
