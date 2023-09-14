package main

import (
	"github.com/0xAckerMan/HNGx/crud/initializers"
	"github.com/0xAckerMan/HNGx/crud/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main(){
	initializers.DB.AutoMigrate(&models.Users{})
}