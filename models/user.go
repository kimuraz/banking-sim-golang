package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Token     string     `json:"token"`
}
