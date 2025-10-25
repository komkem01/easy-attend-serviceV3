package example

import (
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"

	configDTO "github.com/easy-attend-serviceV3/internal/config/dto"

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
