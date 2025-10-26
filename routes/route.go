package routes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/easy-attend-serviceV3/app/modules"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
)

func Router(app *gin.Engine, mod *modules.Modules) {
	app.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	app.Use(otelgin.Middleware(mod.Conf.Svc.Config().AppName),
		// Middleware add trace id to response header
		func(ctx *gin.Context) {
			spanCtx := trace.SpanContextFromContext(ctx.Request.Context())
			if spanCtx.IsValid() {
				ctx.Header("X-Trace-ID", spanCtx.TraceID().String())
			}
			ctx.Next()
		},
	)

	app.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"*"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             false,
	}))

	// Debug middleware to log all requests
	app.Use(gin.Logger())

	// Debug middleware to log request details
	app.Use(func(ctx *gin.Context) {
		fmt.Printf("DEBUG: %s %s - Content-Type: %s\n",
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.GetHeader("Content-Type"))

		// Log request body for POST requests
		if ctx.Request.Method == "POST" {
			body, _ := ctx.GetRawData()
			fmt.Printf("DEBUG: Request Body: %s\n", string(body))
			// Reset body for controller to read
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		ctx.Next()
	})

	api(app.Group("/api/v1"), mod)
}
