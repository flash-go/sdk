# sdk

Reusable utilities written in Go for seamless integration with microservices built on the [Flash Framework](https://github.com/flash-go/flash). Designed to accelerate development, promote consistency, and simplify communication between distributed services.

## config

### Create config

```go
package main

import "github.com/flash-go/sdk/config"

func main() {
	// Create state service
	stateService := {...}

	// Set service name
	serviceName := "name"

	// Create config
	cfg := config.New(stateService, serviceName)
}
```

### Get string value by key

```go
package main

func main() {
	// Create config
	cfg := {...}

	// Get string value by key
	value := cfg.Get("host")
}
```

### Get int value by key

```go
package main

func main() {
	// Create config
	cfg := {...}

	// Get int value by key
	value := cfg.GetInt("host")
}
```

### Set value by key

```go
package main

func main() {
	// Create config
	cfg := {...}

	// Set value by key
	err := cfg.Set("key", "value")
}
```

## errors

### Types

| Type           | Message      |
|----------------|--------------|
| ErrBadRequest  | bad_request  |
| ErrUnauthorized| unauthorized |
| ErrForbidden   | forbidden    |

### Create error

```go
package main

import "github.com/flash-go/sdk/errors"

func main() {
	// Create error
	err := errors.New(errors.ErrUnauthorized, "invalid_credentials")
}
```

## infra

### Create Postgres client

```go
package main

import "github.com/flash-go/sdk/infra"

func main() {
	// Create config
	cfg := {...}

	// Create postgres client
	postgresClient := infra.NewPostgresClient(
		&infra.PostgresClientConfig{
			Cfg:        cfg,
			Telemetry:  nil,
			Migrations: nil,
		},
	)
}
```

### Create Redis client

```go
package main

import "github.com/flash-go/sdk/infra"

func main() {
	// Create config
	cfg := {...}

	// Create redis client
	redisClient := infra.NewRedisClient(
		&infra.RedisClientConfig{
			Cfg:       cfg,
			Telemetry: nil,
		},
	)
}
```

## logger

### Create console logger

```go
package main

import "github.com/flash-go/sdk/logger"

func main() {
	// Create logger service
	loggerService := logger.NewConsole()
}
```

## state

### Create state service

```go
package main

import "github.com/flash-go/sdk/state"

func main() {
	// Set service name
	serviceName := "name"

	// Create state service
	stateService := state.New(serviceName)
}
```

## telemetry

### Create gRPC telemetry service

```go
package main

import "github.com/flash-go/sdk/telemetry"

func main() {
	// Set service name
	serviceName := "name"

	// OTEL collector gRPC endpoint
	endpoint := "localhost:4317"

	// Create telemetry service
	telemetryService := telemetry.NewGrpc(
		serviceName,
		endpoint,
	)
}
```

## services

### users

Create JWT key

```go
package main

import "github.com/flash-go/sdk/services/users"

func main() {
	// Key size
	size := 64

	// Create key
	key := users.NewJwtKey(size)
}
```
