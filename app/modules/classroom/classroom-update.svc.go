package classroom

import (
	"context"
	"log/slog"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type UpdateServiceRequest struct {
	ID       uuid.UUID `json:"id"`
	SchoolID uuid.UUID `json:"school_id"`
	Name     string    `json:"name"`
}

func (s *Service) UpdateService(ctx context.Context, req *UpdateServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`school.svc.update.start`)

	_, err := s.db.UpdateClassroom(ctx, req.ID, req.SchoolID, req.Name)
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	span.AddEvent(`school.svc.update.end`)
	return nil
}
