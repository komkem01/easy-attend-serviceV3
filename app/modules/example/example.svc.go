package example

import (
	entitiesinf "mcop/app/modules/entities/inf"

	configDTO "mcop/internal/config/dto"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	tracer trace.Tracer
	db     entitiesinf.ExampleEntity // Database interface for object entities
}

type Config struct{}

type Options struct {
	*configDTO.Config[Config]
	tracer trace.Tracer
	db     entitiesinf.ExampleEntity // Database interface for object entities
}

func newService(opt *Options) *Service {
	return &Service{
		tracer: opt.tracer,
		db:     opt.db,
	}
}
