package entitiesdto

import "github.com/google/uuid"

type ClassroomMemberCreateRequest struct {
	ClassroomID uuid.UUID `json:"classroom_id" binding:"required,uuid"`
	TeacherID   uuid.UUID `json:"teacher_id" binding:"required,uuid"`
	StudentID   uuid.UUID `json:"student_id" binding:"required,uuid"`
}

type ClassroomMemberUpdateRequest struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id" binding:"required,uuid"`
	TeacherID   uuid.UUID `json:"teacher_id" binding:"required,uuid"`
	StudentID   uuid.UUID `json:"student_id" binding:"required,uuid"`
}

type ClassroomMemberResponse struct {
	ID          uuid.UUID `json:"id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	StudentID   uuid.UUID `json:"student_id"`
}
