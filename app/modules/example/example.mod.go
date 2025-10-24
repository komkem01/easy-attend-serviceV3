package example

import (
	entitiesinf "mcop/app/modules/entities/inf"

	configDTO "mcop/internal/config/dto"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Module struct {
	tracer trace.Tracer
	Svc    *Service
	Ctl    *Controller
}

func New(conf *configDTO.Config[Config], db entitiesinf.ExampleEntity) *Module {
	tracer := otel.Tracer("mcop.modules.example")
	svc := newService(&Options{
		Config: conf,
		tracer: tracer,
		db:     db,
	})
	return &Module{
		tracer: tracer,
		Svc:    svc,
		Ctl:    newController(tracer, svc),
	}
}
