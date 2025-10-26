package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware provides JWT authentication middleware
func AuthMiddleware(tokenManager *TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Check Bearer token format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		token := parts[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Token is required",
			})
			c.Abort()
			return
		}

		// Validate token (basic validation for now)
		claims, err := tokenManager.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_type", claims.UserType)

		c.Next()
	}
}

// RequireAuth is a helper function to create auth middleware
func RequireAuth() gin.HandlerFunc {
	tokenManager := NewTokenManager("your-secret-key-here") // TODO: Move to config
	return AuthMiddleware(tokenManager)
}
