package main

import (
	docs "github.com/MortezaHajilouei/golang-web-server/docs"
	p1 "github.com/MortezaHajilouei/golang-web-server/pkg1"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "GardeshPay"
	docs.SwaggerInfo.Description = "GardeshPay"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		pk1 := v1.Group("/p1")
		{
			pk1.GET("/hello", p1.Helloworld)
			pk1.GET("/ping", p1.PingExample)
		}
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url, ginSwagger.DefaultModelsExpandDepth(-1)))
	r.Run(":8080")

}
