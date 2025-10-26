package student

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
	span.AddEvent(`student.svc.delete.start`)

	err := s.db.DeleteStudent(ctx, req.ID)
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	span.AddEvent(`student.svc.delete.end`)
	return nil
}
