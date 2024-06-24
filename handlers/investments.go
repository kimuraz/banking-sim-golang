package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
)

// GetInvestments godoc
// @Summary Get investments
// @Description Get all investments for the authenticated user
// @Tags investments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Investment
// @Router /investments [get]
func GetInvestments(c *gin.Context) {
	user, _ := c.Get("user")
	var investments []models.Investment
	models.DB.Where("user_id = ?", user.(*models.User).ID).Find(&investments)

	c.JSON(200, investments)
}
