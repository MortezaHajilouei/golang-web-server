package controllers

import (
	"github.com/MortezaHajilouei/golang-web-server/repository"
	"github.com/gin-gonic/gin"
)

type roleController struct {
	roleRepo repository.RoleRepository
}

type RoleController interface {
	AddRole(*gin.Context)
	AddPermission(*gin.Context)
}

func NewRoleController(repo repository.RoleRepository) RoleController {
	return roleController{
		roleRepo: repo,
	}
}

func (uc roleController) AddRole(ctx *gin.Context) {
	ctx.JSON(501, "Not Implemented")
}

func (uc roleController) AddPermission(ctx *gin.Context) {
	ctx.JSON(501, "Not Implemented")
}
