package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type TransactionResponse struct {
	ID          uint    `json:"id"`
	AccountID   uint    `json:"account_id"`
	Amount      float64 `json:"amount"`
	ToAccountID uint    `json:"to_account_id"`
	CategoryID  uint    `json:"category_id"`
	Datetime    string  `json:"datetime"`
	Account     string  `json:"account"`
	ToAccount   string  `json:"to_account"`
}

// GetTransactions godoc
// @Summary Get transactions
// @Description Get all transactions for the authenticated user
// @Tags accounts, transactions
// @Security BearerAuth
// @Param account_id path string true "Account ID"
// @Param date_from query string false "Date from"
// @Param date_to query string false "Date to"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order query string false "Order by date ASC or DESC"
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} TransactionResponse
// @Router /accounts/{account_id}/transactions [get]
func GetTransactions(c *gin.Context) {
	user, _ := c.Get("user")
	accountId := c.Param("account_id")

	account := models.Account{}
	models.DB.Where("id = ?", accountId).First(&account)

	if account.UserID != user.(*models.User).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to view this account"})
		return
	}

	var transactions []models.Transaction
	var total int64
	query := models.DB.
		Model(&models.Transaction{}).
		Where("account_id = ? or to_account_id = ?", accountId, accountId)

	dateFrom := c.Query("date_from")
	if dateFrom != "" {
		query = query.Where("datetime >= ?", dateFrom)
	}

	dateTo := c.Query("date_to")
	if dateTo != "" {
		query = query.Where("datetime <= ?", dateTo)
	}

	order := c.Query("order")
	if order == "ASC" || order == "DESC" {
		query = query.Order("datetime " + order)
	}

	query.Count(&total)
	c.Header("X-Total-Count", strconv.Itoa(int(total)))

	offset, limit := GetOffsetLimit(c)
	query.Offset(offset).Limit(limit)
	query.Preload("Account", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, account_number")
	})
	query.Preload("ToAccount", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, account_number")
	})
	query.Find(&transactions)

	var response []TransactionResponse
	for _, transaction := range transactions {
		response = append(response, TransactionResponse{
			ID:          transaction.ID,
			AccountID:   transaction.AccountID,
			Amount:      transaction.Amount,
			ToAccountID: transaction.ToAccountID,
			CategoryID:  transaction.CategoryID,
			Datetime:    transaction.Datetime.Format("2006-01-02 15:04:05"),
			Account:     transaction.Account.AccountNumber,
			ToAccount:   transaction.ToAccount.AccountNumber,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetTransactionCategories(c *gin.Context) {
	var categories []models.TransactionCategory
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)
}
