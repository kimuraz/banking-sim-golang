package handlers

import (
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetInstruments godoc
// @Summary Get investments instruments
// @Description Get all instruments
// @Param category_id query string false "Category ID"
// @Param q query string false "Search by name"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param order query string false "Order by price ASC or DESC"
// @Tags instruments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Instrument
// @Router /instruments [get]
func GetInstruments(c *gin.Context) {
	var instruments []models.Instrument
	category := c.Query("category_id")
	search := c.Query("q")
	order := c.Query("order")

	query := models.DB.Model(&models.Instrument{})

	if category != "" {
		query = query.Where("instrument_category_id = ?", category)
	}

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if order == "ASC" || order == "DESC" {
		query = query.Order("price " + order)
	}

	var total int64
	query.Count(&total)
	c.Header("X-Total-Count", strconv.Itoa(int(total)))

	offset, limit := GetOffsetLimit(c)
	query.Offset(offset).Limit(limit).Find(&instruments)

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
