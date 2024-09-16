package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"size:100;not null;unique"`
	Password     string `gorm:"size:100;not null"`
	Role         string `gorm:"size:50;not null"`
	Salary       int    `gorm:"not null"`
	VacationDays int    `gorm:"not null"`
}
