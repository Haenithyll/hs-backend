package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() *gorm.DB {
	once.Do(func() {
		LoadEnv()

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			GetEnvOrPanic("DB_HOST"),
			GetEnvOrPanic("DB_USER"),
			GetEnvOrPanic("DB_PASSWORD"),
			GetEnvOrPanic("DB_NAME"),
			GetEnvOrPanic("DB_PORT"),
			GetEnvOrPanic("DB_SSLMODE"),
		)

		logLevel := GetEnvAsInt("GORM_LOG_LEVEL", int(logger.Warn))

		var err error
		const maxRetries = 3
		const retryDelay = 2 // seconds

		for i := 1; i <= maxRetries; i++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
			})

			if err == nil {
				log.Println("📦 Connected to database")
				break
			}

			log.Printf("❌ Failed to connect to DB (attempt %d/%d): %v", i, maxRetries, err)
			if i < maxRetries {
				time.Sleep(time.Duration(retryDelay) * time.Second)
			}
		}

		if err != nil {
			panic("failed to connect to database after retries: " + err.Error())
		}
	})

	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		panic("database is not initialized")
	}
	return db
}
