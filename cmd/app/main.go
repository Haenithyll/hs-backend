// @title HS Backend API
// @version 1.0
// @description This is the backend API for the HS project
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
	"hs-backend/internal/middleware"
	"hs-backend/internal/route"
)

func main() {
	config.LoadEnv()

	db := config.InitDB()

	port := config.GetEnv("PORT", "8080")
	httpLog := config.GetEnv("HTTP_LOG", "false") == "true"

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	if httpLog {
		r.Use(middleware.LoggerMiddleware())
	}
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Panicf("Failed to set trusted proxies: %v", err)
	}

	// Redis Publisher
	// pub, err := config.NewPublisher()
	// if err != nil {
	// 	log.Panicf("Failed to create publisher: %v", err)
	// }
	// defer pub.Close()

	env := config.GetEnvOrPanic("ENV")
	if env == "local" || env == "develop" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DocExpansion("none")))
	}
	route.RegisterRoutes(r, db)

	log.Printf("Server running on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Panicf("Failed to run server: %v", err)
	}
}
