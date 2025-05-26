package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"hs-backend/internal/config"
	"hs-backend/internal/domain"
	"hs-backend/internal/repository"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := config.GetEnvOrPanic("SUPABASE_JWT_SECRET")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		apiKeyHeader := c.GetHeader("API-Key")

		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid Authorization header format"})
				return
			}
			tokenString := parts[1]

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid token signing method")
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid token: " + err.Error()})
				return
			}

			if !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid token"})
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid token claims"})
				return
			}

			userId, ok := claims["sub"].(string)
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid token claims"})
				return
			}
			c.Set("user_id", userId)

			c.Next()
			return
		}

		if apiKeyHeader != "" {
			sum := sha256.Sum256([]byte(apiKeyHeader))
			fingerprint := hex.EncodeToString(sum[:])

			userRepository := repository.NewUserRepository(config.GetDB())

			user, err := userRepository.FindOneByFingerprint(fingerprint)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid API key"})
				return
			}

			if user.ApiKeyHash == nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid API key"})
				return
			}

			if err := bcrypt.CompareHashAndPassword([]byte(*user.ApiKeyHash), []byte(apiKeyHeader)); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "invalid API key"})
				return
			}

			c.Set("user_id", (user.ID).String())
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "missing Authorization or API-Key header"})
	}
}
