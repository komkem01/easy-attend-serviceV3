package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SchoolEntity struct {
	bun.BaseModel `bun:"table:schools"`

	ID        uuid.UUID `bun:"type:uuid,default:gen_random_uuid(),pk"`
	Name      string    `bun:"type:varchar(100),notnull"`
	Address   string    `bun:"type:varchar(255)"`
	Phone     string    `bun:"type:varchar(15)"`
	CreatedAt time.Time `bun:"type:timestamptz,default:current_timestamp,notnull"`
	UpdatedAt time.Time `bun:"type:timestamptz,default:current_timestamp,notnull"`
}
