package repository

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/models"
	"gorm.io/gorm"
)

type fileRepository struct {
	DB *gorm.DB
}

type FileRepository interface {
	Migrate() error
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return fileRepository{
		DB: db,
	}
}

func (u fileRepository) Migrate() error {
	log.Print("[FileRepository]...Migrate")
	return u.DB.AutoMigrate(&models.File{})
}
