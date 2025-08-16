package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	AssetType string  `gorm:"type:varchar(50);not null"`
	AssetName string  `gorm:"type:varchar(100); not null"`
	Quantity  float64 `gorm:"not null"`
	AvgPrice  float64 `gorm:"not null"`
}
