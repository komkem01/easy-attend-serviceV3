package entitiesinf

import (
	"context"

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
}

// teacher
type TeacherEntity interface {
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
