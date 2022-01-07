package main

import (
	"github.com/fplaraujo/currency-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	transferRoutes := router.Group("api/transfer")
	{
		transferRoutes.GET("/", controllers.FindTranfers)
		transferRoutes.POST("/", controllers.CreateTransfer)
	}

	balanceRoutes := router.Group("api/balance")
	{
		balanceRoutes.GET("/", controllers.FindTotalBalance)
		balanceRoutes.GET("/:currency", controllers.FindBalanceByCurrency)
	}

	router.Run()
}
