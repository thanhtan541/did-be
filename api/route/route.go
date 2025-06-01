package route

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("ping-handler")

func Ping(ctx *gin.Context) {
	// Extract request context (instrumented by otelgin)
	ctx_tracer := ctx.Request.Context()

	// Start a new span
	ctx_span, span := tracer.Start(ctx_tracer, "ping-handler")
	defer span.End()

	// Add some metadata to the span
	span.SetAttributes(attribute.String("custom.key", "custom value"))

	result := doSomething(ctx_span)
	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func doSomething(ctx context.Context) string {
	_, span := tracer.Start(ctx, "doSomething")
	defer span.End()

	// ... actual logic
	return "pong"
}
