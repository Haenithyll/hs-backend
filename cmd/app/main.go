// @title HS Backend API
// @version 1.0
// @description This is the backend API for the HS project
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your token.

package main

import (
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "hs-backend/docs"
	"hs-backend/internal/config"
	"hs-backend/internal/routes"
)

func main() {
	config.LoadEnv()

	db := config.InitDB()

	port := config.GetEnv("PORT", "8080")

	r := gin.Default()
	r.SetTrustedProxies(nil)

	if config.GetEnv("ENV", "development") == "development" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	routes.RegisterRoutes(r, db)

	log.Printf("Server running on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
