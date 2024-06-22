package models

import "gorm.io/gorm"

type Investment struct {
	gorm.Model
	UserID       uint    `json:"user_id"`
	InstrumentID uint    `json:"instrument_id"`
	Amount       float64 `json:"amount"`

	Instrument Instrument `gorm:"foreignKey:InstrumentID" json:"-"`
	User       User       `gorm:"foreignKey:UserID" json:"-"`
}
