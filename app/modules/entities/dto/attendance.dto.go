package entitiesdto

import (
	"time"

	"github.com/google/uuid"
)

type AttendanceCreateRequest struct {
	ClassroomID uuid.UUID `json:"classroom_id" binding:"required,uuid"`
	TeacherID   uuid.UUID `json:"teacher_id" binding:"required,uuid"`
	StudentID   uuid.UUID `json:"student_id" binding:"required,uuid"`
	Date        string    `json:"date" binding:"required"`   // YYYY-MM-DD format
	Time        string    `json:"time" binding:"required"`   // HH:MM:SS format
	Status      string    `json:"status" binding:"required"` // present, absent, late, excused
}

type AttendanceUpdateRequest struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id" binding:"required,uuid"`
	TeacherID   uuid.UUID `json:"teacher_id" binding:"required,uuid"`
	StudentID   uuid.UUID `json:"student_id" binding:"required,uuid"`
	Date        string    `json:"date" binding:"required"`
	Time        string    `json:"time" binding:"required"`
	Status      string    `json:"status" binding:"required"`
}

type AttendanceListRequest struct {
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"`
	StudentID   *uuid.UUID `json:"student_id,omitempty"`
	Date        *string    `json:"date,omitempty"`
	Status      *string    `json:"status,omitempty"`
}

type AttendanceResponse struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	StudentID   uuid.UUID `json:"student_id"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
