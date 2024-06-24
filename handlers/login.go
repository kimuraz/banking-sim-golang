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

// Login godoc
// @Summary Login
// @Description Login, note that the password is just `password`
// @Tags auth
// @Accept json
// @Produce json
// @Param input body LoginInput true "Login"
// @Success 200 {object} models.User
// @Router /auth [post]
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

// GetEmails godoc
// @Summary Get emails
// @Description Get all emails, use it for auth
// @Tags emails
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router /emails [get]
func GetEmails(c *gin.Context) {
	var emails []string
	models.DB.Table("users").Pluck("email", &emails)

	c.JSON(200, emails)
}
