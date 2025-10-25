package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type GenderEntity struct {
	bun.BaseModel `bun:"table:genders"`

	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Name      string    `bun:"name,notnull,unique"`
	CreatedAt time.Time `bun:"type:timestamptz,default:current_timestamp,notnull"`
	UpdatedAt time.Time `bun:"type:timestamptz,default:current_timestamp,notnull"`
}
