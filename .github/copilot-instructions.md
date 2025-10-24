# Copilot Instructions

## Project Overview

This is a **Go microservice template** (MCOP - Microservice Cooperative Platform) that provides a robust, production-ready foundation for building scalable microservices. The project follows clean architecture principles with dependency injection and modular design patterns.

## Architecture & Tech Stack

### Core Technologies
- **Language**: Go 1.24.2
- **Web Framework**: Gin (HTTP/1.1 & HTTP/2 support)
- **Database**: PostgreSQL with Bun ORM
- **Message Queue**: Apache Kafka with SASL/TLS
- **Cache**: Redis with JSON support
- **Observability**: OpenTelemetry (traces, metrics, logs)
- **Logging**: Zap with structured logging
- **Configuration**: Environment-based with reflection
- **Build**: Docker multi-stage builds
- **Testing**: Built-in test framework

### Project Structure
```
/app                    # Application layer
  /console             # CLI commands
  /modules             # Business modules
    /entities          # Data entities (Bun models)
    /example           # Example module implementation
    /example-two       # Additional example module
    /net               # Network utilities
  /utils               # Shared utilities
/config                # Configuration definitions
/database              # Database migrations
/internal              # Internal services (DI layer)
  /cmd                 # Command definitions
  /config              # Configuration service
  /database            # Database service
  /http                # HTTP server
  /kafka               # Kafka service
  /log                 # Logging service
  /otel                # OpenTelemetry collector
  /redis               # Redis service
/routes                # HTTP routes definition
```

## Development Guidelines

### 1. Module Development Pattern

When creating new modules, follow the established pattern:

```go
// Module structure
type Module struct {
    Svc *Service      // Business logic service
    Ctl *Controller   // HTTP controllers
}

// Service with dependencies
type Service struct {
    tracer trace.Tracer
    // other dependencies
}

// Controller with service dependency
type Controller struct {
    tracer trace.Tracer
    svc    *Service
    cli    *httpx.Client  // for external HTTP calls
}
```

### 2. Code Generation Commands

Use the built-in generators for consistency:

```bash
# Database migrations
go run . db init     # Initialize migration tables
go run . db create   # Create new migration
go run . db migrate  # Apply migrations
go run . db status   # Check migration status
```

### 3. Configuration Management

- Use environment variables with structured configuration
- Follow the naming pattern: `{MODULE}_{FIELD}` → `DATABASE_SQL__HOST`
- Support nested structures with double underscores `__`
- Use struct tags for validation: `conf:"required"`

```go
type Config struct {
    DatabaseHost string `conf:"required"`
    Port         int    `default:"8080"`
    Debug        bool   `default:"false"`
}
```

### 4. Database Patterns

- Use Bun ORM for database operations
- Implement repository pattern through interfaces
- Support multiple database connections
- Always use context for operations
- Implement proper transaction handling

```go
// Entity interface pattern
type ExampleEntity interface {
    CreateExample(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
    GetExampleByID(ctx context.Context, id uuid.UUID) (*ent.Example, error)
    ListExamplesByStatus(ctx context.Context, status ent.ExampleStatus) ([]*ent.Example, error)
    UpdateExampleByID(ctx context.Context, id uuid.UUID, status ent.ExampleStatus) (*ent.Example, error)
    SoftDeleteExampleByID(ctx context.Context, id uuid.UUID) error
}
```

### 5. HTTP API Conventions

- Use Gin for routing
- Implement proper request/response DTOs
- Add OpenTelemetry tracing to all endpoints
- Use middleware for cross-cutting concerns
- Follow RESTful conventions

```go
// Controller pattern
func (c *Controller) Create(ctx *gin.Context) {
    span := trace.SpanFromContext(ctx.Request.Context())
    span.SetAttributes(attribute.String("operation", "create"))

    var req CreateRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Business logic...
    ctx.JSON(201, response)
}
```

### 6. Observability

- Add tracing to all operations using OpenTelemetry
- Use structured logging with context
- Implement proper error handling and logging
- Add metrics for business operations

```go
// Tracing pattern
func (s *Service) SomeOperation(ctx context.Context) error {
    ctx, span := s.tracer.Start(ctx, "service.some_operation")
    defer span.End()

    span.SetAttributes(
        attribute.String("operation.type", "business_logic"),
        attribute.String("user.id", userID),
    )

    // Operation logic...
    return nil
}
```

### 7. Error Handling

- Use structured error responses
- Implement proper error logging
- Add trace context to errors
- Use appropriate HTTP status codes

### 8. Kafka Integration

- Use the built-in Kafka service for messaging
- Implement proper serialization (JSON)
- Handle consumer groups properly
- Use SSL/TLS for production

```go
// Producer pattern
err := s.kafka.ProduceJSON(ctx, "topic-name", key, messageData)

// Consumer pattern
handler := func(ctx context.Context, message *sarama.ConsumerMessage) error {
    // Process message
    return nil
}
closeFn, err := s.kafka.ConsumerGroup(ctx, "group-id", []string{"topic"}, handler)
```

## Best Practices

### Code Style
- Follow Go conventions and idioms
- Use meaningful variable and function names
- Keep functions small and focused
- Use dependency injection through constructors
- Implement interfaces for testability

### Security
- Use SSL/TLS for all external communications
- Validate all inputs
- Use proper authentication/authorization
- Sanitize database queries (Bun handles this)

### Performance
- Use connection pooling for databases
- Implement proper caching strategies
- Use background contexts appropriately
- Handle graceful shutdowns

### Testing
- Write unit tests for business logic
- Use dependency injection for testability
- Mock external dependencies
- Test error scenarios

## Environment Variables

Key environment variables to configure:

```bash
# Application
APP_NAME=your-service-name
APP_ENV=development|production
PORT=8080
DEBUG=true|false

# Database
DATABASE_SQL__HOST=localhost
DATABASE_SQL__DATABASE=dbname
DATABASE_SQL__USERNAME=username
DATABASE_SQL__PASSWORD=password

# Redis (optional)
DATABASE_REDIS__ADDR=localhost:6379
DATABASE_REDIS__PASSWORD=
DATABASE_REDIS__DB=0

# Kafka
KAFKA_BROKERS=localhost:9092
KAFKA_CA_CERT_PATH=path/to/ca.crt
KAFKA_CERT_PATH=path/to/cert.crt
KAFKA_KEY_PATH=path/to/key.key

# OpenTelemetry
OTEL_ENABLE=true
OTEL_COLLECTOR_ENDPOINT=localhost:4317

# HTTP
HTTP_JSON_NAMING=snake_case|camelCase
```

## Common Tasks

### Adding New Dependencies
1. Update `go.mod` with `go get package`
2. Add to vendor with `go mod vendor`
3. Import and use in your modules

## Troubleshooting

- Check logs for structured error messages
- Use trace IDs for request tracking
- Verify environment variables are set correctly
- Ensure database migrations are applied
- Check Kafka connectivity for messaging issues

This template provides a solid foundation for microservice development with production-ready features like observability, security, and scalability built-in.
