package school

import (
	"context"
	"log/slog"
	"time"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/google/uuid"
)

type ListServiceRequest struct {
	base.RequestPaginate
}

type ListServiceResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Service) ListService(ctx context.Context, request *ListServiceRequest) ([]*ListServiceResponse, *base.ResponsePaginate, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`school.svc.list.start`)

	// Note: Current entity interface doesn't support pagination yet
	// For now, we'll get all schools and apply pagination manually
	data, err := s.db.GetListSchool(ctx)
	if err != nil {
		log.With(slog.Any(`body`, request)).Errf(`internal: %s`, err)
		return nil, nil, err
	}

	var response []*ListServiceResponse
	for _, v := range data {
		response = append(response, &ListServiceResponse{
			ID:        v.ID,
			Name:      v.Name,
			Address:   v.Address,
			Phone:     v.Phone,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	// Create pagination response (simplified for now)
	page := &base.ResponsePaginate{
		Page:  int64(request.Page),
		Size:  int64(request.Size),
		Total: int64(len(response)),
	}

	span.AddEvent(`school.svc.list.success`)
	return response, page, nil
}
