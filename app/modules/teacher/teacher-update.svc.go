package teacher

import (
	"context"
	"log/slog"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type UpdateServiceRequest struct {
	ID          uuid.UUID  `json:"id"`
	SchoolID    uuid.UUID  `json:"school_id"`
	ClassroomID *uuid.UUID `json:"classroom_id,omitempty"` // ใช้ pointer เพื่อรองรับ NULL
	PrefixID    uuid.UUID  `json:"prefix_id"`
	GenderID    uuid.UUID  `json:"gender_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
}

func (s *Service) UpdateService(ctx context.Context, req *UpdateServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.update.start`)

	_, err := s.db.UpdateTeacher(ctx, req.ID, &entitiesdto.TeacherUpdateRequest{
		ID:          req.ID,
		SchoolID:    req.SchoolID,
		ClassroomID: req.ClassroomID,
		PrefixID:    req.PrefixID,
		GenderID:    req.GenderID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
	})
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	span.AddEvent(`teacher.svc.update.end`)
	return nil
}
