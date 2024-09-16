package models

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;not null;unique"`
	Phone     string `gorm:"size:20"`
	Address   string `gorm:"size:255"`
	ProductID uint   `gorm:"not null"`
}
