package routes

import (
	"github.com/ar-tiwari75/wealth-manager/controllers"
	"github.com/ar-tiwari75/wealth-manager/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "API is healthy",
		})
	})

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleWare())
	{
		protected.GET("/profile", controllers.GetProfile)
		protected.GET("/transactions", controllers.GetTransactions)
		protected.POST("/transactions", controllers.CreateTransaction)
	}
}
