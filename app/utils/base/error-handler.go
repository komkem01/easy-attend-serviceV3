package base

import (
	"github.com/easy-attend-serviceV3/config/i18n"
	"github.com/gin-gonic/gin"
)

type ResponseFunction func(ctx *gin.Context, message string, data any, params ...map[string]string) error

type ErrorMapping struct {
	MessageKey   string
	ResponseFunc ResponseFunction
}

// Use error.Error() as key of map for error mappings
var errorMappings = map[string]ErrorMapping{
	i18n.BadRequest:        {i18n.BadRequest, BadRequest},
	i18n.InternalServerError: {i18n.InternalServerError, InternalServerError},
}

func HandleError(ctx *gin.Context, err error) {
	if err != nil {
		// Use map to lookup error mapping
		if mapping, ok := errorMappings[err.Error()]; ok {
			mapping.ResponseFunc(ctx, mapping.MessageKey, nil)
			return
		}

		// If error doesn't match any mapping, return the original error
		InternalServerError(ctx, err.Error(), nil)
	}
}
