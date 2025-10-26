package base

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidationError represents validation errors
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// NotFoundError represents resource not found errors
type NotFoundError struct {
	Resource string `json:"resource"`
	ID       string `json:"id"`
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: %s", e.Resource, e.ID)
}

// ConflictError represents resource conflict errors (e.g., duplicate email)
type ConflictError struct {
	Resource string `json:"resource"`
	Value    string `json:"value"`
}

func (e ConflictError) Error() string {
	return fmt.Sprintf("%s already exists: %s", e.Resource, e.Value)
}

// HandleValidationError handles validation errors with proper HTTP status
func HandleValidationError(ctx *gin.Context, err ValidationError) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    "400",
		"message": "Validation failed",
		"error": gin.H{
			"field":   err.Field,
			"message": err.Message,
		},
		"data": nil,
	})
}

// HandleNotFoundError handles not found errors with proper HTTP status
func HandleNotFoundError(ctx *gin.Context, err NotFoundError) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    "400",
		"message": "Resource not found",
		"error": gin.H{
			"resource": err.Resource,
			"id":       err.ID,
		},
		"data": nil,
	})
}

// HandleConflictError handles conflict errors with proper HTTP status
func HandleConflictError(ctx *gin.Context, err ConflictError) {
	ctx.JSON(http.StatusConflict, gin.H{
		"code":    "409",
		"message": "Resource conflict",
		"error": gin.H{
			"resource": err.Resource,
			"value":    err.Value,
		},
		"data": nil,
	})
}

// Enhanced HandleError with custom error types
func HandleCustomError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case ValidationError:
		HandleValidationError(ctx, e)
	case NotFoundError:
		HandleNotFoundError(ctx, e)
	case ConflictError:
		HandleConflictError(ctx, e)
	default:
		// Fallback to original HandleError
		HandleError(ctx, err)
	}
}
