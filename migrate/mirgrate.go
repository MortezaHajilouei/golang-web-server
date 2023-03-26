package main

import (
	"fmt"
	"log"

	"github.com/MortezaHajilouei/golang-web-server/initializers"
	"github.com/MortezaHajilouei/golang-web-server/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
