package models

import "gorm.io/gorm"

type TransactionCategory struct {
	gorm.Model
	Name string `json:"name"`
}

type Transaction struct {
	gorm.Model
	AccountID   uint    `json:"account_id"`
	Amount      float64 `json:"amount"`
	ToAccountID uint    `json:"to_account"`
	CategoryID  uint    `json:"category_id"`

	Account   Account             `gorm:"foreignKey:AccountID" json:"-"`
	ToAccount Account             `gorm:"foreignKey:ToAccountID" json:"-"`
	Category  TransactionCategory `gorm:"foreignKey:CategoryID" json:"-"`
}
