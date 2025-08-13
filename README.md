# sdk

Software Development Kit written in Go for seamless integration with microservices built on the [Flash Framework](https://github.com/flash-go/flash). Designed to accelerate development, promote consistency, and simplify communication between distributed services.

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

### Get bytes from Base64 value by key

```go
package main

func main() {
	// Create config
	cfg := {...}

	// Get bytes from Base64 value by key
	value := cfg.GetBase64("host")
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

### Get env bool var 

```go
package main

import "github.com/flash-go/sdk/config"

func main() {
	value := config.GetEnvBool("key")
}
```

### Get env int var 

```go
package main

import "github.com/flash-go/sdk/config"

func main() {
	value := config.GetEnvInt("key")
}
```

### Get env bytes from Base64 var 

```go
package main

import "github.com/flash-go/sdk/config"

func main() {
	value := config.GetEnvBase64("key")
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

### Create insecure state service without auth token

```go
package main

import "github.com/flash-go/sdk/state"

func main() {
	// Full address (host:port) of the Consul agent (e.g., `localhost:8500`)
	consulAddress := "localhost:8501"

	// Create state service
	stateService := state.NewWithoutAuth(consulAddress)
}
```

### Create insecure state service with auth token

```go
package main

import "github.com/flash-go/sdk/state"

func main() {
	// Full address (host:port) of the Consul agent (e.g., `localhost:8500`)
	consulAddress := "localhost:8501"

	// Consul ACL token for authenticating requests to the Consul agent or server
	authToken := "secret"

	// Create state service
	stateService := state.NewWithInsecureAuth(
		&state.InsecureAuthConfig{
			Address: consulAddress,
			Token: authToken,
		},
	)
}
```

### Create secure state service with auth token

```go
package main

import "github.com/flash-go/sdk/state"

func main() {
	// Full address (host:port) of the Consul agent (e.g., `localhost:8500`)
	consulAddress := "localhost:8501"

	// Base64 CA certificate file used to verify the Consul server's TLS certificate
	ca := "secret"

	// Base64 client certificate file used for mTLS authentication with Consul
	cert := "secret"

	// Base64 private key corresponding to `CONSUL_CLIENT_CRT` for mTLS authentication
	key := "secret"

	// If set to `true`, disables TLS certificate verification (not recommended for production)
	insecureSkipVerify := false

	// Consul ACL token for authenticating requests to the Consul agent or server
	authToken := "secret"

	// Create state service
	stateService := state.NewWithSecureAuth(
		&state.SecureAuthConfig{
			Address: consulAddress,
			CAPem: ca,
			CertPEM: cert,
			KeyPEM: key,
			InsecureSkipVerify: insecureSkipVerify,
			Token: authToken,
		},
	)
}
```

### Create custom state

```go
package main

import (
	"github.com/flash-go/sdk/state"
	"github.com/hashicorp/consul/api"
)

func main() {
	// Create state service
	stateService := state.New(
		&api.Config{...},
	)
}
```

## telemetry

### Get telemetry KV keys

```go
package main

import "github.com/flash-go/sdk/telemetry"

func main() {
	telemetry.OtelCollectorGrpcOptKey
	telemetry.OtelCollectorCaCrtOptKey
	telemetry.OtelCollectorClientCrtOptKey
	telemetry.OtelCollectorClientKeyOptKey
}
```

### Create insecure grpc telemetry service

```go
package main

import "github.com/flash-go/sdk/telemetry"

func main() {
	// Create config
	cfg := {...}

	// Create telemetry service
	telemetryService := telemetry.NewInsecureGrpc(cfg)
}
```

### Create secure grpc telemetry service

```go
package main

import "github.com/flash-go/sdk/telemetry"

func main() {
	// Create config
	cfg := {...}

	// Create telemetry service
	telemetryService := telemetry.NewSecureGrpc(cfg)
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
