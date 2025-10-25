package school

import (
	"context"
	"log/slog"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type UpdateServiceRequest struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
}

func (s *Service) UpdateService(ctx context.Context, req *UpdateServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`school.svc.update.start`)

	_, err := s.db.UpdateSchool(ctx, req.ID, req.Name, req.Address, req.Phone)
	if err != nil {
		log.With(slog.Any(`body`, req)).Error(err)
		return err
	}

	span.AddEvent(`school.svc.update.end`)
	return nil
}
