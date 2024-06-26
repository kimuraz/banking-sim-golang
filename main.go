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
