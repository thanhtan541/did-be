package telemetry

import (
	"context"
	"io"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer(serviceName string) func(context.Context) error {
	var output io.Writer = os.Stdout
	if os.Getenv("LOG_TO_FILE") == "true" {
		f, _ := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		output = f
	}

	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint(), stdouttrace.WithWriter(output))
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
