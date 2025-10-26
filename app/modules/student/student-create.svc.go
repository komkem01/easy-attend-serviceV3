package student

import (
	"context"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type CreateServiceRequest struct {
	SchoolID    uuid.UUID
	ClassroomID uuid.UUID // จะใช้ uuid.Nil สำหรับกรณีที่ไม่มีค่า
	GenderID    uuid.UUID
	PrefixID    uuid.UUID
	StudentCode string
	FirstName   string
	LastName    string
	Phone       string
}

type CreateServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	SchoolID    uuid.UUID `json:"school_id"`
	ClassroomID uuid.UUID `json:"classroom_id"` // จะใช้ uuid.Nil สำหรับกรณีที่ไม่มีค่า
	PrefixID    uuid.UUID `json:"prefix_id"`
	GenderID    uuid.UUID `json:"gender_id"`
	StudentCode string    `json:"student_code"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Phone       string    `json:"phone"`
}

func (s *Service) CreateService(ctx context.Context, req *CreateServiceRequest) (*CreateServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`student.svc.create.start`)

	student, err := s.db.CreateStudent(ctx, &entitiesdto.StudentCreateRequest{
		School:      req.SchoolID,
		Classroom:   req.ClassroomID,
		Prefix:      req.PrefixID,
		Gender:      req.GenderID,
		StudentCode: req.StudentCode,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Phone:       req.Phone,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &CreateServiceResponse{
		ID:          student.ID,
		SchoolID:    student.SchoolID,
		ClassroomID: student.ClassroomID,
		PrefixID:    student.PrefixID,
		GenderID:    student.GenderID,
		StudentCode: student.StudentCode,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Phone:       student.Phone,
	}

	span.AddEvent(`student.svc.create.end`)
	return response, nil
}
