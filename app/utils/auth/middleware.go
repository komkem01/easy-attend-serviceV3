package auth

import (
	"net/http"
	"strings"
	"time"

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

		// Validate JWT token
		claims, err := tokenManager.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		// Set user information in context for use in handlers
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_type", claims.UserType)
		c.Set("user_first_name", claims.FirstName)
		c.Set("user_last_name", claims.LastName)
		c.Set("user_claims", claims)

		c.Next()
	}
}

// RequireAuth is a helper function to create auth middleware
func RequireAuth() gin.HandlerFunc {
	tokenManager := NewTokenManager("your-secret-key-here") // TODO: Move to config
	return AuthMiddleware(tokenManager)
}

// RequireAuthWithConfig creates auth middleware with custom config
func RequireAuthWithConfig(secretKey string, accessExpiry, refreshExpiry int) gin.HandlerFunc {
	tokenManager := NewTokenManagerWithConfig(
		secretKey,
		time.Duration(accessExpiry)*time.Hour,
		time.Duration(refreshExpiry)*time.Hour,
	)
	return AuthMiddleware(tokenManager)
}

// RequireRole creates middleware that requires specific user role
func RequireRole(tokenManager *TokenManager, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// First run auth middleware
		AuthMiddleware(tokenManager)(c)

		// If auth middleware aborted, return
		if c.IsAborted() {
			return
		}

		// Check user role
		userType, exists := c.Get("user_type")
		if !exists || userType != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "Insufficient permissions",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
