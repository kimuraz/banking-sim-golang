package models

import "time"

type Investment struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	DeletedAt    *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	UserID       uint       `json:"user_id"`
	InstrumentID uint       `json:"instrument_id"`
	Amount       float64    `json:"amount"`

	Instrument Instrument `gorm:"foreignKey:InstrumentID" json:"-"`
	User       User       `gorm:"foreignKey:UserID" json:"-"`
}
