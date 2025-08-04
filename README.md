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

### Get service name

```go
package main

func main() {
	// Create config
	cfg := {...}

	// Get service name
	value := cfg.GetService()
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

### Set env map

```go
package main

func main() {
	// Create config
	cfg := {...}

	var envMap = map[string]string{
		"ENV_KEY": "CONSUL_KV_KEY",
	}

	// Set env map
	err := cfg.SetEnvMap(envMap)
}
```

## errors

### Types

| Type                  | Message             |
|-----------------------|---------------------|
| ErrBadRequest         | bad_request         |
| ErrUnauthorized       | unauthorized        |
| ErrForbidden          | forbidden           |
| ErrServiceUnavailable | service_unavailable |
| ErrNotFound           | not_found           |

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

### Get Postgres KV keys

```go
package main

import "github.com/flash-go/sdk/infra"

func main() {
	infra.PostgresHostOptKey
	infra.PostgresPortOptKey
	infra.PostgresUserOptKey
	infra.PostgresPasswordOptKey
	infra.PostgresDbOptKey
}
```

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

### Get Redis KV keys

```go
package main

import "github.com/flash-go/sdk/infra"

func main() {
	infra.RedisHostOptKey
	infra.RedisPortOptKey
	infra.RedisPasswordOptKey
	infra.RedisDbOptKey
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

### Get telemetry KV keys

```go
package main

import "github.com/flash-go/sdk/telemetry"

func main() {
	telemetry.OtelCollectorGrpcOptKey
}
```

### Create gRPC telemetry service

```go
package main

import "github.com/flash-go/sdk/telemetry"

func main() {
	// Create config
	cfg := {...}

	// Create telemetry service
	telemetryService := telemetry.NewGrpc(cfg)
}
```

## services

### users

#### Create JWT key

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

#### Use users middleware

```go
package main

import "github.com/flash-go/sdk/services/users"

func main() {
	// Create http client
	httpClient := {...}

	// Create http server
	httpServer := {...}

	// Set users service
	usersService := "users-service"

	// Create users middleware
	usersMiddleware := users.NewMiddleware(
		&users.MiddlewareConfig{
			UsersService: usersService,
			HttpClient:  httpClient,
		},
	)

	// Set roles
	roles := []string{"admin"}

	// Use auth middleware
	httpServer.AddRoute(
		http.MethodGet,
		"/",
		func(ctx server.ReqCtx) {
			ctx.WriteString("Hello")
		},
		usersMiddleware.Auth(
			// Restrict access based on roles (optional)
			users.WithAuthRolesOption(roles),
			// Disable MFA validation (optional)
			users.WithoutAuthMfaOption(),
		),
	)
}
```
