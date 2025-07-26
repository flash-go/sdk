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
	redisHostOptKey     = "/redis/host"
	redisPortOptKey     = "/redis/port"
	redisPasswordOptKey = "/redis/password"
	redisDbOptKey       = "/redis/db"
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
			config.Cfg.Get(redisHostOptKey),
			config.Cfg.Get(redisPortOptKey),
		),
		Password: config.Cfg.Get(redisPasswordOptKey),
		DB:       config.Cfg.GetInt(redisDbOptKey),
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
