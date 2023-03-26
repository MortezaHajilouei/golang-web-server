package controllers

import (
	"net/http"

	"github.com/MortezaHajilouei/golang-web-server/initializers"
	"github.com/MortezaHajilouei/golang-web-server/models"
	"github.com/MortezaHajilouei/golang-web-server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// UserDetail godoc
// @Summary      Show User Detail
// @Security 	 ApiKeyAuth
// @Description  get User detail
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      404  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /users/me [get]
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// UsersList godoc
// @Summary      Show Users List
// @Security 	 ApiKeyAuth
// @Description  get users list
// @Param		 info query utils.SearchUser true "Signup"
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /users/all [get]
func (uc *UserController) GetAll(ctx *gin.Context) {
	var users []models.User
	var usersCount int64

	err := initializers.DB.Model(&models.User{}).Scopes(
		utils.FilterByQuery(ctx, utils.SEARCH),
	).Count(&usersCount).Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var users_response []models.UserResponse

	for i := 0; i < len(users); i++ {
		users_response = append(users_response, models.UserResponse{
			ID:        users[i].ID,
			Name:      users[i].Name,
			Email:     users[i].Email,
			CreatedAt: users[i].CreatedAt,
			UpdatedAt: users[i].UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}
