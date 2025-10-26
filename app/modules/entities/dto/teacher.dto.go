package entitiesdto

import (
	"time"

	"github.com/google/uuid"
)

type TeacherCreateRequest struct {
	SchoolID    uuid.UUID  `json:"school_id"`
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"` // ใช้ pointer เพื่อรองรับ NULL
	PrefixID    uuid.UUID  `json:"prefix_id"`
	GenderID    uuid.UUID  `json:"gender_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Phone       string     `json:"phone"`
}

type TeacherUpdateRequest struct {
	ID          uuid.UUID  `json:"id"`
	SchoolID    uuid.UUID  `json:"school_id"`
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"` // ใช้ pointer เพื่อรองรับ NULL
	PrefixID    uuid.UUID  `json:"prefix_id"`
	GenderID    uuid.UUID  `json:"gender_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Phone       string     `json:"phone"`
}

type TeacherInfoResponse struct {
	ID          uuid.UUID `json:"id"`
	SchoolID    uuid.UUID `json:"school_id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	PrefixID    uuid.UUID `json:"prefix_id"`
	GenderID    uuid.UUID `json:"gender_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
