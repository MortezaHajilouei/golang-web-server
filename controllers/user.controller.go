package controllers

import (
	"net/http"
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
	SignUpUser(*gin.Context)
	SignUpUserLegal(*gin.Context)
	SignUpUserAdmin(*gin.Context)
	SignInUser(*gin.Context)
	GetMe(*gin.Context)
}

func NewUserController(repo repository.UserRepository) UserController {
	return userController{
		userRepo: repo,
	}
}

func (uc userController) AddUserAndResponse(ctx *gin.Context, newUser models.User) {
	result, err := uc.userRepo.AddUser(newUser)
	if err != nil && strings.Contains(err.Error(), "duplicated key not allowed") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "کاربری با این مشخصات وجود دارد"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	userResponse := &models.UserResponse{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// RegisterUser godoc
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
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        strings.ToLower(payload.Email),
		Password:     hashedPassword,
		Mobile:       payload.Mobile,
		NationalCode: payload.NationalCode,
		Telephone:    payload.Telephone,
		Address:      payload.Address,
		BirthDay:     payload.BirthDay,
		IBan:         payload.IBan,
		BankId:       payload.BankId,
	}
	uc.AddUserAndResponse(ctx, newUser)
}

// RegisterLegalUser godoc
// @Summary      Register new legal user
// @Description  signup legal user
// @Tags         users
// @Param		 info body models.UserSignUpInputLegal true "Signup"
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.UserResponse
// @Failure      409  {object}	utils.ErrorMessage
// @Failure      502  {object}	utils.ErrorMessage
// @Router       /user/register/legal [post]
func (uc userController) SignUpUserLegal(ctx *gin.Context) {
	var payload *models.UserSignUpInputLegal

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
		FirstName:       payload.FirstName,
		LastName:        payload.LastName,
		Email:           strings.ToLower(payload.Email),
		Password:        hashedPassword,
		Mobile:          payload.Mobile,
		NationalCode:    payload.NationalCode,
		Telephone:       payload.Telephone,
		Address:         payload.Address,
		IBan:            payload.IBan,
		BankId:          payload.BankId,
		BirthDay:        payload.BirthDay,
		IsLegal:         true,
		ZipCode:         payload.ZipCode,
		RegistrationNum: payload.RegistrationNum,
		NationalNum:     payload.NationalNum,
		EconomyNum:      payload.EconomyNum,
	}
	uc.AddUserAndResponse(ctx, newUser)
}

// RegisterAdminUser godoc
// @Summary      Register new admin user
// @Description  signup admin user
// @Tags         users
// @Param		 info body models.UserSignUpInputAdmin true "Signup"
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.UserResponse
// @Failure      409  {object}	utils.ErrorMessage
// @Failure      502  {object}	utils.ErrorMessage
// @Router       /user/register/admin [post]
func (uc userController) SignUpUserAdmin(ctx *gin.Context) {
	var payload *models.UserSignUpInputAdmin

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
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        strings.ToLower(payload.Email),
		Password:     hashedPassword,
		Mobile:       payload.Mobile,
		NationalCode: payload.NationalCode,
		Telephone:    payload.Telephone,
		BirthDay:     payload.BirthDay,
		Address:      payload.Address,
		IsSuperUser:  true,
		UserName:     payload.UserName,
	}
	uc.AddUserAndResponse(ctx, newUser)
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
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.Id, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	resp := gin.H{
		"access_token": access_token,
		"expires_in":   config.AccessTokenMaxAge,
		"isSuperUser":  user.IsSuperUser,
		"isLegal":      user.IsLegal,
	}

	ctx.JSON(http.StatusOK, resp)
}

// ProfileDetail godoc
// @Summary      Show My Profile Detail
// @Security 	 ApiKeyAuth
// @Description  get Profile detail
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.UserResponse
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /user/me/ [get]
func (uc userController) GetMe(ctx *gin.Context) {
	user := ctx.MustGet("currentUser").(models.User)
	userResponse := &models.UserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, gin.H{"user": userResponse})
}
