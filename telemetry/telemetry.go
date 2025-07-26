package telemetry

import (
	"context"
	"log"
	"time"

	"github.com/flash-go/flash/telemetry"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

// Create grpc telemetry service
func NewGrpc(service, otelCollectorGrpc string) telemetry.Telemetry {
	// Create trace exporter
	traceExporter, err := telemetry.NewTraceExporterOtlpGrpc(
		context.Background(),
		otlptracegrpc.WithEndpoint(otelCollectorGrpc),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create telemetry trace exporter: %v", err)
	}

	// Create metric exporter
	metricExporter, err := telemetry.NewMetricExporterPeriodicOtlpGrpc(
		30*time.Second, // interval
		10*time.Second, // timeout
		context.Background(),
		otlpmetricgrpc.WithEndpoint(otelCollectorGrpc),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create telemetry metric exporter: %v", err)
	}

	// Return telemetry service
	return telemetry.New(service, traceExporter, metricExporter)
}
