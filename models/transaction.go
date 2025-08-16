package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID          uint      `gorm:"not null"`
	User            User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	AssetType       string    `gorm:"type:varchar(50);not null"`
	AssetName       string    `gorm:"type:varchar(100);not null"`
	Quantity        int       `gorm:"not null"`
	Price           float64   `gorm:"not null"`
	TransactionType string    `gorm:"type:varchar(50);not null"`
	TransactionDate time.Time `gorm:"type:date;not null"`
}
