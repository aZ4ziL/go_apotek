package controllers

import (
	"github.com/aZ4zil/go_apotek/handlers"
	"github.com/gin-gonic/gin"
)

func Controllers(router *gin.Engine) {
	router.GET("/login", handlers.Login)
	router.POST("/login", handlers.Login)
	router.GET("/register", handlers.Register)
	router.POST("/register", handlers.Register)

	// Index
	router.GET("", handlers.Index)

	// Drugs
	router.GET("/drugs", handlers.GetALlDrug)

	// Transactions
	router.GET("/transactions", handlers.GetAllTransactions)
	router.GET("/transactions/pay", handlers.CreateNewTransaction)
	router.POST("/transactions/pay", handlers.CreateNewTransaction)
}
