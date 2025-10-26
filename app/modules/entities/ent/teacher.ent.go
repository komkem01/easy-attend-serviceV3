package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TeacherEntity struct {
	bun.BaseModel `bun:"table:teachers"`

	ID          uuid.UUID  `bun:"type:uuid,default:gen_random_uuid(),pk"`
	SchoolID    uuid.UUID  `bun:"type:uuid,notnull"`
	ClassroomID *uuid.UUID `bun:"type:uuid"` // ใช้ pointer เพื่อรองรับ NULL
	PrefixID    uuid.UUID  `bun:"type:uuid,notnull"`
	GenderID    uuid.UUID  `bun:"type:uuid,notnull"`
	FirstName   string     `bun:"type:varchar(100),notnull"`
	LastName    string     `bun:"type:varchar(100),notnull"`
	Email       string     `bun:"type:varchar(100),notnull,unique"`
	Password    string     `bun:"type:varchar(255),notnull"`
	Phone       string     `bun:"type:varchar(15)"`
	CreatedAt   time.Time  `bun:"type:timestamptz,default:current_timestamp,notnull"`
	UpdatedAt   time.Time  `bun:"type:timestamptz,default:current_timestamp,notnull"`
}
