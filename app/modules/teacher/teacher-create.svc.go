package teacher

import (
	"context"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type CreateServiceRequest struct {
	SchoolID    uuid.UUID `json:"school_id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	PrefixID    uuid.UUID `json:"prefix_id"`
	GenderID    uuid.UUID `json:"gender_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
}

type CreateServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	SchoolID    uuid.UUID `json:"school_id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	PrefixID    uuid.UUID `json:"prefix_id"`
	GenderID    uuid.UUID `json:"gender_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
}

func (s *Service) CreateService(ctx context.Context, req *CreateServiceRequest) (*CreateServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.create.start`)

	teacher, err := s.db.CreateTeacher(ctx, &entitiesdto.TeacherCreateRequest{
		SchoolID:    req.SchoolID,
		ClassroomID: req.ClassroomID,
		PrefixID:    req.PrefixID,
		GenderID:    req.GenderID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		Phone:       req.Phone,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	response := &CreateServiceResponse{
		ID:          teacher.ID,
		SchoolID:    teacher.SchoolID,
		ClassroomID: teacher.ClassroomID,
		PrefixID:    teacher.PrefixID,
		GenderID:    teacher.GenderID,
		FirstName:   teacher.FirstName,
		LastName:    teacher.LastName,
		Email:       teacher.Email,
		Phone:       teacher.Phone,
	}

	span.AddEvent(`teacher.svc.create.end`)
	return response, nil
}
