package infra

import (
	"fmt"
	"log"

	"github.com/flash-go/flash/telemetry"
	"github.com/flash-go/sdk/config"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

const (
	PostgresHostOptKey     = "/postgres/host"
	PostgresPortOptKey     = "/postgres/port"
	PostgresUserOptKey     = "/postgres/user"
	PostgresPasswordOptKey = "/postgres/password"
	PostgresDbOptKey       = "/postgres/db"
)

type PostgresClientConfig struct {
	Cfg        config.Config
	Telemetry  telemetry.Telemetry
	Migrations []*gormigrate.Migration
}

func NewPostgresClient(config *PostgresClientConfig) *gorm.DB {
	// Create DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.Get(PostgresHostOptKey),
		config.Cfg.Get(PostgresPortOptKey),
		config.Cfg.Get(PostgresUserOptKey),
		config.Cfg.Get(PostgresPasswordOptKey),
		config.Cfg.Get(PostgresDbOptKey),
	)

	// Connect to server
	client, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}

	// Use OpenTelemetry tracing plugin
	if config.Telemetry != nil {
		if err := client.Use(
			tracing.NewPlugin(
				tracing.WithTracerProvider(
					config.Telemetry.TraceProvider(),
				),
			),
		); err != nil {
			log.Fatalf("failed to use OpenTelemetry tracing plugin: %v", err)
		}
	}

	// Run migrations
	if len(config.Migrations) > 0 {
		mirgations := gormigrate.New(client, gormigrate.DefaultOptions, config.Migrations)
		if err := mirgations.Migrate(); err != nil {
			log.Fatalf("could not migrate: %v", err)
		}
	}

	return client
}
