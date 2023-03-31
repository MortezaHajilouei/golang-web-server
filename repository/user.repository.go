package repository

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/config"
	"github.com/MortezaHajilouei/golang-web-server/models"
	"github.com/MortezaHajilouei/golang-web-server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	AddUser(models.User) (models.User, error)
	GetUser(string) (models.User, error)
	GetByEmail(string) (models.User, error)
	GetAllUser(ctx *gin.Context) ([]models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser1(models.User) (models.User, error)
	Migrate() error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&models.User{})
}

func (u userRepository) GetUser(id string) (user models.User, err error) {
	return user, u.DB.First(&user, "id=?", id).Error
}

func (u userRepository) GetByEmail(email string) (user models.User, err error) {
	return user, u.DB.First(&user, "email=?", email).Error
}

func (u userRepository) GetAllUser(ctx *gin.Context) (users []models.User, err error) {
	var usersCount int64
	err = config.DB.Model(&models.User{}).Scopes(
		utils.FilterByQuery(ctx, utils.SEARCH)).Count(&usersCount).Find(&users).Error

	return users, err
}

func (u userRepository) AddUser(user models.User) (models.User, error) {
	return user, u.DB.Create(&user).Error
}

func (u userRepository) UpdateUser(user models.User) (models.User, error) {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, u.DB.Model(&user).Updates(&user).Error
}

func (u userRepository) DeleteUser1(user models.User) (models.User, error) {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, u.DB.Delete(&user).Error
}
