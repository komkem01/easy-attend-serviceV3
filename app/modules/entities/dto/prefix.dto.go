package entitiesdto

import "github.com/google/uuid"

type PrefixResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
