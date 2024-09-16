package models

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	ID                uint    `gorm:"primaryKey"`
	InvoiceNumber     uint    `gorm:"not null;autoIncrement"`
	TotalAmount       float64 `gorm:"not null"`
	MoneyFromCustomer float64 `gorm:"not null"`
	PaymentMethod     string  `gorm:"not null"`
	PaymentStatus     string  `gorm:"not null"`
	UserID            uint    `gorm:"not null"`
}

type TransactionProduct struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	InvoiceID uint
	ProductID uint
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
}
