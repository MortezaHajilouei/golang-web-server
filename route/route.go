package route

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/controllers"
	"github.com/MortezaHajilouei/golang-web-server/middleware"
	"github.com/MortezaHajilouei/golang-web-server/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, httpRouter *gin.Engine) {

	userRepository := repository.NewUserRepository(db)
	fileRepository := repository.NewFileRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	permissionRepository := repository.NewPermissionRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	if err := fileRepository.Migrate(); err != nil {
		log.Fatal("File migrate err", err)
	}
	if err := roleRepository.Migrate(); err != nil {
		log.Fatal("Role migrate err", err)
	}
	if err := permissionRepository.Migrate(); err != nil {
		log.Fatal("Permission migrate err", err)
	}

	userController := controllers.NewUserController(userRepository)
	apiRoutes := httpRouter.Group("/api/v1")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			//userRoutes.GET("/:id", userController.GetUser)
			userRoutes.POST("/register/", userController.SignUpUser)
			userRoutes.POST("/register/legal", userController.SignUpUserLegal)
			userRoutes.POST("/register/admin", userController.SignUpUserAdmin)
			userRoutes.POST("/login/", userController.SignInUser)
			userRoutes.GET("/me/", middleware.Authenticate(), userController.GetMe)
		}
	}
}
