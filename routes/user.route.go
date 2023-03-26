package routes

import (
	"github.com/MortezaHajilouei/golang-web-server/controllers"
	"github.com/MortezaHajilouei/golang-web-server/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/all", uc.userController.GetAll)
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
