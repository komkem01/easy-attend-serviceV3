package entitiesdto

import "github.com/google/uuid"

type GenderResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
