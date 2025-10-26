package classroommember

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
		db          entitiesinf.ClassroomMemberEntity
		classroomDB entitiesinf.ClassroomEntity
		schoolDB    entitiesinf.SchoolEntity
		teacherDB   entitiesinf.TeacherEntity
		studentDB   entitiesinf.StudentEntity
	}
	Controller struct {
		tracer trace.Tracer
		svc    *Service
	}
)

type Options struct {
	// *configDTO.Config[Config]
	tracer      trace.Tracer
	db          entitiesinf.ClassroomMemberEntity
	classroomDB entitiesinf.ClassroomEntity
	schoolDB    entitiesinf.SchoolEntity
	teacherDB   entitiesinf.TeacherEntity
	studentDB   entitiesinf.StudentEntity
}

func New(db entitiesinf.ClassroomMemberEntity, classroomDB entitiesinf.ClassroomEntity, schoolDB entitiesinf.SchoolEntity, teacherDB entitiesinf.TeacherEntity, studentDB entitiesinf.StudentEntity) *Module {
	tracer := otel.Tracer("easy-attend-serviceV3.modules.classroom_member")
	svc := newService(&Options{
		// Config: conf,
		tracer:      tracer,
		db:          db,
		classroomDB: classroomDB,
		schoolDB:    schoolDB,
		teacherDB:   teacherDB,
		studentDB:   studentDB,
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
		classroomDB: opt.classroomDB,
		schoolDB:    opt.schoolDB,
		teacherDB:   opt.teacherDB,
		studentDB:   opt.studentDB,
	}
}

func newController(trace trace.Tracer, svc *Service) *Controller {
	return &Controller{
		tracer: trace,
		svc:    svc,
	}
}
