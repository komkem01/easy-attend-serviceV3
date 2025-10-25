package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ClassroomEntity struct {
	bun.BaseModel `bun:"table:classrooms"`

	ID        uuid.UUID `bun:"type:uuid,default:gen_random_uuid(),pk"`
	SchoolID  uuid.UUID `bun:"type:uuid,notnull"`
	Name      string    `bun:"type:varchar(255),notnull"`
	CreatedAt time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`
}
