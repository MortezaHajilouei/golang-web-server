package config

import "github.com/MortezaHajilouei/golang-web-server/docs"

func Swagger() {
	docs.SwaggerInfo.Title = "GardeshPay"
	docs.SwaggerInfo.Description = "GardeshPay"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

}
