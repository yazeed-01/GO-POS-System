package controllers

import (
	"POSSystem/initializers"
	"POSSystem/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardData struct {
	TotalRevenue      float64 `json:"totalRevenue"`
	TotalLoss         float64 `json:"totalLoss"`
	TotalProfit       float64 `json:"totalProfit"`
	MostSoldProductID uint    `json:"mostSoldProductId"`
}

func AdminDashboardData(c *gin.Context) {
	var totalRevenue, totalLoss, totalProfit float64
	var mostSoldProductID uint

	db := initializers.DB

	// Calculate total revenue
	if err := db.Model(&models.Invoice{}).Select("SUM(total_amount)").Row().Scan(&totalRevenue); err != nil {
		log.Println("Error calculating total revenue:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Calculate total loss
	if err := db.Model(&models.Invoice{}).Select("SUM(money_from_customer)").Row().Scan(&totalLoss); err != nil {
		log.Println("Error calculating total loss:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Calculate total profit
	totalProfit = totalRevenue - totalLoss

	// Find the most sold product ID
	if err := db.Model(&models.TransactionProduct{}).Select("product_id").Group("product_id").Order("SUM(quantity) DESC").Limit(1).Row().Scan(&mostSoldProductID); err != nil {
		log.Println("Error finding most sold product ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, DashboardData{
		TotalRevenue:      totalRevenue,
		TotalLoss:         totalLoss,
		TotalProfit:       totalProfit,
		MostSoldProductID: mostSoldProductID,
	})
}
