package controllers

import (
	"POSSystem/initializers"
	"POSSystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PurchaseHistoryAll(c *gin.Context) {
	// Fetch all invoices
	var invoices []models.Invoice
	if err := initializers.DB.Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching invoices"})
		return
	}
	var result []map[string]interface{}
	for _, invoice := range invoices {
		var products []models.TransactionProduct
		if err := initializers.DB.Where("invoice_id = ?", invoice.ID).
			Preload("Product").Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
			return
		}

		// Construct response for this invoice
		invoiceData := map[string]interface{}{
			"InvoiceNumber":     invoice.InvoiceNumber,
			"Date":              invoice.CreatedAt, // Or another date field if available
			"TotalAmount":       invoice.TotalAmount,
			"MoneyFromCustomer": invoice.MoneyFromCustomer,
			"Status":            invoice.PaymentStatus,
			"PaymentMethod":     invoice.PaymentMethod,
			"Products":          []map[string]interface{}{},
		}

		for _, product := range products {
			productData := map[string]interface{}{
				"ProductID":   product.ProductID,
				"Quantity":    product.Quantity,
				"ProductName": product.Product.Name,
				"Price":       product.Product.Price,
			}
			invoiceData["Products"] = append(invoiceData["Products"].([]map[string]interface{}), productData)
		}

		result = append(result, invoiceData)
	}

	c.JSON(http.StatusOK, result)
}
