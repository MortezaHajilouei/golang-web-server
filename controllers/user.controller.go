package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/MortezaHajilouei/golang-web-server/config"
	"github.com/MortezaHajilouei/golang-web-server/models"
	"github.com/MortezaHajilouei/golang-web-server/repository"
	"github.com/MortezaHajilouei/golang-web-server/utils"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepo repository.UserRepository
}

type UserController interface {
	GetUser(*gin.Context)
	GetMe(*gin.Context)
	GetAllUser(*gin.Context)
	SignInUser(*gin.Context)
	SignUpUser(*gin.Context)
	//AddUser(enforcer *casbin.Enforcer) gin.HandlerFunc
	//UpdateUser(*gin.Context)
	//DeleteUser(*gin.Context)
}

func NewUserController(repo repository.UserRepository) UserController {
	return userController{
		userRepo: repo,
	}
}

// UserDetail godoc
// @Summary      Show User Detail
// @Security 	 ApiKeyAuth
// @Description  get User detail
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id  path string  true  "user id"
// @Success      200  {object}  models.User
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      404  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /user/{id} [get]
func (uc userController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.userRepo.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "user not found"})
		return
	}
	userResponse := &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// ProfileDetail godoc
// @Summary      Show My Profile Detail
// @Security 	 ApiKeyAuth
// @Description  get Profile detail
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /user/ [get]
func (uc userController) GetMe(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(models.User)
	userResponse := &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
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
// @Router       /users/ [get]
func (uc userController) GetAllUser(ctx *gin.Context) {
	var users []models.User

	users, err := uc.userRepo.GetAllUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
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

// Register godoc
// @Summary      Register new user
// @Description  signup user
// @Tags         users
// @Param		 info body models.UserSignUpInput true "Signup"
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.UserResponse
// @Failure      409  {object}	utils.ErrorMessage
// @Failure      502  {object}	utils.ErrorMessage
// @Router       /user/register/ [post]
func (uc userController) SignUpUser(ctx *gin.Context) {
	var payload *models.UserSignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: hashedPassword,
	}
	result, err := uc.userRepo.AddUser(newUser)
	if err != nil && strings.Contains(err.Error(), "duplicated key not allowed") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userResponse := &models.UserResponse{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// Login godoc
// @Summary      Login user
// @Description  Login user
// @Tags         users
// @Param		 info body models.UserSignInInput true "SignIn"
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /user/login/ [post]
func (ac userController) SignInUser(ctx *gin.Context) {
	var payload *models.UserSignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := ac.userRepo.GetByEmail(strings.ToLower(payload.Email))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config := config.Config_

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (uc *userController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config := config.Config_

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	//result := uc.userRepo.DB.First(&user, "id = ?", fmt.Sprint(sub))
	fmt.Printf("ss: ?", sub)
	str := strconv.Itoa(1)
	user, err := uc.userRepo.GetUser(str)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (uc *userController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
