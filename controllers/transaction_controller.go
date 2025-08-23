package controllers

import (
	"net/http"

	"github.com/ar-tiwari75/wealth-manager/models"
	"github.com/ar-tiwari75/wealth-manager/storage"
	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var transactions []models.Transaction
	if err := storage.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	// Convert []Transaction model to []TransactionResponse DTO
	var response []models.TransactionResponse
	for _, t := range transactions {
		response = append(response, models.TransactionResponse{
			ID:              t.ID,
			AssetType:       t.AssetType,
			AssetName:       t.AssetName,
			Quantity:        t.Quantity,
			Price:           t.Price,
			TransactionType: t.TransactionType,
			TransactionDate: t.TransactionDate,
			CreatedAt:       t.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transactions fetched successfully",
		"transactions": response,
	})
}

func CreateTransaction(c *gin.Context) {
	var req models.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert request to a transaction model
	transaction := models.Transaction{
		UserID:          c.GetUint("user_id"),
		AssetType:       req.AssetType,
		AssetName:       req.AssetName,
		Quantity:        req.Quantity,
		Price:           req.Price,
		TransactionType: req.TransactionType,
		TransactionDate: req.TransactionDate,
	}

	if err := storage.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Transaction created successfully",
		"transaction": transaction,
	})
}
