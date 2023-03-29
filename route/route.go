package route

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/controllers"
	"github.com/MortezaHajilouei/golang-web-server/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, httpRouter *gin.Engine) {

	// Initialize  casbin adapter
	//adapter := gormadapter.NewAdapterByDB(db)

	// if err != nil {
	// 	panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	// }

	// // Load model configuration file and policy store adapter
	//enforcer := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	// }

	// //add policy
	// if hasPolicy := enforcer.HasPolicy("doctor", "report", "read"); !hasPolicy {
	// 	enforcer.AddPolicy("doctor", "report", "read")
	// }
	// if hasPolicy := enforcer.HasPolicy("doctor", "report", "write"); !hasPolicy {
	// 	enforcer.AddPolicy("doctor", "report", "write")
	// }
	// if hasPolicy := enforcer.HasPolicy("patient", "report", "read"); !hasPolicy {
	// 	enforcer.AddPolicy("patient", "report", "read")
	// }

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}

	userController := controllers.NewUserController(userRepository)
	apiRoutes := httpRouter.Group("/api/v1")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			userRoutes.GET("/:id", userController.GetUser)
			userRoutes.GET("/", userController.GetMe)
			userRoutes.POST("/register/", userController.SignUpUser)
			userRoutes.POST("/login/", userController.SignInUser)
		}
		usersRoutes := apiRoutes.Group("/users")
		{
			usersRoutes.GET("/", userController.GetAllUser)
		}
	}

	// userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	// {
	// 	userProtectedRoutes.GET("/", middleware.Authorize("report", "read", enforcer), userController.GetAllUser)
	// 	userProtectedRoutes.GET("/:user", middleware.Authorize("report", "read", enforcer), userController.GetUser)
	// 	userProtectedRoutes.PUT("/:user", middleware.Authorize("report", "write", enforcer), userController.UpdateUser)
	// 	userProtectedRoutes.DELETE("/:user", middleware.Authorize("report", "write", enforcer), userController.DeleteUser)
	// }

}
