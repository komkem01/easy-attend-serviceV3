package entitiesinf

import (
	"context"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/modules/entities/ent"

	"github.com/google/uuid"
)

// ObjectEntity defines the interface for object entity operations such as create, retrieve, update, and soft delete.
type ExampleEntity interface {
	CreateExample(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
	GetExampleByID(ctx context.Context, id uuid.UUID) (*ent.Example, error)
	UpdateExampleByID(ctx context.Context, id uuid.UUID, status ent.ExampleStatus) (*ent.Example, error)
	SoftDeleteExampleByID(ctx context.Context, id uuid.UUID) error
	ListExamplesByStatus(ctx context.Context, status ent.ExampleStatus) ([]*ent.Example, error)
}
type ExampleTwoEntity interface {
	CreateExampleTwo(ctx context.Context, userID uuid.UUID) (*ent.Example, error)
}

// student
type StudentEntity interface {
	CreateStudent(ctx context.Context, req *entitiesdto.StudentCreateRequest) (*ent.StudentEntity, error)
	GetListStudent(ctx context.Context, resp *entitiesdto.StudentListResponse) ([]*ent.StudentEntity, error)
	GetStudentByID(ctx context.Context, id uuid.UUID, resp *entitiesdto.StudentInfoResponse) (*ent.StudentEntity, error)
	UpdateStudent(ctx context.Context, id uuid.UUID, req *entitiesdto.StudentUpdateRequest) (*ent.StudentEntity, error)
	DeleteStudent(ctx context.Context, id uuid.UUID) error
}

// teacher
type TeacherEntity interface {
	CreateTeacher(ctx context.Context, req *entitiesdto.TeacherCreateRequest) (*ent.TeacherEntity, error)
	GetListTeacher(ctx context.Context) ([]*ent.TeacherEntity, error)
	GetByIDTeacher(ctx context.Context, id uuid.UUID) (*ent.TeacherEntity, error)
	GetTeacherByEmail(ctx context.Context, email string) (*ent.TeacherEntity, error)
	UpdateTeacher(ctx context.Context, id uuid.UUID, req *entitiesdto.TeacherUpdateRequest) (*ent.TeacherEntity, error)
	DeleteTeacher(ctx context.Context, id uuid.UUID) error
	CheckExistTeacher(ctx context.Context, id uuid.UUID) (bool, error)
}

// prefix
type PrefixEntity interface {
	GetListPrefix(ctx context.Context) ([]*ent.PrefixEntity, error)
	GetByIDPrefix(ctx context.Context, id uuid.UUID) (*ent.PrefixEntity, error)
}

// gender
type GenderEntity interface {
	GetListGender(ctx context.Context) ([]*ent.GenderEntity, error)
	GetByIDGender(ctx context.Context, id uuid.UUID) (*ent.GenderEntity, error)
}

// school
type SchoolEntity interface {
	GetListSchool(ctx context.Context) ([]*ent.SchoolEntity, error)
	GetByIDSchool(ctx context.Context, id uuid.UUID) (*ent.SchoolEntity, error)
	CreateSchool(ctx context.Context, name, address, phone string) (*ent.SchoolEntity, error)
	UpdateSchool(ctx context.Context, id uuid.UUID, name, address, phone string) (*ent.SchoolEntity, error)
	DeleteSchool(ctx context.Context, id uuid.UUID) error
	CheckExistSchool(ctx context.Context, id uuid.UUID) (bool, error)
}

// classroom
type ClassroomEntity interface {
	GetListClassroom(ctx context.Context) ([]*ent.ClassroomEntity, error)
	GetByIDClassroom(ctx context.Context, id uuid.UUID) (*ent.ClassroomEntity, error)
	CreateClassroom(ctx context.Context, schoolID uuid.UUID, name string) (*ent.ClassroomEntity, error)
	UpdateClassroom(ctx context.Context, id uuid.UUID, schoolID uuid.UUID, name string) (*ent.ClassroomEntity, error)
	DeleteClassroom(ctx context.Context, id uuid.UUID) error
	CheckExistClassroom(ctx context.Context, id uuid.UUID) (bool, error)
}

// classroom member
type ClassroomMemberEntity interface {
	CreateClassroomMember(ctx context.Context, req *entitiesdto.ClassroomMemberCreateRequest) (*ent.ClassroomMemberEntity, error)
	GetListClassroomMember(ctx context.Context, classroomID uuid.UUID) ([]*ent.ClassroomMemberEntity, error)
	GetAllClassroomMembers(ctx context.Context, limit int) ([]*ent.ClassroomMemberEntity, error)
	GetClassroomMemberByID(ctx context.Context, id uuid.UUID) (*ent.ClassroomMemberEntity, error)
	UpdateClassroomMember(ctx context.Context, id uuid.UUID, req *entitiesdto.ClassroomMemberUpdateRequest) (*ent.ClassroomMemberEntity, error)
	DeleteClassroomMember(ctx context.Context, id uuid.UUID) error
	CheckExistClassroomMember(ctx context.Context, id uuid.UUID) (bool, error)
	GetClassroomMembersByStudentID(ctx context.Context, studentID uuid.UUID) ([]*ent.ClassroomMemberEntity, error)
}

// Attendance
type AttendanceEntity interface {
	CreateAttendance(ctx context.Context, req *entitiesdto.AttendanceCreateRequest) (*ent.AttendanceEntity, error)
	GetListAttendance(ctx context.Context, classroomID uuid.UUID, date string) ([]*ent.AttendanceEntity, error)
	GetAllAttendance(ctx context.Context, limit int) ([]*ent.AttendanceEntity, error)
	GetAttendanceByID(ctx context.Context, id uuid.UUID) (*ent.AttendanceEntity, error)
	UpdateAttendance(ctx context.Context, id uuid.UUID, req *entitiesdto.AttendanceUpdateRequest) (*ent.AttendanceEntity, error)
	DeleteAttendance(ctx context.Context, id uuid.UUID) error
	CheckExistAttendance(ctx context.Context, id uuid.UUID) (bool, error)
	GetAttendanceByStudentID(ctx context.Context, studentID uuid.UUID, date string) (*ent.AttendanceEntity, error)
	GetAttendanceByClassroomAndDate(ctx context.Context, classroomID uuid.UUID, date string) ([]*ent.AttendanceEntity, error)
}
