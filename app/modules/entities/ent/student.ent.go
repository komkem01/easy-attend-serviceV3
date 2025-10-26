package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type StudentEntity struct {
	bun.BaseModel `bun:"table:students"`

	ID          uuid.UUID `bun:"type:uuid,default:gen_random_uuid(),pk"`
	SchoolID    uuid.UUID `bun:"type:uuid,notnull"`
	ClassroomID uuid.UUID `bun:"type:uuid"`
	PrefixID    uuid.UUID `bun:"type:uuid,notnull"`
	GenderID    uuid.UUID `bun:"type:uuid,notnull"`
	StudentCode string    `bun:"type:varchar(20),notnull"`
	FirstName   string    `bun:"type:varchar(100),notnull"`
	LastName    string    `bun:"type:varchar(100),notnull"`
	Phone       string    `bun:"type:varchar(15)"`
	CreatedAt   time.Time `bun:"type:timestamptz,default:current_timestamp,notnull"`
	UpdatedAt   time.Time `bun:"type:timestamptz,default:current_timestamp,notnull"`
}
