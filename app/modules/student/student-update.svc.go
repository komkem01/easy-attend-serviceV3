package student

import (
	"context"
	"log/slog"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type UpdateServiceRequest struct {
	ID          uuid.UUID `json:"id"`
	SchoolID    uuid.UUID `json:"school_id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	PrefixID    uuid.UUID `json:"prefix_id"`
	GenderID    uuid.UUID `json:"gender_id"`
	StudentCode string    `json:"student_code"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Phone       string    `json:"phone"`
}

func (s *Service) UpdateService(ctx context.Context, req *UpdateServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`student.svc.update.start`)

	_, err := s.db.UpdateStudent(ctx, req.ID, &entitiesdto.StudentUpdateRequest{
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
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	span.AddEvent(`student.svc.update.end`)
	return nil
}
