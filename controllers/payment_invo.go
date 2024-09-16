package controllers

import (
	"POSSystem/initializers"
	"POSSystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateInvoice handles creating a new invoice
func CreateInvoice(c *gin.Context) {
	var invoiceData struct {
		InvoiceNumber     uint    `json:"invoice_number"`
		TotalAmount       float64 `json:"total_amount"`
		MoneyFromCustomer float64 `json:"money_from_customer"`
		PaymentMethod     string  `json:"payment_method"`
		PaymentStatus     string  `json:"payment_status"`
		UserID            uint    `json:"user_id"`
		Products          []struct {
			ID       uint `json:"id"`
			Quantity int  `json:"quantity"`
		} `json:"products"`
	}

	if err := c.BindJSON(&invoiceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new invoice
	invoice := models.Invoice{
		InvoiceNumber:     invoiceData.InvoiceNumber,
		TotalAmount:       invoiceData.TotalAmount,
		MoneyFromCustomer: invoiceData.MoneyFromCustomer,
		PaymentMethod:     invoiceData.PaymentMethod,
		PaymentStatus:     invoiceData.PaymentStatus,
		UserID:            invoiceData.UserID,
	}

	// Save the invoice
	if err := initializers.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create transaction products
	for _, p := range invoiceData.Products {
		if p.ID == 0 || p.Quantity <= 0 {
			continue // Skip invalid product entries
		}
		transactionProduct := models.TransactionProduct{
			InvoiceID: invoice.ID, // Corrected to InvoiceID
			ProductID: p.ID,
			Quantity:  p.Quantity,
		}
		if err := initializers.DB.Create(&transactionProduct).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, invoice)
}
