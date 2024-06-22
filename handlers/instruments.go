package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInstruments(c *gin.Context) {
	var instruments []models.Instrument
	models.DB.Find(&instruments)

	c.JSON(http.StatusOK, instruments)
}

func GetInstrumentsCategories(c *gin.Context) {
	var categories []models.InstrumentCategory
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)
}
