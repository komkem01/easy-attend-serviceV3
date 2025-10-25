package classroom

import (
	"context"
	"time"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type InfoServiceResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SchoolID  uuid.UUID `json:"school_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Service) InfoService(ctx context.Context, id uuid.UUID) (*InfoServiceResponse, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`school.svc.info.start`)

	data, err := s.db.GetByIDClassroom(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	span.AddEvent(`school.svc.info.end`)
	return &InfoServiceResponse{
		ID:        data.ID,
		Name:      data.Name,
		SchoolID:  data.SchoolID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}
