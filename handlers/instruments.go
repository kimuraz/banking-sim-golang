package handlers

import (
	"banking_sim/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
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

type InstrumentUpdateMsg struct {
	ID           uint      `json:"id"`
	Datetime     time.Time `json:"datetime"`
	CurrentPrice float64   `json:"price"`
}

func GenerateRandomInstrumentUpdateMsg(ids []uint) *InstrumentUpdateMsg {
	var instrument models.Instrument
	randomId := ids[rand.Intn(len(ids))]
	models.DB.Where("id = ?", randomId).First(&instrument)
	if instrument.ID == 0 {
		return nil
	}
	increaseDecreaseRatio := rand.Float64() + float64(rand.Intn(1))

	simPrice := instrument.Price * increaseDecreaseRatio
	simPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", simPrice), 64)

	return &InstrumentUpdateMsg{
		ID:           instrument.ID,
		Datetime:     time.Now().UTC(),
		CurrentPrice: simPrice,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InstrumentsWebsocket(c *gin.Context) {
	idsString := c.Query("ids")
	ids := []uint{}
	if idsString != "" {
		idsStringArr := strings.Split(idsString, ",")
		for _, id := range idsStringArr {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ids"})
				return
			}
			ids = append(ids, uint(idInt))
		}
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	done := make(chan struct{})

	// Start a goroutine to close the connection if any message is received
	go func() {
		defer close(done)
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				return
			}
			conn.Close()
			return
		}
	}()

	for {
		select {
		case <-done:
			return
		default:
			msg := GenerateRandomInstrumentUpdateMsg(ids)
			if msg == nil {
				continue
			}
			jsonMsg, _ := json.Marshal(msg)
			err := conn.WriteMessage(websocket.TextMessage, jsonMsg)
			if err != nil {
				log.Println(err)
				return
			}
			delay := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(delay)
		}
	}
}
