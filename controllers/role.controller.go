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
func (uc roleController) AddRole(ctx *gin.Context) {
	ctx.JSON(501, "Not Implemented")
}

func (uc roleController) AddPermission(ctx *gin.Context) {
	ctx.JSON(501, "Not Implemented")
}
