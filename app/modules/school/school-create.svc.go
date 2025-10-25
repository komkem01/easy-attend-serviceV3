package school

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
)

type CreateServiceRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (s *Service) CreateService(ctx context.Context, req *CreateServiceRequest) error {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`school.svc.create.start`)

	_, err := s.db.CreateSchool(ctx, req.Name, req.Address, req.Phone)
	if err != nil {
		log.Error(err)
		return err
	}
	span.AddEvent(`school.svc.create.end`)
	return nil
}
