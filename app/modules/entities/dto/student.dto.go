package entitiesdto

import "github.com/google/uuid"

type StudentListResponse struct {
	ID          uuid.UUID `json:"id"`
	School      uuid.UUID `json:"school_id"`
	Classroom   uuid.UUID `json:"classroom_id"` // จะใช้ uuid.Nil สำหรับกรณีที่ไม่มีค่า
	Prefix      uuid.UUID `json:"prefix_id"`
	Gender      uuid.UUID `json:"gender_id"`
	StudentCode string    `json:"student_code"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Phone       string    `json:"phone"`
}

type StudentInfoResponse struct {
	ID          uuid.UUID         `json:"id"`
	School      SchoolResponse    `json:"school"`
	Classroom   ClassroomResponse `json:"classroom"`
	Prefix      PrefixResponse    `json:"prefix"`
	Gender      GenderResponse    `json:"gender"`
	StudentCode string            `json:"student_code"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Phone       string            `json:"phone"`
}

type StudentCreateRequest struct {
	School      uuid.UUID `json:"school_id" validate:"required"`
	Classroom   uuid.UUID `json:"classroom_id"` // ไม่บังคับ - จะใช้ uuid.Nil สำหรับกรณีที่ไม่มีค่า
	Prefix      uuid.UUID `json:"prefix_id" validate:"required"`
	Gender      uuid.UUID `json:"gender_id" validate:"required"`
	StudentCode string    `json:"student_code" validate:"required"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	Phone       string    `json:"phone"`
}

type StudentUpdateRequest struct {
	School      uuid.UUID `json:"school_id" validate:"required"`
	Classroom   uuid.UUID `json:"classroom_id"` // ไม่บังคับ - จะใช้ uuid.Nil สำหรับกรณีที่ไม่มีค่า
	Prefix      uuid.UUID `json:"prefix_id" validate:"required"`
	Gender      uuid.UUID `json:"gender_id" validate:"required"`
	StudentCode string    `json:"student_code" validate:"required"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	Phone       string    `json:"phone"`
}
