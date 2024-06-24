package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetInstruments godoc
// @Summary Get investments instruments
// @Description Get all instruments
// @Tags instruments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Instrument
// @Router /instruments [get]
func GetInstruments(c *gin.Context) {
	var instruments []models.Instrument
	models.DB.Find(&instruments)

	c.JSON(http.StatusOK, instruments)
}

// GetInstrumentsCategories godoc
// @Summary Get investments instruments categories
// @Description Get all instruments categories
// @Tags instruments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {array} models.InstrumentCategory
// @Router /instruments_categories [get]
func GetInstrumentsCategories(c *gin.Context) {
	var categories []models.InstrumentCategory
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)
}
