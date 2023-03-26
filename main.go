package main

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/controllers"
	"github.com/MortezaHajilouei/golang-web-server/docs"
	"github.com/MortezaHajilouei/golang-web-server/initializers"
	"github.com/MortezaHajilouei/golang-web-server/middleware"
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

	gin.SetMode(config.GIN_MODE)
	server = gin.Default()

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

}

func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	server.Use(middleware.DefaultStructuredLogger())
	server.Use(gin.Recovery())

	router := server.Group("/api/v1")
	{
		AuthRouteController.AuthRoute(router)
		UserRouteController.UserRoute(router)
		router.GET("/job/", controllers.Job1)
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url,
		ginSwagger.DefaultModelsExpandDepth(-1)))

	log.Fatal(server.Run(":" + config.ServerPort))
}
