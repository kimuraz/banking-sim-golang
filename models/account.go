package models

import "time"

type Account struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	UserID        uint       `json:"user_id"`
	Balance       float64    `json:"balance"`
	AccountNumber string     `json:"account_number"`
	User          User       `gorm:"foreignKey:UserID" json:"-"`
}
