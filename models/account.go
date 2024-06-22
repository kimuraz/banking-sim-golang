package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserID        uint    `json:"user_id"`
	Balance       float64 `json:"balance"`
	AccountNumber string  `json:"account_number"`
	User          User    `gorm:"foreignKey:UserID" json:"-"`
}
