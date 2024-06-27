package main

import (
	_ "banking_sim/docs"
	"banking_sim/handlers"
	"banking_sim/middleware"
	"banking_sim/models"
	"github.com/cnjack/throttle"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// @title Banking Simulation API
// @version 1.0

// @description This is a read-only API for simulating banking operations.
// @description It allows you to get accounts, transactions, instruments, and investments.
// @description The endpoints that support pagination have a default page size of 15 and a maximum of 100. To check the total number of items, you can look at the `X-Total-Count` header.
// @description It requires authentication for most of the endpoints. You can use any email from the emails endpoint with password `password`.
// @description The API is rate-limited to 100 requests per second.
// @description There's a special websocket endpoint that allows you to get real-time updates on the instruments prices. You can connect to it using the `ws` endpoint.
// @description The websocket endpoint requires authentication, you can use the token from auth endpoint with `?token=[token]`
// @description The websocket will send you a message in random intervals with the updated prices of the instruments with id's that you pass as query param `?id=1,2,3`.
// @description The message will be in the following format: `{"id":1,"price":100.0, "datetime": "2024-01-01T00:00:00Z"}`

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

//@securityDefinitions.apikey BearerAuth
//@in header
//@name Authorization

// @BasePath /api/v1
func main() {
	models.InitDB()

	router := gin.Default()
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(configCors))

	auth := router.Group("/api/v1")
	auth.Use(throttle.Policy(&throttle.Quota{
		Limit:  100,
		Within: time.Second,
	}))
	auth.POST("/auth", handlers.Login)
	auth.GET("/emails", handlers.GetEmails)
	auth.GET("/ws", middleware.WsAuth(), handlers.InstrumentsWebsocket)
	auth.Use(middleware.Auth())
	{
		auth.GET("/accounts", handlers.GetAccounts)
		auth.GET("/accounts/:account_id/transactions", handlers.GetTransactions)
		auth.GET("/instruments", handlers.GetInstruments)
		auth.GET("/investments", handlers.GetInvestments)

		auth.GET("/transactions_categories", handlers.GetTransactionCategories)
		auth.GET("/instruments_categories", handlers.GetInstrumentsCategories)
	}

	router.Run(":8080")
}
