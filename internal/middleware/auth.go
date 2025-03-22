package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"hs-backend/internal/config"
	customError "hs-backend/internal/error"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := config.GetEnvOrPanic("SUPABASE_JWT_SECRET")

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customError.ErrorResponse{Error: "missing Authorization header"})
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customError.ErrorResponse{Error: "invalid Authorization header format"})
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, customError.ErrorResponse{Error: "invalid token: " + err.Error()})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customError.ErrorResponse{Error: "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customError.ErrorResponse{Error: "invalid token claims"})
			return
		}

		userId, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, customError.ErrorResponse{Error: "invalid token claims"})
			return
		}
		c.Set("user_id", userId)

		c.Next()
	}
}
