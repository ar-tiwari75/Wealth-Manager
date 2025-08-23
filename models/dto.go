package models

import "time"

type ProfileResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTransactionRequest struct {
	AssetType       string    `json:"asset_type" binding:"required"`
	AssetName       string    `json:"asset_name" binding:"required"`
	Quantity        int       `json:"quantity" binding:"required"`
	Price           float64   `json:"price" binding:"required"`
	TransactionType string    `json:"transaction_type" binding:"required,oneof=buy sell"`
	TransactionDate time.Time `json:"transaction_date" binding:"required"`
}

type TransactionResponse struct {
	ID              uint      `json:"id"`
	AssetType       string    `json:"asset_type"`
	AssetName       string    `json:"asset_name"`
	Quantity        int       `json:"quantity"`
	Price           float64   `json:"price"`
	TransactionType string    `json:"transaction_type"`
	TransactionDate time.Time `json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at"`
}
