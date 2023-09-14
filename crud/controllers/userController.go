package controllers

import (
	"github.com/0xAckerMan/HNGx/crud/initializers"
	"github.com/0xAckerMan/HNGx/crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
)

func UsersCreate (c *gin.Context){
	var body struct{
		Name string `json:"name"`
		Email string `json:"email"`
	}
	c.BindJSON(&body)
	//create a new user
	user := models.Users{Name: body.Name, Email: body.Email}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersIndex (c *gin.Context){
	var users []models.Users
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UsersShow (c *gin.Context){
	var user models.Users
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersUpdate(c *gin.Context) {
	var user models.Users
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	user.Name = body.Name
	user.Email = body.Email

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func UsersDelete(c *gin.Context) {
	var user models.Users
	result := initializers.DB.First(&user, c.Param("id"))

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(204, gin.H{"status": "success"})
}

