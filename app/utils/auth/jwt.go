package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// TokenManager handles simple token operations
type TokenManager struct {
	secretKey     string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

// TokenClaims represents simple token claims
type TokenClaims struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserType  string    `json:"user_type"` // "teacher", "student", etc.
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
}

// TokenPair represents access and refresh tokens
type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	TokenType    string    `json:"token_type"`
}

// NewTokenManager creates a new token manager
func NewTokenManager(secretKey string) *TokenManager {
	return &TokenManager{
		secretKey:     secretKey,
		accessExpiry:  24 * time.Hour,     // 24 hours
		refreshExpiry: 7 * 24 * time.Hour, // 7 days
	}
}

// GenerateTokenPair generates simple access and refresh tokens
func (tm *TokenManager) GenerateTokenPair(userID uuid.UUID, email, firstName, lastName, userType string) (*TokenPair, error) {
	now := time.Now()

	// Generate random tokens
	accessToken, err := tm.generateRandomToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := tm.generateRandomToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    now.Add(tm.accessExpiry),
		TokenType:    "Bearer",
	}, nil
}

// generateRandomToken generates a secure random token
func (tm *TokenManager) generateRandomToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// ValidateToken validates a simple token (for now, just check if not empty)
// In production, you'd store tokens in Redis/DB and validate against them
func (tm *TokenManager) ValidateToken(tokenString string) (*TokenClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("empty token")
	}

	// In a real implementation, you'd:
	// 1. Look up the token in Redis/Database
	// 2. Check if it's expired
	// 3. Return associated user claims

	// For now, return empty claims (you'd implement proper validation)
	return &TokenClaims{}, nil
}
