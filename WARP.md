# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

This is a Go Clean Architecture template demonstrating clean separation of concerns across layers. The project implements a translation service with three server types: REST API (Fiber), gRPC, and AMQP RPC (RabbitMQ). 

## Architecture

### Clean Architecture Layers

The codebase follows Uncle Bob's Clean Architecture with dependency inversion:

- **Inner Layer (Business Logic)**: `internal/usecase` and `internal/entity`
  - Contains pure business logic using only Go standard library
  - Interfaces define contracts for outer layer dependencies
  - Independent of frameworks, databases, and external services

- **Outer Layer (Infrastructure)**: `internal/controller`, `internal/repo`, `pkg`
  - Controllers handle HTTP/gRPC/AMQP requests
  - Repositories implement data persistence (PostgreSQL) and external APIs
  - Infrastructure packages provide reusable components

### Key Directories

- `cmd/app/main.go` - Application entry point, configuration initialization
- `internal/app/app.go` - Dependency injection container and server startup
- `config/config.go` - Configuration management via Viper (reads `settings.yaml`)
- `internal/controller/` - Server handlers (HTTP/gRPC/AMQP layers)
- `internal/usecase/` - Business logic layer with interfaces
- `internal/repo/` - Data access layer (database + external APIs)
- `internal/entity/` - Domain entities (models)
- `pkg/` - Reusable infrastructure components

### Dependency Injection Pattern

Dependencies flow inward through constructor injection:
```
HTTP Controller -> UseCase -> Repository (Postgres/WebAPI)
```

All layer boundaries crossed via interfaces, enabling easy testing and swapping implementations.

## Development Commands

### Local Development Setup
```bash
# Start PostgreSQL and RabbitMQ
make compose-up

# Run application with migrations
make run
```

### Building and Running
```bash
# Install development dependencies
make bin-deps

# Run with automatic migrations
make run

# Build dependencies (tidy + verify)
make deps

# Check for security vulnerabilities
make deps-audit
```

### Code Generation
```bash
# Generate Swagger documentation
make swag-v1

# Generate gRPC code from proto files
make proto-v1

# Generate mocks for testing
make mock
```

### Testing
```bash
# Run unit tests with coverage
make test

# Run integration tests
make integration-test

# Full docker stack integration tests
make compose-up-integration-test
```

### Code Quality
```bash
# Format code (gofumpt + gci)
make format

# Run Go linters
make linter-golangci

# Run Dockerfile linter
make linter-hadolint

# Run dotenv linter
make linter-dotenv

# Complete pre-commit workflow
make pre-commit
```

### Database Migrations
```bash
# Create new migration
make migrate-create migration_name

# Run migrations
make migrate-up
```

## Adding New Features

### Adding a New Entity/UseCase
1. Create entity in `internal/entity/`
2. Define interfaces in `internal/usecase/contracts.go`
3. Implement usecase in `internal/usecase/[feature]/`
4. Create repository interfaces in `internal/repo/contracts.go`
5. Implement repositories in `internal/repo/persistent/` and `internal/repo/webapi/`
6. Wire dependencies in `internal/app/app.go`

### Adding New API Endpoints
1. Add routes in `internal/controller/http/v[X]/router.go`
2. Implement handlers in `internal/controller/http/v[X]/[feature].go`
3. Add Swagger annotations for documentation
4. Regenerate docs with `make swag-v1`

### API Versioning
- HTTP: Add `internal/controller/http/v2/` and register in router
- gRPC: Add `internal/controller/grpc/v2/` and `docs/proto/v2/`
- AMQP: Add `internal/controller/amqp_rpc/v2/`

## Configuration

Configuration is managed via:
- `settings.yaml` - Main configuration file
- Environment variables (override settings.yaml values)
- `config/config.go` - Configuration struct definitions

Key configuration sections:
- `app` - Application name and version
- `http` - HTTP server settings
- `pg` - PostgreSQL connection settings
- `grpc` - gRPC server settings  
- `rmq` - RabbitMQ settings
- `metrics` - Prometheus metrics
- `swagger` - API documentation

## Testing Strategy

- **Unit Tests**: Mock dependencies using generated interfaces
- **Integration Tests**: Use Docker containers for real database/services
- **Mocking**: Auto-generated mocks via `go generate` and mockgen
- **Coverage**: Atomic coverage mode for race condition detection

## Key Technologies

- **Web Framework**: Fiber (HTTP), gRPC, AMQP RPC
- **Database**: PostgreSQL with pgx driver
- **Query Builder**: Squirrel for SQL generation
- **Migrations**: golang-migrate for schema versioning
- **Logging**: zerolog for structured logging
- **Validation**: go-playground/validator
- **Documentation**: Swagger/OpenAPI auto-generation
- **Metrics**: Prometheus integration
- **Testing**: Testify framework with auto-generated mocks

## Service Access Points

When running with `make compose-up-all`:
- REST API: http://localhost:8080 or http://app.lvh.me
- gRPC: tcp://localhost:8081 or tcp://grpc.lvh.me:8081  
- Health Check: http://localhost:8080/healthz
- Metrics: http://localhost:8080/metrics
- Swagger UI: http://localhost:8080/swagger
- PostgreSQL: localhost:5432
- RabbitMQ UI: http://localhost:15672 (guest/guest)