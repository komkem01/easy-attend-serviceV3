package classroom

import (
	"context"
	"time"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/google/uuid"
)

type CreateServiceRequest struct {
	SchoolID  uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Service) CreateService(ctx context.Context, req *CreateServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`classroom.svc.create.start`)

	_, err := s.db.CreateClassroom(ctx, req.SchoolID, req.Name)
	if err != nil {
		log.Error(err)
		return err
	}
	span.AddEvent(`classroom.svc.create.end`)
	return nil
}
