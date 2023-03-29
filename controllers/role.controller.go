package controllers

import (
	"github.com/MortezaHajilouei/golang-web-server/repository"
	"github.com/gin-gonic/gin"
)

type RoleController interface {
	AddRole(*gin.Context)
}

func NewRoleController(repo repository.UserRepository) UserController {
	return userController{
		userRepo: repo,
	}
}
