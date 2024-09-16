package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"size:255"`
	Price       float64 `gorm:"not null"`
	Quantity    uint    `gorm:"not null"`
	Stock       bool    `gorm:"not null"`
	SupplierID  uint    `gorm:"not null"`
}
