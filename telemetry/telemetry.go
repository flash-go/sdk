package telemetry

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc/credentials"

	"github.com/flash-go/flash/telemetry"
	"github.com/flash-go/sdk/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

const (
	OtelCollectorGrpcOptKey      = "/telemetry/collector/grpc"
	OtelCollectorCaCrtOptKey     = "/telemetry/collector/certs/ca.crt"
	OtelCollectorClientCrtOptKey = "/telemetry/collector/certs/client.crt"
	OtelCollectorClientKeyOptKey = "/telemetry/collector/certs/client.key"
)

// Create insecure grpc telemetry service
func NewInsecureGrpc(config config.Config) telemetry.Telemetry {
	// Get grpc endpoint
	endpoint := config.Get(OtelCollectorGrpcOptKey)

	// Create trace exporter without TLS
	traceExporter, err := telemetry.NewTraceExporterOtlpGrpc(
		context.Background(),
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create telemetry trace exporter: %v", err)
	}

	// Create metric exporter without TLS
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

	// Create telemetry service
	return telemetry.New(
		config.GetService(),
		traceExporter,
		metricExporter,
	)
}

// Create secure grpc telemetry service
func NewSecureGrpc(config config.Config) telemetry.Telemetry {
	// Get grpc endpoint
	endpoint := config.Get(OtelCollectorGrpcOptKey)

	// Get certs
	ca_crt := config.GetBase64(OtelCollectorCaCrtOptKey)
	client_crt := config.GetBase64(OtelCollectorClientCrtOptKey)
	client_key := config.GetBase64(OtelCollectorClientKeyOptKey)

	// Create CertPool
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(ca_crt) {
		log.Fatalf("failed to append CA certificate")
	}

	// Parse a public/private client key
	clientCert, err := tls.X509KeyPair(client_crt, client_key)
	if err != nil {
		log.Fatalf("failed to load client certificate/key: %v", err)
	}

	// New TLS credentials
	creds := credentials.NewTLS(
		// TLS config for mTLS
		&tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
			// Check server name
			ServerName: strings.Split(endpoint, ":")[0],
		},
	)

	// Create trace exporter with TLS
	traceExporter, err := telemetry.NewTraceExporterOtlpGrpc(
		context.Background(),
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithTLSCredentials(creds),
	)
	if err != nil {
		log.Fatalf("failed to create telemetry trace exporter: %v", err)
	}

	// Create metric exporter with TLS
	metricExporter, err := telemetry.NewMetricExporterPeriodicOtlpGrpc(
		30*time.Second, // interval
		10*time.Second, // timeout
		context.Background(),
		otlpmetricgrpc.WithEndpoint(endpoint),
		otlpmetricgrpc.WithTLSCredentials(creds),
	)
	if err != nil {
		log.Fatalf("failed to create telemetry metric exporter: %v", err)
	}

	// Create telemetry service
	return telemetry.New(
		config.GetService(),
		traceExporter,
		metricExporter,
	)
}
