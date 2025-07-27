package infra

import (
	"fmt"
	"log"

	"github.com/flash-go/flash/telemetry"
	"github.com/flash-go/sdk/config"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

const (
	RedisHostOptKey     = "/redis/host"
	RedisPortOptKey     = "/redis/port"
	RedisPasswordOptKey = "/redis/password"
	RedisDbOptKey       = "/redis/db"
)

type RedisClientConfig struct {
	Cfg       config.Config
	Telemetry telemetry.Telemetry
}

func NewRedisClient(config *RedisClientConfig) *redis.Client {
	// Create options
	options := &redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			config.Cfg.Get(RedisHostOptKey),
			config.Cfg.Get(RedisPortOptKey),
		),
		Password: config.Cfg.Get(RedisPasswordOptKey),
		DB:       config.Cfg.GetInt(RedisDbOptKey),
	}

	// Create client
	client := redis.NewClient(options)

	// Use otel instrument tracing
	if config.Telemetry != nil {
		if err := redisotel.InstrumentTracing(
			client,
			redisotel.WithTracerProvider(
				config.Telemetry.TraceProvider(),
			),
		); err != nil {
			log.Fatalf("otel instrument tracing error: %v", err)
		}
	}

	return client
}
