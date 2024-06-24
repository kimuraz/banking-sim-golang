package models

import (
	"time"
)

type InstrumentCategory struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
}

type Instrument struct {
	ID                   uint               `gorm:"primaryKey" json:"id"`
	DeletedAt            *time.Time         `gorm:"index" json:"deleted_at,omitempty"`
	Name                 string             `json:"name"`
	Symbol               string             `json:"symbol"`
	Price                float64            `json:"price"`
	InstrumentCategoryID uint               `json:"instrument_category_id"`
	InstrumentCategory   InstrumentCategory `gorm:"foreignKey:InstrumentCategoryID" json:"-"`
}
