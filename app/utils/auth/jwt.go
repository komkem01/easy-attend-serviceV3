package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TokenManager handles simple token operations
type TokenManager struct {
	secretKey     string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserType  string    `json:"user_type"` // "teacher", "student", etc.
	jwt.RegisteredClaims
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

// NewTokenManagerWithConfig creates a new token manager with custom durations
func NewTokenManagerWithConfig(secretKey string, accessExpiry, refreshExpiry time.Duration) *TokenManager {
	return &TokenManager{
		secretKey:     secretKey,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

// GenerateTokenPair generates JWT access and refresh tokens
func (tm *TokenManager) GenerateTokenPair(userID uuid.UUID, email, firstName, lastName, userType string) (*TokenPair, error) {
	now := time.Now()

	// Generate JWT access token
	accessClaims := &TokenClaims{
		UserID:    userID,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserType:  userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tm.accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "easy-attend-service",
			Subject:   userID.String(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(tm.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Generate JWT refresh token
	refreshClaims := &TokenClaims{
		UserID:    userID,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserType:  userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tm.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "easy-attend-service",
			Subject:   userID.String(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(tm.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    now.Add(tm.accessExpiry),
		TokenType:    "Bearer",
	}, nil
}

// ValidateToken validates a JWT token and returns claims
func (tm *TokenManager) ValidateToken(tokenString string) (*TokenClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("empty token")
	}

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Extract claims
	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	// Check if token is expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token is expired")
	}

	return claims, nil
}

// RefreshToken generates a new access token from a valid refresh token
func (tm *TokenManager) RefreshToken(refreshTokenString string) (*TokenPair, error) {
	// Validate the refresh token
	claims, err := tm.ValidateToken(refreshTokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Generate new token pair
	return tm.GenerateTokenPair(
		claims.UserID,
		claims.Email,
		claims.FirstName,
		claims.LastName,
		claims.UserType,
	)
}
