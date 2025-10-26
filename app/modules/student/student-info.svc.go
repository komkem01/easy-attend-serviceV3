package student

import (
	"context"
	"log/slog"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type InfoServiceRequest struct {
	ID uuid.UUID `json:"id"`
}

type InfoServiceResponse struct {
	ID            uuid.UUID `json:"id"`
	SchoolID      uuid.UUID `json:"school_id"`
	SchoolName    string    `json:"school_name"`
	ClassroomID   uuid.UUID `json:"classroom_id"`
	ClassroomName string    `json:"classroom_name"`
	PrefixID      uuid.UUID `json:"prefix_id"`
	PrefixName    string    `json:"prefix_name"`
	GenderID      uuid.UUID `json:"gender_id"`
	GenderName    string    `json:"gender_name"`
	StudentCode   string    `json:"student_code"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Phone         string    `json:"phone"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (s *Service) InfoService(ctx context.Context, req *InfoServiceRequest) (*InfoServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`student.svc.info.start`)

	data, err := s.db.GetStudentByID(ctx, req.ID, &entitiesdto.StudentInfoResponse{})
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return nil, err
	}

	schoolData, err := s.dbSchool.GetByIDSchool(ctx, data.SchoolID)
	if err != nil {
		log.With(slog.Any(`school_id`, data.SchoolID)).Error(err)
		return nil, err
	}

	classroomData, err := s.dbClassroom.GetByIDClassroom(ctx, data.ClassroomID)
	if err != nil {
		log.With(slog.Any(`classroom_id`, data.ClassroomID)).Error(err)
		return nil, err
	}

	prefixData, err := s.dbPrefix.GetByIDPrefix(ctx, data.PrefixID)
	if err != nil {
		log.With(slog.Any(`prefix_id`, data.PrefixID)).Error(err)
		return nil, err
	}

	genderData, err := s.dbGender.GetByIDGender(ctx, data.GenderID)
	if err != nil {
		log.With(slog.Any(`gender_id`, data.GenderID)).Error(err)
		return nil, err
	}

	response := &InfoServiceResponse{
		ID:            data.ID,
		SchoolID:      schoolData.ID,
		SchoolName:    schoolData.Name,
		ClassroomID:   classroomData.ID,
		ClassroomName: classroomData.Name,
		PrefixID:      prefixData.ID,
		PrefixName:    prefixData.Name,
		GenderID:      genderData.ID,
		GenderName:    genderData.Name,
		StudentCode:   data.StudentCode,
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		Phone:         data.Phone,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	span.AddEvent(`student.svc.info.end`)
	return response, nil
}
