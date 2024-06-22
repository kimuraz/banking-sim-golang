package handlers

import (
	"banking_sim/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	models.DB.Debug().Where("email = ?", input.Email).First(&user)
	result := models.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user)
	fmt.Println(result)
	fmt.Println(user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, user)
}
