package controllers

import (
	"POSSystem/initializers"
	"POSSystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SellProductTable - Controller to fetch product details by ID
func SellProductTable(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	if err := initializers.DB.Where("ID = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    product.ID,
		"name":  product.Name,
		"price": product.Price,
	})
}
