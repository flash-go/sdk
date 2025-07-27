package telemetry

import (
	"context"
	"log"
	"time"

	"github.com/flash-go/flash/telemetry"
	"github.com/flash-go/sdk/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

const (
	OtelCollectorGrpcOptKey = "/telemetry/collector/grpc"
)

// Create grpc telemetry service
func NewGrpc(config config.Config) telemetry.Telemetry {
	// Get grpc endpoint
	endpoint := config.Get(OtelCollectorGrpcOptKey)

	// Create trace exporter
	traceExporter, err := telemetry.NewTraceExporterOtlpGrpc(
		context.Background(),
		otlptracegrpc.WithEndpoint(endpoint),
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
		otlpmetricgrpc.WithEndpoint(endpoint),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create telemetry metric exporter: %v", err)
	}

	// Return telemetry service
	return telemetry.New(
		config.GetService(),
		traceExporter,
		metricExporter,
	)
}
