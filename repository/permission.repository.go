package repository

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/models"
	"gorm.io/gorm"
)

type permissionRepository struct {
	DB *gorm.DB
}

type PermissionRepository interface {
	Migrate() error
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return permissionRepository{
		DB: db,
	}
}

func (u permissionRepository) Migrate() error {
	log.Print("[PermissionRepository]...Migrate")
	return u.DB.AutoMigrate(&models.Permission{})
}
