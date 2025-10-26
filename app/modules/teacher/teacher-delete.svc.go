package teacher

import (
	"context"
	"log/slog"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type DeleteServiceRequest struct {
	ID uuid.UUID `json:"id"`
}

func (s *Service) DeleteService(ctx context.Context, req *DeleteServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.delete.start`)

	// Check if teacher exists first
	exists, err := s.db.CheckExistTeacher(ctx, req.ID)
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	if !exists {
		log.With(slog.Any(`body`, req)).Warnf("teacher not found")
		return nil // Teacher already doesn't exist, consider it successful
	}

	err = s.db.DeleteTeacher(ctx, req.ID)
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	span.AddEvent(`teacher.svc.delete.end`)
	return nil
}
