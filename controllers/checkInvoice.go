package controllers

import (
	"POSSystem/initializers"
	"POSSystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnProduct(c *gin.Context) {
	invoiceID := c.Param("id")
	var invoice models.Invoice

	if err := initializers.DB.Where("invoice_number = ?", invoiceID).First(&invoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	// Fetch products related to this invoice and preload the product details
	var products []models.TransactionProduct
	if err := initializers.DB.Where("invoice_id = ?", invoice.ID).
		Preload("Product").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}

	// Construct a response with product details
	var result []map[string]interface{}
	for _, p := range products {
		result = append(result, map[string]interface{}{
			"ProductID":   p.ProductID,
			"Quantity":    p.Quantity,
			"ProductName": p.Product.Name,
			"Price":       p.Product.Price,
		})
	}

	c.JSON(http.StatusOK, result)
}
