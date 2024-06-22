package main

import (
	"banking_sim/handlers"
	"banking_sim/middleware"
	"banking_sim/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Banking Simulation API
// @version 1.0
// @description This is a read-only API for simulating banking operations.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
func main() {
	models.InitDB()

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/api/v1")
	auth.POST("/auth", handlers.Login)
	auth.Use(middleware.Auth())
	{
		auth.GET("/accounts", handlers.GetAccounts)
		auth.GET("/instruments", handlers.GetInstruments)
		auth.GET("/transactions", handlers.GetTransactions)

		auth.GET("/transactions_categories", handlers.GetTransactionCategories)
		auth.GET("/instruments_categories", handlers.GetInstrumentsCategories)
	}

	router.Run(":8080")
}
