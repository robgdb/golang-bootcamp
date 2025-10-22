package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robertbonadeo/go-bootcamp-project/config"
	"github.com/robertbonadeo/go-bootcamp-project/internal/db"
	"github.com/robertbonadeo/go-bootcamp-project/internal/handlers"
	"github.com/robertbonadeo/go-bootcamp-project/internal/middleware"
)

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	cfg := config.LoadConfig()

	db.InitDB(cfg)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	})

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// User routes
		protected.GET("/user", handlers.GetUser)
		protected.PUT("/user", handlers.UpdateUser)
		protected.POST("/user/change-password", handlers.ChangePassword)
		protected.POST("/refresh-token", handlers.RefreshToken)

		// Content routes
		protected.POST("/content", handlers.CreateContent)
		protected.GET("/content/:id", handlers.GetContent)
		protected.GET("/content", handlers.GetAllContent)
		protected.GET("/my-content", handlers.GetMyContent)
		protected.PUT("/content/:id", handlers.UpdateContent)
		protected.DELETE("/content/:id", handlers.DeleteContent)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
