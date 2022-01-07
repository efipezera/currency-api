package controllers

import (
	"net/http"

	"github.com/fplaraujo/currency-api/config"
	"github.com/fplaraujo/currency-api/models"
	"github.com/gin-gonic/gin"
)

//TODO: COMO EU SEI QUE A CONEXÃO FOI FECHADA ENTRE A API E O BANCO DE DADOS?

//FindTranfers brings all transfers made.
func FindTranfers(c *gin.Context) {
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	var transfers []models.Transfer
	config.DB.Find(&transfers)

	c.JSON(http.StatusOK, gin.H{
		"transfers": transfers,
	})
}

//CreateTransfer creates a transfer.
func CreateTransfer(c *gin.Context) {
	var input models.CreateTransferInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	transfer := models.Transfer{
		Value: input.Value,
	}
	config.DB.Create(&transfer)

	c.JSON(http.StatusCreated, gin.H{
		"amountTransferred": transfer,
	})
}

//FindTotalBalance sum all the transfers and return the total balance.
func FindTotalBalance(c *gin.Context) {
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	var totalBalance float64
	config.DB.Table("transfers").Select("sum(value)").Scan(&totalBalance)

	c.JSON(http.StatusOK, gin.H{
		"totalBalance": totalBalance,
	})
}

//FindBalanceByCurrency brings the total amount according to the currency chosen.
func FindBalanceByCurrency(c *gin.Context) {
	config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection()

	var totalBalance float64
	config.DB.Table("transfers").Select("sum(value)").Scan(&totalBalance)

	currency := c.Param("currency")
	var currencyValue float64
	config.DB.Table("currency").Select("currency_value").Where("currency = ?", currency).Scan(&currencyValue)

	finalValue := totalBalance / currencyValue

	c.JSON(http.StatusOK, gin.H{
		"currency": currency,
		"exchange": finalValue,
	})
}

// func calcExchange(exchangeValue float64) float64 {
// 	if exchangeValue <= 0 {
// 		panic("The exchange value can't be less than o equal to zero!")
// 	}

// 	spread := 0.04
// 	iof := 0.638
// 	exchangeTax := 0.16

// 	finalValue := exchangeValue + (exchangeValue * spread) + (exchangeValue * iof) + (exchangeValue * exchangeTax)
// 	return finalValue
// }
