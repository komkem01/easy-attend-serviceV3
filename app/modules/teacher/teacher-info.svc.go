package teacher

import (
	"context"
	"log/slog"
	"time"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type InfoServiceRequest struct {
	ID uuid.UUID `json:"id"`
}

type InfoServiceResponse struct {
	ID            uuid.UUID  `json:"id"`
	SchoolID      uuid.UUID  `json:"school_id"`
	SchoolName    string     `json:"school_name"`
	ClassroomID   *uuid.UUID `json:"classroom_id,omitempty"`   // อาจเป็น null
	ClassroomName *string    `json:"classroom_name,omitempty"` // อาจเป็น null
	PrefixID      uuid.UUID  `json:"prefix_id"`
	PrefixName    string     `json:"prefix_name"`
	GenderID      uuid.UUID  `json:"gender_id"`
	GenderName    string     `json:"gender_name"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (s *Service) InfoService(ctx context.Context, req *InfoServiceRequest) (*InfoServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.info.start`)

	data, err := s.db.GetByIDTeacher(ctx, req.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	schoolData, err := s.dbSchool.GetByIDSchool(ctx, data.SchoolID)
	if err != nil {
		log.With(slog.Any(`school_id`, data.SchoolID)).Error(err)
		return nil, err
	}

	// จัดการ classroom data (อาจเป็น null)
	var classroomIDPtr *uuid.UUID
	var classroomNamePtr *string
	if data.ClassroomID != nil {
		classroomData, err := s.dbClassroom.GetByIDClassroom(ctx, *data.ClassroomID)
		if err != nil {
			log.With(slog.Any(`classroom_id`, *data.ClassroomID)).Error(err)
			return nil, err
		}
		classroomIDPtr = &classroomData.ID
		classroomNamePtr = &classroomData.Name
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
		SchoolID:      data.SchoolID,
		SchoolName:    schoolData.Name,
		ClassroomID:   classroomIDPtr,
		ClassroomName: classroomNamePtr,
		PrefixID:      data.PrefixID,
		PrefixName:    prefixData.Name,
		GenderID:      data.GenderID,
		GenderName:    genderData.Name,
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		Email:         data.Email,
		Phone:         data.Phone,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	span.AddEvent(`teacher.svc.info.end`)
	return response, nil
}
