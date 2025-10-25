package classroom

import (
	entitiesinf "github.com/easy-attend-serviceV3/app/modules/entities/inf"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Module struct {
	Svc *Service
	Ctl *Controller
}
type (
	Service struct {
		tracer trace.Tracer
		db     entitiesinf.ClassroomEntity
	}
	Controller struct {
		tracer trace.Tracer
		svc    *Service
	}
)

type Options struct {
	// *configDTO.Config[Config]
	tracer trace.Tracer
	db     entitiesinf.ClassroomEntity
}

func New(db entitiesinf.ClassroomEntity) *Module {
	tracer := otel.Tracer("easy-attend-serviceV3.modules.classroom")
	svc := newService(&Options{
		// Config: conf,
		tracer: tracer,
		db:     db,
	})
	return &Module{
		Svc: svc,
		Ctl: newController(tracer, svc),
	}
}

func newService(opt *Options) *Service {
	return &Service{
		tracer: opt.tracer,
		db:     opt.db,
	}
}

func newController(trace trace.Tracer, svc *Service) *Controller {
	return &Controller{
		tracer: trace,
		svc:    svc,
	}
}
