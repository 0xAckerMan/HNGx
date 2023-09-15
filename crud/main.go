package main

import (
	"github.com/0xAckerMan/HNGx/crud/controllers"
	"github.com/0xAckerMan/HNGx/crud/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	// initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main(){
	r := gin.Default()
	r.POST("/api", controllers.UsersCreate)
	r.GET("/api", controllers.UsersIndex)
	r.GET("/api/:id", controllers.UsersShow)
	r.PUT("/api/:id", controllers.UsersUpdate)
	r.DELETE("/api/:id", controllers.UsersDelete)

	r.Run(":80") // listen and serve on
}