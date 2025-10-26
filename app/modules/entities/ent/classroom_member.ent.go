package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ClassroomMemberEntity struct {
	bun.BaseModel `bun:"table:classroom_members"`

	ID          uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	ClassroomID uuid.UUID `bun:"classroom_id,type:uuid,notnull"`
	StudentID   uuid.UUID `bun:"student_id,type:uuid,notnull"`
	TeacherID   uuid.UUID `bun:"teacher_id,type:uuid,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
