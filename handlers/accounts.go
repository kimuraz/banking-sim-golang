package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAccounts(c *gin.Context) {
	user, _ := c.Get("user")
	var accounts []models.Account
	models.DB.Where("user_id = ?", user.(*models.User).ID).Find(&accounts)

	c.JSON(http.StatusOK, accounts)
}
