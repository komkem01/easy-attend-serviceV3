package example

import (
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"

	configDTO "github.com/easy-attend-serviceV3/internal/config/dto"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Module struct {
	tracer trace.Tracer
	Svc    *Service
	Ctl    *Controller
}

func New(conf *configDTO.Config[Config], db entitiesinf.ExampleEntity) *Module {
	tracer := otel.Tracer("easy-attend-serviceV3.modules.example")
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
