package student

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
		tracer      trace.Tracer
		db          entitiesinf.StudentEntity
		dbSchool    entitiesinf.SchoolEntity
		dbClassroom entitiesinf.ClassroomEntity
		dbPrefix    entitiesinf.PrefixEntity
		dbGender    entitiesinf.GenderEntity
	}
	Controller struct {
		tracer trace.Tracer
		svc    *Service
	}
)

type Options struct {
	// *configDTO.Config[Config]
	tracer      trace.Tracer
	db          entitiesinf.StudentEntity
	dbSchool    entitiesinf.SchoolEntity
	dbClassroom entitiesinf.ClassroomEntity
	dbPrefix    entitiesinf.PrefixEntity
	dbGender    entitiesinf.GenderEntity
}

func New(db entitiesinf.StudentEntity, dbSchool entitiesinf.SchoolEntity, dbClassroom entitiesinf.ClassroomEntity, dbPrefix entitiesinf.PrefixEntity, dbGender entitiesinf.GenderEntity) *Module {
	tracer := otel.Tracer("easy-attend-serviceV3.modules.student")
	svc := newService(&Options{
		// Config: conf,
		tracer:      tracer,
		db:          db,
		dbSchool:    dbSchool,
		dbClassroom: dbClassroom,
		dbPrefix:    dbPrefix,
		dbGender:    dbGender,
	})
	return &Module{
		Svc: svc,
		Ctl: newController(tracer, svc),
	}
}

func newService(opt *Options) *Service {
	return &Service{
		tracer:      opt.tracer,
		db:          opt.db,
		dbSchool:    opt.dbSchool,
		dbClassroom: opt.dbClassroom,
		dbPrefix:    opt.dbPrefix,
		dbGender:    opt.dbGender,
	}
}

func newController(trace trace.Tracer, svc *Service) *Controller {
	return &Controller{
		tracer: trace,
		svc:    svc,
	}
}
