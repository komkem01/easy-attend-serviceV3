package school

import (
	"context"
	"log/slog"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type DeleteServiceRequest struct {
	ID uuid.UUID
}

func (s *Service) DeleteService(ctx context.Context, req *DeleteServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`school.svc.delete.start`)

	// Check if school exists first
	exists, err := s.db.CheckExistSchool(ctx, req.ID)
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	if !exists {
		log.With(slog.Any(`body`, req)).Warnf("school not found")
		return nil // School already doesn't exist, consider it successful
	}

	err = s.db.DeleteSchool(ctx, req.ID)
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	span.AddEvent(`school.svc.delete.end`)
	return nil
}
