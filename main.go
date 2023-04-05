package main

import (
	"log"

	"github.com/MortezaHajilouei/golang-web-server/config"
	"github.com/MortezaHajilouei/golang-web-server/middleware"
	"github.com/MortezaHajilouei/golang-web-server/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@securityDefinitions.apikey JWT
//@in header
//@name Authorization

var (
	server *gin.Engine
)

func init() {

	config.InitConfig(".")
	config.ConnectDB(&config.Config_)
	config.Auth(config.DB)
	config.Swagger()
	gin.SetMode(config.Config_.GIN_MODE)

	server = gin.Default()
}

func main() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.Config_.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	server.Use(middleware.DefaultStructuredLogger())
	server.Use(gin.Recovery())
	server.Use(middleware.Session())
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	route.SetupRoutes(config.DB, server)

	log.Fatal(server.Run(":" + config.Config_.ServerPort))
}
