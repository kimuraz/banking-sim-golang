package models

import "gorm.io/gorm"

type InstrumentCategory struct {
	gorm.Model
	Name string `json:"name"`
}

type Instrument struct {
	gorm.Model
	Name                 string             `json:"name"`
	Symbol               string             `json:"symbol"`
	Price                float64            `json:"price"`
	InstrumentCategoryID uint               `json:"instrument_category_id"`
	InstrumentCategory   InstrumentCategory `gorm:"foreignKey:InstrumentCategoryID" json:"-"`
}
