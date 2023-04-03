package repository

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/models"
	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

type RoleRepository interface {
	Migrate() error
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{
		DB: db,
	}
}

func (u roleRepository) Migrate() error {
	log.Print("[RoleRepository]...Migrate")
	return u.DB.AutoMigrate(&models.Role{})
}
