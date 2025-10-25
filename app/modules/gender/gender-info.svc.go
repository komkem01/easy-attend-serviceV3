package gender

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type InfoServiceResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (s *Service) InfoService(ctx context.Context, id uuid.UUID) (*InfoServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`gender.svc.info.start`)

	data, err := s.db.GetByIDGender(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	span.AddEvent(`gender.svc.info.end`)
	return &InfoServiceResponse{
		ID:   data.ID,
		Name: data.Name,
	}, nil
}
