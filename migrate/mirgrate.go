package main

import (
	"fmt"

	"github.com/MortezaHajilouei/golang-web-server/config"
	"github.com/MortezaHajilouei/golang-web-server/models"
)

func init() {
	config.InitConfig(".")
	config.ConnectDB(&config.Config_)
}

func main() {
	config.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
