package main

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/controllers"
	docs "github.com/MortezaHajilouei/golang-web-server/docs"
	"github.com/MortezaHajilouei/golang-web-server/initializers"
	"github.com/MortezaHajilouei/golang-web-server/routes"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	docs.SwaggerInfo.Title = "GardeshPay"
	docs.SwaggerInfo.Description = "GardeshPay"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	//AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	//UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

// Default
func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api/v1")
	{
		router_user := router.Group("/user")
		{
			AuthRouteController.AuthRoute(router_user)
			UserRouteController.UserRoute(router_user)
		}
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url, ginSwagger.DefaultModelsExpandDepth(-1)))
	log.Fatal(server.Run(":" + config.ServerPort))

}
