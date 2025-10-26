package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserContext holds user information from JWT token
type UserContext struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	UserType string    `json:"user_type"`
}

// GetUserFromContext extracts user information from Gin context
func GetUserFromContext(c *gin.Context) (*UserContext, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return nil, errors.New("user ID not found in context")
	}

	email, exists := c.Get("user_email")
	if !exists {
		return nil, errors.New("user email not found in context")
	}

	userType, exists := c.Get("user_type")
	if !exists {
		return nil, errors.New("user type not found in context")
	}

	// Convert userID to UUID
	var userUUID uuid.UUID
	switch v := userID.(type) {
	case uuid.UUID:
		userUUID = v
	case string:
		parsed, err := uuid.Parse(v)
		if err != nil {
			return nil, errors.New("invalid user ID format")
		}
		userUUID = parsed
	default:
		return nil, errors.New("invalid user ID type")
	}

	return &UserContext{
		UserID:   userUUID,
		Email:    email.(string),
		UserType: userType.(string),
	}, nil
}

// MustGetUserFromContext extracts user information and panics if not found
func MustGetUserFromContext(c *gin.Context) *UserContext {
	user, err := GetUserFromContext(c)
	if err != nil {
		panic("User not found in context: " + err.Error())
	}
	return user
}

// GetUserID extracts only user ID from context
func GetUserID(c *gin.Context) (uuid.UUID, error) {
	user, err := GetUserFromContext(c)
	if err != nil {
		return uuid.Nil, err
	}
	return user.UserID, nil
}

// MustGetUserID extracts user ID and panics if not found
func MustGetUserID(c *gin.Context) uuid.UUID {
	userID, err := GetUserID(c)
	if err != nil {
		panic("User ID not found in context: " + err.Error())
	}
	return userID
}
