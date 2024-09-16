package controllers

import (
	"POSSystem/initializers"
	"POSSystem/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReturnPayment(c *gin.Context) {
	invoiceNumberStr := c.Param("id")
	invoiceNumber, err := strconv.ParseUint(invoiceNumberStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice number"})
		return
	}

	var invoiceData struct {
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

	// Parse and bind JSON request body
	if err := c.BindJSON(&invoiceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if MoneyFromCustomer & TotalAmount is 0 so delete the invoice
	if invoiceData.MoneyFromCustomer == 0 && invoiceData.TotalAmount == 0 {
		if err := initializers.DB.Where("invoice_number = ?", invoiceNumber).Delete(&models.Invoice{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
		return
	}

	// Retrieve the existing invoice using InvoiceNumber
	var invoice models.Invoice
	if err := initializers.DB.Where("invoice_number = ?", invoiceNumber).First(&invoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	// Update the invoice details
	invoice.TotalAmount = invoiceData.TotalAmount
	invoice.MoneyFromCustomer = invoiceData.MoneyFromCustomer
	invoice.PaymentMethod = invoiceData.PaymentMethod
	invoice.PaymentStatus = invoiceData.PaymentStatus
	invoice.UserID = invoiceData.UserID

	if err := initializers.DB.Save(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete existing transaction products
	if err := initializers.DB.Where("invoice_id = ?", invoice.ID).Delete(&models.TransactionProduct{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add new transaction products
	for _, p := range invoiceData.Products {
		if p.ID == 0 || p.Quantity <= 0 {
			continue // Skip invalid product entries
		}

		transactionProduct := models.TransactionProduct{
			InvoiceID: invoice.ID,
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
