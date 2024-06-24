package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAccounts godoc
// @Summary Get accounts
// @Description Get all accounts for the authenticated user
// @Tags accounts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Account
// @Router /accounts [get]
func GetAccounts(c *gin.Context) {
	user, _ := c.Get("user")
	var accounts []models.Account
	models.DB.Where("user_id = ?", user.(*models.User).ID).Find(&accounts)

	c.JSON(http.StatusOK, accounts)
}
