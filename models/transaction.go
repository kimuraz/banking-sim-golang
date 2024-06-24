package models

import "time"

type TransactionCategory struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
}

type Transaction struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	AccountID   uint       `json:"account_id"`
	Amount      float64    `json:"amount"`
	ToAccountID uint       `json:"to_account"`
	CategoryID  uint       `json:"category_id"`

	Account   Account             `gorm:"foreignKey:AccountID" json:"-"`
	ToAccount Account             `gorm:"foreignKey:ToAccountID" json:"-"`
	Category  TransactionCategory `gorm:"foreignKey:CategoryID" json:"-"`
}
