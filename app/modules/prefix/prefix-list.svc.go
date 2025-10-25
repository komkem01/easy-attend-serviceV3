package prefix

import (
	"context"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/google/uuid"
)

type ListServiceRequest struct {
	base.RequestPaginate
}

type ListServiceResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (s *Service) ListService(ctx context.Context, request *ListServiceRequest) ([]*ListServiceResponse, *base.ResponsePaginate, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`prefix.svc.list.start`)

	data, err := s.db.GetListPrefix(ctx)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	var responses []*ListServiceResponse
	for _, item := range data {
		responses = append(responses, &ListServiceResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return responses, nil, nil
}
