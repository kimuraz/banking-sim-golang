package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTransactions(c *gin.Context) {
	user, _ := c.Get("user")
	var accounts []models.Account
	models.DB.Where("user_id = ?", user.(*models.User).ID).Find(&accounts)

	var accountIDs []uint
	for _, account := range accounts {
		accountIDs = append(accountIDs, account.ID)
	}

	var transactions []models.Transaction
	models.DB.Preload("Category").Where("from_account_id IN ? OR to_account_id IN ?", accountIDs, accountIDs).Find(&transactions)

	c.JSON(http.StatusOK, transactions)
}

func GetTransactionCategories(c *gin.Context) {
	var categories []models.TransactionCategory
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)
}
